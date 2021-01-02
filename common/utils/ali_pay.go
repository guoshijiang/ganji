package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/smartwalle/alipay/v3"
)


// 支付宝支付
func AliPayZfb(notify_url, return_url, order_number, pay_amount string) string {
	privateKey := beego.AppConfig.String("private_key")
	appId := beego.AppConfig.String("app_id")
	//aliPublicKey :=  beego.AppConfig.String("public_key")
	client, err := alipay.New(appId, privateKey, true)
	if err != nil {
		logs.Error(err.Error())
	}
	var p = alipay.TradeAppPay{}
	p.NotifyURL = notify_url
	p.ReturnURL = return_url
	p.Subject = beego.AppConfig.String("project_name")
	p.OutTradeNo = order_number
	p.TotalAmount = pay_amount
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err := client.TradeAppPay(p)
	if err != nil {
		logs.Error(err.Error())
	}
	return url
}

