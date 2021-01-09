package api

import (
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
func (this *UserController) ZhifubaoNotify() {
	logs.Info("========zfb_notifyzfb_notifyzfb_notify==========")
	logs.Info(this.Ctx.Input.RequestBody)
	logs.Info("========zfb_notifyzfb_notifyzfb_notify==========")
}