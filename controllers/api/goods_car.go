package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_goods_car "ganji/types/goods_car"
	"github.com/astaxie/beego"
	"strings"
)

type GoodsCarController struct {
	beego.Controller
}

// @Title AddGoodsToCar
// @Description 将商品添加到购物车 AddGoodsToCar
// @Success 200 status bool, data interface{}, msg string
// @router /add_goods_to_car [post]
func (this *UserController) AddGoodsToCar() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_token, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	goods_car := type_goods_car.AddGoodCarCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_car); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_car.AddGoodCarCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	if user_token.Id != goods_car.UserId {
		this.Data["json"] = RetResource(false, types.UserTokenUserIdNotEqual, nil, "传入的用户ID和用户Token不相符")
		this.ServeJSON()
		return
	}
	goods_dtl, _, _ := models.GetGoodsDetail(goods_car.GoodsId)
	if goods_car.PayAmount != goods_dtl.GoodsPrice && goods_car.PayAmount != goods_dtl.GoodsDisPrice {
		this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "传人的商品价格和系统内部的商品价格不一样")
		this.ServeJSON()
		return
	}
	if goods_dtl != nil {
		gdc := models.GoodsCar {
			GoodsId: goods_dtl.Id,
			Logo: goods_dtl.Logo,
			GoodsTitle: goods_dtl.Title,
			GoodsName: goods_dtl.GoodsName,
			UserId: user_token.Id,
			BuyNums: goods_car.BuyNums,
			PayAmount: goods_car.PayAmount,
		}
		err := gdc.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库操作错误")
			this.ServeJSON()
			return
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "添加购物车成功")
		this.ServeJSON()
		return
	}
}


// @Title EditGoodsToCar
// @Description 编辑购物车 EditGoodsToCar
// @Success 200 status bool, data interface{}, msg string
// @router /edit_goods_car [post]
func (this *UserController) EditGoodsCar() {

}


// @Title DelGoodsToCar
// @Description 删除购物车 DelGoodsToCar
// @Success 200 status bool, data interface{}, msg string
// @router /del_goods_car [post]
func (this *UserController) DelGoodsCar() {

}


// @Title GoodsCarList
// @Description 获取购物车列表 GoodsCarList
// @Success 200 status bool, data interface{}, msg string
// @router /goods_car_list [post]
func (this *UserController) GoodsCarList() {

}

