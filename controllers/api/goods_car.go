package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_goods_car "ganji/types/goods_car"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type GoodsCarController struct {
	beego.Controller
}

// @Title AddGoodsToCar
// @Description 将商品添加到购物车 AddGoodsToCar
// @Success 200 status bool, data interface{}, msg string
// @router /add_goods_to_car [post]
func (this *GoodsCarController) AddGoodsToCar() {
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
	//if goods_car.PayAmount != goods_dtl.GoodsPrice && goods_car.PayAmount != goods_dtl.GoodsDisPrice {
	//	this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "传人的商品价格和系统内部的商品价格不一样")
	//	this.ServeJSON()
	//	return
	//}
	if goods_car.IsDis == 0 { //非打折商品
		if goods_car.PayAmount != goods_dtl.GoodsPrice * float64(goods_car.BuyNums ){
			this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "无效的商品价格")
			this.ServeJSON()
			return
		}
	} else if goods_car.IsDis == 1 { //打折商品
		if goods_car.PayAmount != goods_dtl.GoodsDisPrice * float64(goods_car.BuyNums ){
			this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "无效的商品价格")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "无效的商品价格方式")
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
func (this *GoodsCarController) EditGoodsCar() {
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
	goods_car_edit := type_goods_car.EditGoodCarCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_car_edit); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_car_edit.EditGoodCarCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	if user_token.Id != goods_car_edit.UserId {
		this.Data["json"] = RetResource(false, types.UserTokenUserIdNotEqual, nil, "传入的用户ID和用户Token不相符")
		this.ServeJSON()
		return
	}
	goods_dtl, _, _ := models.GetGoodsDetail(goods_car_edit.GoodsId)
	//if goods_car_edit.PayAmount != goods_dtl.GoodsPrice && goods_car_edit.PayAmount != goods_dtl.GoodsDisPrice {
	//	this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "传人的商品价格和系统内部的商品价格不一样")
	//	this.ServeJSON()
	//	return
	//}
	goods_car, code, err :=  models.GetGoodsCarDetail(goods_car_edit.GoodsCarId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	if goods_car_edit.IsDis == 0 { //非打折商品
		if goods_car_edit.PayAmount != goods_dtl.GoodsPrice * float64(goods_car_edit.BuyNums ){
			this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "无效的商品价格")
			this.ServeJSON()
			return
		}
	} else if goods_car_edit.IsDis == 1 { //打折商品
		if goods_car_edit.PayAmount != goods_dtl.GoodsDisPrice * float64(goods_car_edit.BuyNums ){
			this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, nil, "无效的商品价格")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "无效的商品价格方式")
		this.ServeJSON()
		return
	}
	goods_car.BuyNums = goods_car.BuyNums + goods_car_edit.BuyNums
	goods_car.PayAmount = goods_car.PayAmount + goods_car_edit.PayAmount
	err = goods_car.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库操作错误")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "编辑购物车成功")
	this.ServeJSON()
	return
}


// @Title DelGoodsToCar
// @Description 删除购物车 DelGoodsToCar
// @Success 200 status bool, data interface{}, msg string
// @router /del_goods_car [post]
func (this *GoodsCarController) DelGoodsCar() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	goods_car_del := type_goods_car.DelGoodCarCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_car_del); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_car_del.DelGoodCarCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	str_list := strings.Split(goods_car_del.GoodsIds, ",")
	for _, v := range str_list {
		gcid, _ := strconv.Atoi(v)
		gcr, _, _ := models.GetGoodsCarDetail(int64(gcid))
		err = gcr.Delete()
		if err != nil {
			this.Data["json"] = RetResource(true, types.SystemDbErr, nil, "删除购物车操作数据库失败")
			this.ServeJSON()
			return
		}
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "删除购物车成功")
	this.ServeJSON()
	return
}


// @Title GoodsCarList
// @Description 获取购物车列表 GoodsCarList
// @Success 200 status bool, data interface{}, msg string
// @router /goods_car_list [post]
func (this *GoodsCarController) GoodsCarList() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	ut, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	goods_car_lst := type_goods_car.GoodCarListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_car_lst); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_car_lst.GoodCarListCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	goods_car_list, total, err := models.GetGoodsCarList(goods_car_lst.Page, goods_car_lst.PageSize, ut.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsCarListFail, nil, err.Error())
		this.ServeJSON()
		return
	}
	var gds_car_lst []type_goods_car.GoodsCarList
	for _, value := range goods_car_list {
		gds, _, _ :=models.GetGoodsDetail(value.GoodsId)
		mct, _, _ := models.GetMerchantDetail(gds.MerchantId)
		var goods_price  float64
		if gds.IsDiscount == 0 {
			goods_price = gds.GoodsPrice
		} else {
			goods_price = gds.GoodsDisPrice
		}
		gdsc := type_goods_car.GoodsCarList{
			MerchantId: gds.MerchantId,
			MerchantName:mct.MerchantName,
			GoodsCarId: value.Id,
			GoodsId: gds.Id,
			GoodsLogo: value.Logo,
			GoodsTitle: gds.Title,
			GoodsMark: gds.GoodsMark,
			GoodsName: gds.GoodsName,
			GoodsPrice: goods_price,
			UserId: value.UserId,
			BuyNums: value.BuyNums,
			PayAmount: value.PayAmount,
		}
		gds_car_lst = append(gds_car_lst, gdsc)
	}
	data := map[string]interface{}{
		"total": total,
		"gds_car_lst": gds_car_lst,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取购物车列表成功")
	this.ServeJSON()
	return
}
