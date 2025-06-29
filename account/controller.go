package account

import (
	"errors"
	"reflect"
	"time"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/limiter"
	"github.com/zlsgo/app_module/model"

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
	accoutModel *model.Schema
	permModel   *model.Schema
	roleModel   *model.Schema
	module      *Module
	Path        string
}

var _ = reflect.TypeOf(&Index{})

const saltLen = 4

func (h *Index) Init(r *znet.Engine) error {
	{
		// 无需权限校验
		noPerm := r.Group("/", func(e *znet.Engine) {
			e.Use(limiter.IPMiddleware())
		})
		// 登录
		noPerm.POST("/login", h.login)

		// 获取系统信息
		noPerm.GET("/site", h.getSite)

		// 刷新 token
		noPerm.Any("/refresh-token", h.refreshToken)

		// 用户注册
		noPerm.POST("/register", h.register)
	}

	err := PermisMiddleware(r)
	if err != nil {
		return err
	}

	return nil
}

func (h *Index) refreshToken(c *znet.Context) (interface{}, error) {
	token := jwt.GetToken(c)
	refreshToken := c.DefaultFormOrQuery("refresh_token", "")
	if refreshToken == "" {
		refreshToken = c.GetJSON("refresh_token").String()
	}

	// _, err := jwt.Parse(token, h.module.Options.key)
	// err = jwt.ParseError(err)
	// if err != nil {
	// 	return nil, err
	// }

	info, err := jwt.Parse(refreshToken, h.module.Options.key)
	if err != nil || !info.IsRefresh {
		return nil, zerror.InvalidInput.Text("refresh_token 无效")
	}

	salt := info.Info[:saltLen]
	uid := info.Info[saltLen:]
	f, err := model.FindCols(h.accoutModel, "salt", uid)
	if err != nil || f.Index(0).String() != salt {
		return nil, zerror.InvalidInput.Text("refresh_token 已失效")
	}

	salt = zstring.Rand(saltLen)
	err = updateUser(h.accoutModel, uid, ztype.Map{
		"salt": salt,
	})
	if err != nil {
		return nil, err
	}

	clearCache(token, uid)

	accessToken, refreshToken, err := jwt.GenToken(salt+uid, h.module.Options.key, h.module.Options.Expire, h.module.Options.RefreshExpire)
	if err != nil {
		return nil, err
	}

	return ztype.Map{
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, nil
}

// GetInfo 获取用户信息
func (h *Index) GetInfo(c *znet.Context) (interface{}, error) {
	// TODO: 考虑做缓存处理
	info, err := model.FindOne(h.accoutModel, h.module.Request.UID(c), func(so *model.CondOptions) {
		so.Fields = h.accoutModel.GetFields("password", "salt")
	})
	if err != nil {
		return nil, err
	}

	if info.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	perms, _ := model.FindCols(h.roleModel, "permission", ztype.Map{
		"alias": info.Get("role").SliceString(),
	})
	permIDs := make([]int, 0)
	for i := range perms {
		permIDs = append(permIDs, perms[i].SliceInt()...)
	}
	permission, _ := model.FindCols(h.permModel, "alias", ztype.Map{
		model.IDKey(): zarray.Unique(permIDs),
		"alias !=":    "",
	}, func(o *model.CondOptions) {
		o.Fields = []string{"alias"}
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
	loginLimit.Set(ip, data, zutil.BackOffDelay(ztype.ToInt(total), time.Hour/2, time.Hour))
}

// login 登录
func (h *Index) login(c *znet.Context) (result interface{}, err error) {
	if isBusyLogin(c) {
		return nil, zerror.WrapTag(zerror.Unauthorized)(errors.New("登录失败次数过多，请稍后再试"))
	}

	json, _ := c.GetJSONs()
	account := json.Get("account").String()
	password := json.Get("password").String()

	if account == "" {
		err = zerror.InvalidInput.Text("请输入账号")
		return
	}

	if password == "" {
		err = zerror.InvalidInput.Text("请输入密码")
		return
	}

	user, err := model.FindOne(h.accoutModel, ztype.Map{
		"account": account,
	})
	if err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	defer func() {
		if err != nil {
			loginFailed(c)
		}
	}()

	if user.IsEmpty() {
		err = zerror.InvalidInput.Text("账号或密码错误")
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), zstring.String2Bytes(password))
	if err != nil {
		err = zerror.InvalidInput.Text("账号或密码错误")
		return
	}

	status := user.Get("status").Int()
	if status != 1 {
		switch status {
		case 0:
			return nil, zerror.Unauthorized.Text("账号待激活")
		default:
			return nil, zerror.Unauthorized.Text("账号已停用")
		}
	}

	salt := user.Get("salt").String()

	if h.module.Options.Only || salt == "" {
		salt = zstring.Rand(saltLen)
	}

	uid := user.Get(model.IDKey()).String()
	err = updateUser(h.accoutModel, uid, ztype.Map{
		"salt":     salt,
		"login_at": ztime.Now(),
	})
	if err != nil {
		return nil, err
	}

	info := salt + uid
	accessToken, refreshToken, err := jwt.GenToken(info, h.module.Options.key, h.module.Options.Expire, h.module.Options.RefreshExpire)
	if err != nil {
		return nil, err
	}

	if mLog, ok := h.module.mods.Get(logsName); ok {
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
	uid := h.module.Request.UID(c)
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
			h.module.Request.WithLog(c, "修改密码", err.Error())
		} else {
			h.module.Request.WithLog(c, "修改密码", "修改成功")
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
	if err != nil {
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	uid := h.module.Request.UID(c)
	user, _ := model.FindOne(h.accoutModel, uid, func(so *model.CondOptions) {
		so.Fields = []string{model.IDKey(), "password", "salt"}
	})
	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), zstring.String2Bytes(old))
	if err != nil {
		return nil, zerror.InvalidInput.Text("旧密码错误")
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
	accessToken, refreshToken, err := jwt.GenToken(info, h.module.Options.key, h.module.Options.Expire, h.module.Options.RefreshExpire)

	return ztype.Map{
		"token":         accessToken,
		"refresh_token": refreshToken,
	}, err
}

// PatchMe 修改当前用户信息
func (h *Index) PatchMe(c *znet.Context) (any, error) {
	uid := h.module.Request.UID(c)
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
	uid := h.module.Request.UID(c)
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

func updateUser(m *model.Schema, id string, data ztype.Map) error {
	_, err := model.Update(m, id, data)
	return err
}

// register 用户注册
func (h *Index) register(c *znet.Context) (any, error) {
	if !h.module.Options.EnableRegister {
		return nil, zerror.InvalidInput.Text("系统未开启注册")
	}

	j, _ := c.GetJSONs()
	data := ztype.Map{
		"nickname": j.Get("nickname").String(),
		"account":  j.Get("account").String(),
		"password": j.Get("password").String(),
	}
	return h.module.Inside.CreateUser(data)
}
