package api

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

// @Title AppIndex
// @Description 首页接口 AppIndex
// @Success 200 status bool, data interface{}, msg string
// @router /index_list [post]
func (this *IndexController) AppIndex() {

}

