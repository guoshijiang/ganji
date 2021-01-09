package utils

import "github.com/iGoogle-ink/gopay/wechat"

func wx_pay()  {
	cli := wechat.NewClient("wxdaa2ab9ef87b5497", "mchId", "apiKey", false)
	cli.SetCountry(wechat.China)
	cli.AddCertFilePath("", "", "")
}



