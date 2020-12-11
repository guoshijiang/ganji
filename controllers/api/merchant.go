package api

import "github.com/astaxie/beego"

type MerchantController struct {
	beego.Controller
}

// @Title MerchantList
// @Description 商家列表接口 MerchantList
// @Success 200 status bool, data interface{}, msg string
// @router /marchant_list [post]
func (this *UserController) MerchantList() {

}

