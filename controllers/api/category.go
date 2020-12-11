package api

import "github.com/astaxie/beego"

type CategoryController struct {
	beego.Controller
}


// @Title CategoryList
// @Description 分类列表接口 CategoryList
// @Success 200 status bool, data interface{}, msg string
// @router /category_list [post]
func (this *UserController) CategoryList() {

}


// @Title GoodsCategoryList
// @Description 分类商品接口 GoodsCategoryList
// @Success 200 status bool, data interface{}, msg string
// @router /goods_category_list [post]
func (this *UserController) GoodsCategoryList() {

}


