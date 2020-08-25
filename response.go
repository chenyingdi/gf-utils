package gfUtils

import "github.com/gogf/gf/net/ghttp"

type Resp struct {
	req  *ghttp.Request
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(r *ghttp.Request) *Resp {
	return &Resp{req: r}
}

func (r *Resp) JSON(code int, msg string, data interface{}) {
	r.Code = code
	r.Msg = msg
	r.Data = data
	_ = r.req.Response.WriteJsonExit(r)
}

func (r *Resp) SUCCESS(data interface{}) {
	r.JSON(
		200,
		"操作成功！",
		data,
	)
}

func (r *Resp) FAIL(data interface{}) {
	r.JSON(
		500,
		"操作失败！",
		data,
	)
}

func (r *Resp) UNAUTHORIZED(data interface{}) {
	r.JSON(
		400,
		"授权失败！",
		data,
	)
}
