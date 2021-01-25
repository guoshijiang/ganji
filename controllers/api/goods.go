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

const (
	COLOR = 0
	SIZE = 1
	OTHER = 2
)

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
	image_path := beego.AppConfig.String("img_root_path")
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range good_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:   value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: image_path + value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			SendIntegral: value.SendIntegral,
			LeftTime: value.LeftTime,
			IsHot: value.IsHot,
			IsDiscount: value.IsDiscount,
			IsIgSend: value.IsIgSend,
			IsGroup: value.IsGroup,
			IsIntegral: value.IsIntegral,
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
	img_path := beego.AppConfig.String("img_root_path")
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range goods_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:   value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: img_path + value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			SendIntegral: value.SendIntegral,
			LeftTime: value.LeftTime,
			IsHot: value.IsHot,
			IsDiscount: value.IsDiscount,
			IsIgSend: value.IsIgSend,
			IsGroup: value.IsGroup,
			IsIntegral: value.IsIntegral,
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


// @Title GetLimitTimeGoodsList
// @Description 限时购买产品列表 GetLimitTimeGoodsList
// @Success 200 status bool, data interface{}, msg string
// @router /limit_time_goods_list [post]
func (this *GoodsController) GetLimitTimeGoodsList() {
	lt_gds := type_goods.LTGoodsListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &lt_gds); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	goods_list, total, err := models.GetLtGoodsList(lt_gds.Page, lt_gds.PageSize)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, "获取商品列表失败")
		this.ServeJSON()
		return
	}
	img_path := beego.AppConfig.String("img_root_path")
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range goods_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:   value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: img_path + value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			SendIntegral: value.SendIntegral,
			LeftTime: value.LeftTime,
			IsHot: value.IsHot,
			IsDiscount: value.IsDiscount,
			IsIgSend: value.IsIgSend,
			IsGroup: value.IsGroup,
			IsIntegral: value.IsIntegral,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"goods_lst": goods_ret_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取限时购买商品列表成功")
	this.ServeJSON()
	return
}


// @Title GetHotGoodsList
// @Description 获取爆款产品列表 GetHotGoodsList
// @Success 200 status bool, data interface{}, msg string
// @router /hot_goods_list [post]
func (this *GoodsController) GetHotGoodsList() {
	lt_gds := type_goods.LTGoodsListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &lt_gds); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	goods_list, total, err := models.GetOrderDownHotGoodsList(lt_gds.Page, lt_gds.PageSize)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, "获取商品列表失败")
		this.ServeJSON()
		return
	}
	img_path := beego.AppConfig.String("img_root_path")
	var goods_ret_list []type_goods.CategoryGoodsRet
	for _, value := range goods_list {
		gds_ret := type_goods.CategoryGoodsRet{
			GoodsId:   value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: img_path + value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			SendIntegral: value.SendIntegral,
			LeftTime: value.LeftTime,
			IsHot: value.IsHot,
			IsDiscount: value.IsDiscount,
			IsIgSend: value.IsIgSend,
			IsGroup: value.IsGroup,
			IsIntegral: value.IsIntegral,
		}
		goods_ret_list = append(goods_ret_list, gds_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"goods_lst": goods_ret_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取获取爆款产品列表成功")
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
		user_addr, _, err := models.GetUserAddressDefault(goods_detil.UserId)
		if err != nil {
			user_address = nil
		} else {
			user_address["address_id"] = user_addr.Id
			user_address["address_name"] = user_addr.Address
		}
	} else {
		user_address = nil
	}
	color_type_list, _, err := models.GetGoodsTypeList(goods_dtl.Id, COLOR)
	size_type_list, _, err := models.GetGoodsTypeList(goods_dtl.Id, SIZE)
	other_type_list, _, err := models.GetGoodsTypeList(goods_dtl.Id, OTHER)
	var gds_color_type_list, gds_size_type_list, gds_other_type_list []type_goods.GoodsTypeRet
	// 颜色属性
	if err != nil || color_type_list == nil {
		gds_color_type_list = nil
	} else {
		for _, value_t := range color_type_list {
			c_gds_type := type_goods.GoodsTypeRet{
				GoodsTypeId: value_t.Id,
				GoodsTypeName: value_t.TypeName,
			}
			gds_color_type_list = append(gds_color_type_list, c_gds_type)
		}
	}
	// 大小属性
	if err != nil || size_type_list == nil {
		gds_size_type_list = nil
	} else {
		for _, value_t := range size_type_list {
			s_gds_type := type_goods.GoodsTypeRet{
				GoodsTypeId: value_t.Id,
				GoodsTypeName: value_t.TypeName,
			}
			gds_size_type_list = append(gds_size_type_list, s_gds_type)
		}
	}
	// 其他属性
	if err != nil || other_type_list == nil {
		gds_other_type_list = nil
	} else {
		for _, value_0_t := range other_type_list {
			o_gds_type := type_goods.GoodsTypeRet{
				GoodsTypeId: value_0_t.Id,
				GoodsTypeName: value_0_t.TypeName,
			}
			gds_other_type_list = append(gds_other_type_list, o_gds_type)
		}
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
		"goods_integral": goods_dtl.GoodsIntegral,
		"group_number": goods_dtl.GroupNumber,
		"send_interal": goods_dtl.SendIntegral,
		"goods_name": goods_dtl.GoodsName,
		"goods_params": goods_dtl.GoodsParams,
		"goods_detail": goods_dtl.GoodsDetail,
		"goods_img": gds_img_lst,
		"user_address": user_address,
		"merchant_info": merchant_info,
		"is_hot": goods_dtl.IsHot,
		"is_discount": goods_dtl.IsDiscount,
		"is_ig_send": goods_dtl.IsIgSend,
		"is_group": goods_dtl.IsGroup,
		"is_integral": goods_dtl.IsIntegral,
		"color_types": gds_color_type_list,
		"size_types": gds_size_type_list,
		"other_types": gds_other_type_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, goods_detail, "获取商品详情成功")
	this.ServeJSON()
	return
}

