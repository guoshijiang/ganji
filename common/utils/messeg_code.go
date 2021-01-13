package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func SendMesseageCode(phone string, verify_code int) bool {
	region_id := beego.AppConfig.String("region_id")
	access_key_id := beego.AppConfig.String("access_key_id")
	fmt.Println("access_key_id", access_key_id)
	access_secret := beego.AppConfig.String("access_secret")
	client, err := dysmsapi.NewClientWithAccessKey(region_id, access_key_id, access_secret)
	if err != nil {
		logs.Info(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = beego.AppConfig.String("sign_name")
	request.TemplateCode = beego.AppConfig.String("template_code")
	verify_code_str := fmt.Sprintf("{\"code\":\"%d\"}", verify_code)
	request.TemplateParam = verify_code_str
	response, err := client.SendSms(request)
	if err != nil {
		logs.Info(err.Error())
		return false
	}
	logs.Info("response is %#v\n", response)
	return true
}
