package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_goods "ganji/types/goods"
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}


// @Title GoodsCategoryList
// @Description 分类商品列表接口 GoodsCategoryList
// @Success 200 status bool, data interface{}, msg string
// @router /goods_category_list [post]
func (this *GoodsController) GoodsCategoryList() {
	goods_category := type_goods.GoodsCategoryCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_category); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_category.GoodsCategoryCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	good_list, total, err := models.GetCategoryGoodsList(goods_category.Page, goods_category.PageSize, goods_category.FirstLevetCatId, goods_category.LastLevelCatId)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, "获取商品列表失败")
		this.ServeJSON()
		return
	}
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range good_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:   value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			LeftTime: value.LeftTime,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"gds_lst":   goods_ret_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取分类商品列表成功")
	this.ServeJSON()
	return
}


// @Title MerchantGoodsList
// @Description 商家商品列表接口 MerchantGoodsList
// @Success 200 status bool, data interface{}, msg string
// @router /merchant_goods_list [post]
func (this *GoodsController) MerchantGoodsList() {
	merchant_gds := type_goods.MerchantGoodsListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &merchant_gds); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := merchant_gds.MerchantGoodsListCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	goods_list, total, err := models.GetMerchantGoodsList(merchant_gds.Page, merchant_gds.PageSize, merchant_gds.MerchantId, merchant_gds.QueryWay)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, "获取商品列表失败")
		this.ServeJSON()
		return
	}
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range goods_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:   value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			LeftTime: value.LeftTime,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"order_lst": goods_ret_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取分类商品列表成功")
	this.ServeJSON()
	return
}


// @Title GoodsDetail
// @Description 商品详情接口 GoodsDetail
// @Success 200 status bool, data interface{}, msg string
// @router /goods_detail [post]
func (this *GoodsController) GoodsDetail() {
	goods_detil := type_goods.GoodsDetailCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &goods_detil); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := goods_detil.GoodsDetailCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	goods_dtl, code, err := models.GetGoodsDetail(goods_detil.GoodsId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商品列表失败")
		this.ServeJSON()
		return
	}
	img_path := beego.AppConfig.String("img_root_path")
	merchant, code, err := models.GetMerchantDetail(goods_dtl.MerchantId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商家信息失败")
		this.ServeJSON()
		return
	}
	merchant_info :=  map[string]interface{}{
		"merchant_id": merchant.Id,
		"merchant_logo": img_path + merchant.Logo,
		"merchant_name": merchant.MerchantName,
	}
	goods_img_lst, code, err := models.GetGoodsImgList(goods_dtl.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "获取商品图片失败")
		this.ServeJSON()
		return
	}
	gds_img_lst := []type_goods.GoodsImagesRet{}
	for _, v := range goods_img_lst {
		gds_img := type_goods.GoodsImagesRet{
			GoodsImgId:v.Id,
			ImageUrl: img_path + v.Image,
		}
		gds_img_lst = append(gds_img_lst, gds_img)
	}
	user_address := make(map[string]interface{})
	if goods_detil.UserId > 0 {
		user_addr, code, err := models.GetUserAddressDefault(goods_detil.UserId)
		if err != nil {
			this.Data["json"] = RetResource(false, code, err.Error(), "获取用户默认地址失败失败")
			this.ServeJSON()
			return
		}
		user_address["address_id"] = user_addr.Id
		user_address["address_name"] = user_addr.Address
	} else {
		user_address = nil
	}
	goods_detail := map[string]interface{}{
		"id": goods_dtl.Id,
		"title": goods_dtl.Title,
		"mark": goods_dtl.GoodsMark,
		"logo": img_path + goods_dtl.Logo,
		"serveice": goods_dtl.Serveice,
		"calc_way": goods_dtl.CalcWay,
		"sell_nums": goods_dtl.SellNums,
		"total_amount": goods_dtl.TotalAmount,
		"left_amount": goods_dtl.LeftAmount,
		"goods_price": goods_dtl.GoodsPrice,
		"goods_dis_price": goods_dtl.GoodsDisPrice,
		"goods_name": goods_dtl.GoodsName,
		"goods_params": goods_dtl.GoodsParams,
		"goods_detail": goods_dtl.GoodsDetail,
		"goods_img": gds_img_lst,
		"user_address": user_address,
		"merchant_info": merchant_info,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, goods_detail, "获取商品详情成功")
	this.ServeJSON()
	return
}


