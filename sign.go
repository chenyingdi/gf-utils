package gfUtils

import (
	"errors"
	"github.com/chenyingdi/wx-service"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type WxSignInParam struct {
	Code      string `json:"code"`
	Nickname  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
}

// 密码登录
func SignInWithPassword(tag, val, password, table, key string) (gdb.Record, *Err) {
	e := NewErr()

	user, err := g.DB().Table(table).FindOne(tag, val)
	e.Append(err)

	if user.IsEmpty() {
		e.Append(errors.New("用户名或密码错误！"))
		return nil, e
	}

	if !PBKDF2Decode(password, user["password"].String(), 1024, func() string {
		return key
	}) {
		e.Append(errors.New("用户名或密码错误！"))
		return nil, e
	}

	return user, e
}

func SignInWithWx(p WxSignInParam, appID, appSecret string) (string, *Err) {
	e := NewErr()

	c := WxService.NewClient(&WxService.ClientConfig{
		AppID:     appID,
		AppSecret: appSecret,
		IsSandBox: false,
	})

	openid := c.Code2Session(p.Code)
	e.Append(c.Err())

	return openid, e
}
