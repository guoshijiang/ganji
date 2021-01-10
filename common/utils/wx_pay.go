package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
	"strconv"
	"time"
)


type WxPayData struct {
	Timestamp string  `json:"timestamp"`
    Appid     string  `json:"appid"`
	Partnerid string  `json:"partnerid"`
	Prepayid  string  `json:"prepayid"`
	Noncestr  string  `json:"noncestr"`
	Package   string  `json:"package"`
	Paysign   string  `json:"paysign"`
	Signtype  string  `json:"signtype"`
}


// 微信支付代码实现
func WxPayOrder(order_number string, total_fee float64) (*WxPayData, error) {
	wx_app_id := beego.AppConfig.String("wx_app_id")
	// wx_serial_number := beego.AppConfig.String("wx_serial_number")
	mch_id := beego.AppConfig.String("wx_mch_id")
	v3_api_key := beego.AppConfig.String("wx_api_v3_key")
	// wx_pk_content := beego.AppConfig.String("wx_pk_content")
	wx_ip := beego.AppConfig.String("wx_ip")
	notify_url := beego.AppConfig.String("notify_url")
	wx_body := beego.AppConfig.String("wx_body")
	client := wechat.NewClient(wx_app_id, mch_id, v3_api_key, true)
	client.DebugSwitch = gopay.DebugOn
	client.SetCountry(wechat.China)
	err := client.AddCertFilePath(
		"/root/market/src/ganji/crt/apiclient_cert.pem",
		"/root/market/src/ganji/crt/apiclient_key.pem",
		"/root/market/src/ganji/crt/apiclient_cert.p12")
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	bm := make(gopay.BodyMap)
	nonce_str := gotil.GetRandomString(32)
	bm.Set("nonce_str", nonce_str).
		Set("body", wx_body).
		Set("out_trade_no", order_number).
		Set("total_fee", total_fee).
		Set("spbill_create_ip", wx_ip).
		Set("notify_url", notify_url).
		Set("trade_type",  wechat.TradeType_App).
		Set("sign_type",  wechat.SignType_MD5).
		SetBodyMap("scene_info", func(bm gopay.BodyMap) {
			bm.SetBodyMap("app_info", func(bm gopay.BodyMap) {
				bm.Set("type", "App")
				bm.Set("wap_url", "notify_url")
				bm.Set("wap_name", "市集APP支付")
			})
		})
	sign := wechat.GetParamSign(wx_app_id, mch_id, v3_api_key, bm)
	bm.Set("sign", sign)
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	paySign := wechat.GetAppPaySign(wx_app_id, mch_id, wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, v3_api_key)
	data_ret := WxPayData{
		Timestamp: timeStamp,
		Appid: wx_app_id,
		Partnerid: mch_id,
		Prepayid: wxRsp.PrepayId,
		Noncestr: nonce_str,
		Package: "Sign=WXPay",
		Paysign: paySign,
		Signtype: wechat.SignType_MD5,
	}
	return &data_ret, nil
}

