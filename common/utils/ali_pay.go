package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	alipay_zfb "github.com/smartwalle/alipay/v3"
)


// 支付宝支付
func AliPayZfb(notify_url, return_url, order_number, pay_amount string) string {
	privateKey := beego.AppConfig.String("ali_private_key")
	appId := beego.AppConfig.String("ali_app_id")
	client, err := alipay_zfb.New(appId, privateKey, true)
	if err != nil {
		logs.Error(err.Error())
	}

	var p = alipay_zfb.TradeAppPay{}
	p.NotifyURL = notify_url
	p.ReturnURL = return_url
	p.Subject = beego.AppConfig.String("ali_project_name")
	p.OutTradeNo = order_number
	p.TotalAmount = pay_amount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	if client != nil {
		err = client.LoadAliPayPublicCertFromFile("./gingernet.vip.csr")
		if err != nil {
			logs.Error("load crt fail")
			return ""
		}
		err = client.LoadAliPayPublicKey(beego.AppConfig.String("ali_public_key"))
		if err != nil {
			logs.Error("load public key fail")
			return ""
		}
		url, err := client.TradeAppPay(p)
		if err != nil {
			logs.Error(err.Error())
		}
		return url
	} else {
		return ""
	}
}

