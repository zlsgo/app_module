package account

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/restapi"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztype"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/sohaha/zlsgo/zvalid"
	"github.com/zlsgo/app_core/common"
	"github.com/zlsgo/app_core/service"
	"golang.org/x/crypto/bcrypt"
)

type Index struct {
	service.App
	accoutModel *restapi.Model
	permModel   *restapi.Model
	roleModel   *restapi.Model
	plugin      *Module
	Path        string
}

var (
	_ = reflect.TypeOf(&Index{})
)

const saltLen = 4

func (h *Index) Init(r *znet.Engine) error {
	// 登录无需验证
	r.POST("/login", h.login)

	// 获取系统信息无需验证
	r.GET("/site", h.getSite)

	// 获取系统信息无需验证
	r.Any("/refresh-token", h.refreshToken)

	return h.plugin.RegMiddleware(r)
}

func (h *Index) refreshToken(c *znet.Context) (interface{}, error) {
	token := jwt.GetToken(c)
	refreshToken := c.DefaultFormOrQuery("refresh_token", "")
	if refreshToken == "" {
		refreshToken = c.GetJSON("refresh_token").String()
	}

	_, err := jwt.Parse(token, h.plugin.Options.key)
	if err != nil && !strings.Contains(err.Error(), "expired") {
		return nil, zerror.WrapTag(zerror.InvalidInput)(errors.New("旧 token 不合法"))
	}

	info, err := jwt.Parse(refreshToken, h.plugin.Options.key)
	if err != nil {
		return nil, zerror.WrapTag(zerror.InvalidInput)(errors.New("refresh_token 无效"))
	}
	if !info.IsRefresh {
		return nil, zerror.WrapTag(zerror.InvalidInput)(errors.New("非合法 refresh_token"))
	}

	salt := info.Info[:saltLen]
	uid := info.Info[saltLen:]
	f, err := restapi.FindCols(h.accoutModel, "salt", uid)
	if err != nil || f.Index(0).String() != salt {
		return nil, zerror.WrapTag(zerror.InvalidInput)(errors.New("refresh_token 已失效"))
	}

	salt = zstring.Rand(saltLen)
	err = updateUser(h.accoutModel, uid, ztype.Map{
		"salt": salt,
	})
	if err != nil {
		return nil, err
	}

	clearCache(token, uid)

	accessToken, refreshToken, err := jwt.GenToken(salt+uid, h.plugin.Options.key, h.plugin.Options.Expire)
	if err != nil {
		return nil, err
	}

	return ztype.Map{
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// GetMe 获取当前用户信息
func (h *Index) GetMe(c *znet.Context) (interface{}, error) {
	// TODO: 考虑做缓存处理
	info, err := restapi.FindOne(h.accoutModel, common.VarUID(c), func(so *restapi.CondOptions) error {
		so.Fields = h.accoutModel.GetFields("password", "salt")
		return nil
	})
	if err != nil {
		return nil, err
	}
	if info.IsEmpty() {
		return nil, zerror.WrapTag(zerror.InvalidInput)(errors.New("用户不存在"))
	}

	perms, _ := restapi.FindCols(h.roleModel, "permission", ztype.Map{
		"alias": info.Get("role").SliceString(),
	})
	permIDs := make([]int, 0)
	for i := range perms {
		permIDs = append(permIDs, perms[i].SliceInt()...)
	}
	permission, _ := restapi.FindCols(h.permModel, "alias", ztype.Map{
		restapi.IDKey: zarray.Unique(permIDs),
		"alias !=":    "",
	}, func(o *restapi.CondOptions) error {
		o.Fields = []string{"alias"}
		return nil
	})

	data := ztype.Map{
		"info":       info,
		"permission": permission,
	}
	return data, nil
}

var loginLimit = zcache.NewFast()

// isBusyLogin 短时间内登录失败超过指定次数禁止登录
func isBusyLogin(c *znet.Context) (b bool) {
	ip := c.GetClientIP()
	total, ok := loginLimit.Get(ip)
	if !ok {
		total = 0
	}

	b = ztype.ToInt(total) >= 5
	if b {
		loginFailed(c)
	}
	return
}

// loginFailed 登录失败
func loginFailed(c *znet.Context) {
	ip := c.GetClientIP()
	total, _ := loginLimit.Get(ip)
	data := ztype.ToInt(total) + 1
	loginLimit.Set(ip, data, zutil.BackOffDelay(ztype.ToInt(total), time.Hour/2))
}

// login 登录
func (h *Index) login(c *znet.Context) (result interface{}, err error) {
	if isBusyLogin(c) {
		return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("登录失败次数过多，请稍后再试"))
	}

	json, _ := c.GetJSONs()
	account := json.Get("account").String()
	password := json.Get("password").String()

	invalidInput := zerror.WrapTag(zerror.InvalidInput)
	if account == "" {
		err = invalidInput(errors.New("请输入账号"))
		return
	}

	if password == "" {
		err = invalidInput(errors.New("请输入密码"))
		return
	}

	user, err := restapi.FindOne(h.accoutModel, ztype.Map{
		"account": account,
	})
	if err != nil {
		return nil, invalidInput(err)
	}

	defer func() {
		if err != nil {
			loginFailed(c)
		}
	}()

	if user.IsEmpty() {
		err = invalidInput(errors.New("账号或密码错误"))
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), zstring.String2Bytes(password))
	if err != nil {
		err = invalidInput(errors.New("账号或密码错误"))
		return
	}

	status := user.Get("status").Int()
	if status != 1 {
		switch status {
		case 0:
			return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("账号待激活"))
		default:
			return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("账号已停用"))
		}
	}

	salt := user.Get("salt").String()

	if h.plugin.Options.Only || salt == "" {
		salt = zstring.Rand(saltLen)
	}

	uid := user.Get(restapi.IDKey).String()
	err = updateUser(h.accoutModel, uid, ztype.Map{
		"salt":     salt,
		"login_at": ztime.Now(),
	})

	if err != nil {
		return nil, err
	}

	info := salt + uid
	accessToken, refreshToken, err := jwt.GenToken(info, h.plugin.Options.key, h.plugin.Options.Expire)

	if err != nil {
		return nil, err
	}

	if mLog, ok := h.plugin.mods.Get(logsName); ok {
		_, _ = insertLog(c, mLog, user.Get("account").String(), 200, "登录成功")
	}

	return ztype.Map{
		"uid":           uid,
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// AnyLogout 用户退出
func (h *Index) AnyLogout(c *znet.Context) (any, error) {
	uid := common.VarUID(c)
	if uid == "" {
		return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("请先登录"))
	}

	err := updateUser(h.accoutModel, uid, ztype.Map{
		"salt": "",
	})

	if err == nil {
		clearCache(jwt.GetToken(c), uid)
	}
	return nil, err
}

// AnyPassword 修改密码
func (h *Index) AnyPassword(c *znet.Context) (data any, err error) {
	defer func() {
		if err != nil {
			WithLog(c, "修改密码", []byte(err.Error()))
		} else {
			WithLog(c, "修改密码", []byte("修改成功"))
		}
	}()
	var (
		old      string
		password string
	)
	rule := c.ValidRule().Required()
	err = zvalid.Batch(
		zvalid.BatchVar(&old, c.Valid(rule, "old_password", "旧密码")),
		zvalid.BatchVar(&password, c.Valid(rule, "password", "新密码")),
	)

	invalidInput := zerror.WrapTag(zerror.InvalidInput)
	if err != nil {
		return nil, invalidInput(err)
	}

	uid := common.VarUID(c)
	user, _ := restapi.FindOne(h.accoutModel, uid, func(so *restapi.CondOptions) error {
		so.Fields = []string{restapi.IDKey, "password", "salt"}
		return nil
	})
	if user.IsEmpty() {
		return nil, invalidInput(errors.New("用户不存在"))
	}

	err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), zstring.String2Bytes(old))
	if err != nil {
		return nil, invalidInput(errors.New("旧密码错误"))
	}

	salt := zstring.Rand(saltLen)
	err = updateUser(h.accoutModel, uid, ztype.Map{
		"salt":     salt,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	clearCache(jwt.GetToken(c), uid)

	info := salt + uid
	accessToken, refreshToken, err := jwt.GenToken(info, h.plugin.Options.key, h.plugin.Options.Expire)

	return ztype.Map{
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, err
}

// PatchMe 修改当前用户信息
func (h *Index) PatchMe(c *znet.Context) (any, error) {
	uid := common.VarUID(c)
	data, _ := c.GetJSONs()
	keys := []string{"avatar", "nickname"}
	update := make(ztype.Map, 0)
	for k, v := range data.Map() {
		if !zarray.Contains(keys, k) {
			continue
		}
		update[k] = v
	}
	err := updateUser(h.accoutModel, uid, update)
	return nil, err
}

// POSTAvatar 上传用户头像
func (h *Index) POSTAvatar(c *znet.Context) (any, error) {
	uid := common.VarUID(c)
	res, err := common.Upload(c, "account", func(o *common.UploadOption) {
		o.Dir = "/avatar"
		o.MimeType = []string{"image/*"}
	})
	if err != nil {
		return nil, err
	}

	err = updateUser(h.accoutModel, uid, ztype.Map{
		"avatar": res[0].Path,
	})
	return res[0].Path, err
}

func updateUser(m *restapi.Model, id string, data ztype.Map) error {
	_, err := restapi.Update(m, id, data)
	return err
}
