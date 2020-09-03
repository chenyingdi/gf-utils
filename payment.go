package gfUtils

import (
	WxService "github.com/chenyingdi/wx-service"
	"github.com/gogf/gf/frame/g"
)

// 微信支付
func WxPayment(fee int, orderSn, appID, mchID, apiKey, openid string) (string, *Err) {
	e := NewErr()

	c := WxService.NewClient(&WxService.ClientConfig{
		AppID:     appID,
		MchID:     mchID,
		ApiKey:    apiKey,
		IsSandBox: false,
	})

	p := WxService.NewParams()

	ip, err := WxService.GetIp()
	e.Append(err)

	p.SetString("appid", c.Client().AppID).
		SetString("mch_id", c.Client().MchID).
		SetString("nonce_str", WxService.GeneNonceStr(32)).
		SetString("body", g.Cfg().GetString("app.appName")+"-商品购买").
		SetString("out_trade_no", orderSn).
		SetInt("total_fee", fee).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", g.Cfg().GetString("app.appNotifyUrl")).
		SetString("openid", openid).
		SetString("trade_type", "JSAPI")

	prepayID := c.UnifiedOrder(p)
	e.Append(c.Err())

	return prepayID, e
}
