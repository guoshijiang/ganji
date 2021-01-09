package api

import (
	"encoding/json"
	"ganji/types"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type NotifyController struct {
	beego.Controller
}


// @Title ZhifubaoNotify
// @Description 支付支付成功回调函数 ZhifubaoNotify
// @Success 200 status bool, data interface{}, msg string
// @router /zfb_notify [post]
func (this *NotifyController) ZhifubaoNotify() {
	logs.Info("========body==========")
	var zhifubao_param interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &zhifubao_param); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err.Error(), "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	logs.Info(zhifubao_param)
	logs.Info("========body==========")
	this.Data["json"] = RetResource(true, 200, nil, "success")
	this.ServeJSON()
	return
}