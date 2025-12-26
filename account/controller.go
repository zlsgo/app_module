package account

import (
	"errors"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/zlsgo/app_module/account/jwt"
	"github.com/zlsgo/app_module/account/limiter"
	"github.com/zlsgo/app_module/model"

	"github.com/sohaha/zlsgo/zarray"
	"github.com/sohaha/zlsgo/zerror"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	zsession "github.com/sohaha/zlsgo/znet/session"
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
	r.Use(limiter.IPMiddleware())

	noPermRoutes := []string{
		"/login",
		"/refresh-token",
		"/register",
		"/site",
	}
	err := h.module.UsePermisMiddleware(r, nil, zarray.Map(noPermRoutes, func(i int, v string) string {
		return strings.TrimRight(h.Path, "/") + v
	})...)
	if err != nil {
		return err
	}

	// 登录
	r.POST("/login", h.login)

	// 获取系统信息
	r.GET("/site", h.getSite)

	// 刷新 token
	r.Any("/refresh-token", h.refreshToken)

	// 用户注册
	r.POST("/register", h.register)

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
	f, err := model.FindCols[string](h.accoutModel.Model(), "salt", model.ID(uid))
	if err != nil || len(f) == 0 || f[0] != salt {
		return nil, zerror.InvalidInput.Text("refresh_token 已失效")
	}

	salt = zstring.Rand(saltLen)
	err = h.module.updateUser(h.accoutModel, uid, ztype.Map{
		"salt": salt,
	})
	if err != nil {
		return nil, err
	}

	h.module.clearCache(token, uid)

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
	// 使用缓存获取用户信息
	uid := h.module.Request.UID(c)
	userInfo, err := h.module.getUserForCache(h.accoutModel, uid)
	if err != nil {
		// 如果缓存失败，直接查询数据库
		info, err := model.FindOne[ztype.Map](h.accoutModel.Model(), model.ID(uid), func(so *model.CondOptions) {
			so.Fields = h.accoutModel.GetFields("password", "salt")
		})
		if err != nil {
			if errors.Is(err, model.ErrNoRecord) {
				return nil, zerror.InvalidInput.Text("用户不存在")
			}
			return nil, err
		}
		userInfo = info
	}
	_ = userInfo.Delete("password")
	_ = userInfo.Delete("salt")

	perms, _ := model.FindCols[ztype.Type](h.roleModel.Model(), "permission", model.Filter{
		"alias": userInfo.Get("role").SliceString(),
	})
	permIDs := make([]int, 0)
	for i := range perms {
		permIDs = append(permIDs, perms[i].SliceInt()...)
	}
	permission, _ := model.FindCols[string](h.permModel.Model(), "alias", model.Filter{
		model.IDKey(): zarray.Unique(permIDs),
		"alias !=":    "",
	}, func(o *model.CondOptions) {
		o.Fields = []string{"alias"}
	})

	data := ztype.Map{
		"info":       userInfo,
		"permission": permission,
	}
	return data, nil
}

// isBusyLogin 短时间内登录失败超过指定次数禁止登录
func (m *Module) isBusyLogin(c *znet.Context) (b bool) {
	ip := c.GetClientIP()
	total, ok := m.loginLimit.Get(ip)
	if !ok {
		total = 0
	}

	b = ztype.ToInt(total) >= 5
	if b {
		m.loginFailed(c)
	}
	return
}

// loginFailed 登录失败
func (m *Module) loginFailed(c *znet.Context) {
	ip := c.GetClientIP()
	total, _ := m.loginLimit.Get(ip)
	data := ztype.ToInt(total) + 1
	m.loginLimit.Set(ip, data, zutil.BackOffDelay(ztype.ToInt(total), time.Hour/2, time.Hour))
}

// login 登录
func (h *Index) login(c *znet.Context) (result interface{}, err error) {
	if h.module.isBusyLogin(c) {
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

	user, err := model.FindOne[ztype.Map](h.accoutModel.Model(), model.Filter{
		"account": account,
	})
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, zerror.InvalidInput.Text("账号或密码错误")
		}
		return nil, zerror.InvalidInput.Text(err.Error())
	}

	defer func() {
		if err != nil {
			h.module.loginFailed(c)
		}
	}()

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
	err = h.module.updateUser(h.accoutModel, uid, ztype.Map{
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
		ip := ""
		if !h.module.noLogIP {
			ip = c.GetClientIP()
		}
		_, _ = insertLog(mLog, user.Get("account").String(), ip, c.Request.Method, c.Request.URL.String(), 200, "登录成功", c.Request.URL.Query().Encode(), "")
	}

	if h.module.Options.Session != nil {
		s, _ := zsession.Get(c)
		if s != nil {
			s.Set("token", accessToken)
			s.Set("refresh_token", refreshToken)
			s.Save()
		}
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

	err := h.module.updateUser(h.accoutModel, uid, ztype.Map{
		"salt": "",
	})

	if err == nil {
		h.module.clearCache(jwt.GetToken(c), uid)
		if h.module.Options.Session != nil {
			s, _ := zsession.Get(c)
			if s != nil {
				s.Destroy()
			}
		}
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

	if ok, msg := ValidatePassword(password, DefaultPasswordConfig); !ok {
		return nil, zerror.InvalidInput.Text(msg)
	}

	uid := h.module.Request.UID(c)
	user, err := model.FindOne[ztype.Map](h.accoutModel.Model(), model.ID(uid), func(so *model.CondOptions) {
		so.Fields = []string{model.IDKey(), "password", "salt"}
	})
	if err != nil {
		if errors.Is(err, model.ErrNoRecord) {
			return nil, zerror.InvalidInput.Text("用户不存在")
		}
		return nil, err
	}
	if user.IsEmpty() {
		return nil, zerror.InvalidInput.Text("用户不存在")
	}

	err = bcrypt.CompareHashAndPassword(user.Get("password").Bytes(), zstring.String2Bytes(old))
	if err != nil {
		return nil, zerror.InvalidInput.Text("旧密码错误")
	}

	salt := zstring.Rand(saltLen)
	err = h.module.updateUser(h.accoutModel, uid, ztype.Map{
		"salt":     salt,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	h.module.clearCache(jwt.GetToken(c), uid)

	info := salt + uid
	accessToken, refreshToken, err := jwt.GenToken(info, h.module.Options.key, h.module.Options.Expire, h.module.Options.RefreshExpire)

	if h.module.Options.Session != nil {
		s, _ := zsession.Get(c)
		if s != nil {
			s.Set("token", accessToken)
			s.Set("refresh_token", refreshToken)
			s.Save()
		}
	}
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
	err := h.module.updateUser(h.accoutModel, uid, update)
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

	avatarPath := res[0].Path
	basePath := filepath.Base(avatarPath)
	ext := filepath.Ext(basePath)
	baseName := strings.TrimSuffix(basePath, ext)
	newAvatarPath := strings.Replace(avatarPath, baseName, uid, -1)

	info, _ := h.module.getUserForCache(h.accoutModel, uid)
	oldAvatarPath := info.Get("avatar").String()
	if oldAvatarPath != "" {
		_ = zfile.Remove("." + oldAvatarPath)
	}
	err = zfile.MoveFile("."+avatarPath, "."+newAvatarPath)
	if err != nil {
		return nil, errors.New("头像保存失败")
	}

	err = h.module.updateUser(h.accoutModel, uid, ztype.Map{
		"avatar": newAvatarPath,
	})

	return newAvatarPath, err
}

func (m *Module) updateUser(schema *model.Schema, id string, data ztype.Map) error {
	_, err := model.Update(schema, model.ID(id), data)
	m.clearCache("", id)
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
