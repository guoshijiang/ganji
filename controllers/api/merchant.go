package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_merchant "ganji/types/merchant"
	"github.com/astaxie/beego"
)

type MerchantController struct {
	beego.Controller
}


// @Title MerchantList
// @Description 商家列表接口 MerchantList
// @Success 200 status bool, data interface{}, msg string
// @router /marchant_list [post]
func (this *MerchantController) MerchantList() {
	gds_merchant := type_merchant.MerchantListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &gds_merchant); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := gds_merchant.MerchantListCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	merchant_list, total, err := models.GetMerchantList(gds_merchant.Page, gds_merchant.PageSize, gds_merchant.MerchantName, gds_merchant.MerchantAddress)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetMerchantListFail, nil, "获取商家列表失败")
		this.ServeJSON()
		return
	}
	image_path := beego.AppConfig.String("img_root_path")
	var mct_list_ret []type_merchant.MerchantListRet
	for _, merchant := range merchant_list {
		mct_ret := type_merchant.MerchantListRet{
			MctId: merchant.Id,
			MctName: merchant.MerchantName,
			MctIntroduce: merchant.MerchantIntro,
			MctLogo: image_path + merchant.Logo,
			MctWay: merchant.MerchantWay,
			ShopLevel: merchant.ShopLevel,
			ShopServer: merchant.ShopServer,
		}
		mct_list_ret = append(mct_list_ret, mct_ret)
	}
	data := map[string]interface{}{
		"total":     total,
		"mct_lst": mct_list_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取商家列表成功")
	this.ServeJSON()
	return
}


// @Title MerchantDetail
// @Description 商家详情接口 MerchantDetail
// @Success 200 status bool, data interface{}, msg string
// @router /marchant_detail [post]
func (this *MerchantController) MerchantDetail() {
	merchant_dtil := type_merchant.MerchantDetailCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &merchant_dtil); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := merchant_dtil.MerchantDetailCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	image_path := beego.AppConfig.String("img_root_path")
	mcrt_detail, code, err := models.GetMerchantDetail(merchant_dtil.MerchantId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	m_goods_nums := models.GetMerchantGoodsNums(merchant_dtil.MerchantId)
	mct_ret_dtl := type_merchant.MerchantDetailRet{
		MctId: mcrt_detail.Id,
		MctLogo: image_path + mcrt_detail.Logo,
		MctName: mcrt_detail.MerchantName,
		MctIntroduce: mcrt_detail.MerchantIntro,
		MerchantDetail: mcrt_detail.MerchantDetail,
		Address: mcrt_detail.Address,
		GoodsNum: m_goods_nums,
		MctWay: mcrt_detail.MerchantWay,
		ShopLevel: mcrt_detail.ShopLevel,
		ShopServer: mcrt_detail.ShopServer,
		CreatedAt:mcrt_detail.CreatedAt,
		UpdatedAt: mcrt_detail.UpdatedAt,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, mct_ret_dtl, "获取商家详情成功")
	this.ServeJSON()
	return
}
