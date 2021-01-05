package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_index "ganji/types/index"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}


// @Title AppIndexUp
// @Description 首页上面的商品列表接口 AppIndexUp
// @Success 200 status bool, data interface{}, msg string
// @router /index_up_list [post]
func (this *IndexController) AppIndexUp() {
	var banner_ret_list []type_index.IndexBannerRet
	var cat_ret_list []type_index.IndexCatRet
	var limit_buy_list []type_index.IndexLimitTimeGoodsRet
	var hot_sell_list, best_goods_list []type_index.IndexGoodsBuyRet
    image_path := beego.AppConfig.String("img_root_path")
	banner_list, code, err := models.GetBannerList()
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	for _, banner := range banner_list {
		banner_ret := type_index.IndexBannerRet{
			BannerId: banner.Id,
			BannerImg: image_path + banner.Avator,
			BannerUrl: banner.Url,
		}
		banner_ret_list = append(banner_ret_list, banner_ret)
	}
	cat_list, code, err := models.GetOneLevelCategoryList()
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	for _, cat := range cat_list {
		cat_ret := type_index.IndexCatRet{
			CatId: cat.Id,
			CatName: cat.Name,
			CatIcon: image_path + cat.Icon,
		}
		cat_ret_list = append(cat_ret_list, cat_ret)
	}
	time_goods_list, code, err := models.GetLimitTimeGoodsList()
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	for _, gd := range time_goods_list {
		time_gd_ret := type_index.IndexLimitTimeGoodsRet{
			GoodsId: gd.Id,
			GoodsMark: gd.GoodsMark,
			Title: gd.Title,
			Logo: image_path + gd.Logo,
			GoodsPrice: gd.GoodsPrice,
			GoodsDisPrice: gd.GoodsDisPrice,
			GoodsIntegral: gd.GoodsIntegral,
			LeftTime: gd.LeftTime,
			IsDiscount: gd.IsDiscount,
			IsIgSend: gd.IsIgSend,
			IsGroup: gd.IsGroup,
			IsIntegral: gd.IsIntegral,
		}
		limit_buy_list = append(limit_buy_list, time_gd_ret)
	}

	hot_gds_list, code, err := models.GetHotGoodsList()
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	for _, hgd := range hot_gds_list {
		hot_gd_ret := type_index.IndexGoodsBuyRet{
			GoodsId: hgd.Id,
			GoodsMark: hgd.GoodsMark,
			Title: hgd.Title,
			Logo: image_path + hgd.Logo,
			GoodsPrice: hgd.GoodsPrice,
			GoodsDisPrice: hgd.GoodsDisPrice,
			GoodsIntegral: hgd.GoodsIntegral,
			IsDiscount: hgd.IsDiscount,
			IsIgSend: hgd.IsIgSend,
			IsGroup: hgd.IsGroup,
			IsIntegral: hgd.IsIntegral,
		}
		hot_sell_list = append(hot_sell_list, hot_gd_ret)
	}

	bst_gds_list, code, err := models.GetDiscountGoodsList()
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	for _, bgd := range bst_gds_list {
		bst_gd_ret := type_index.IndexGoodsBuyRet{
			GoodsId: bgd.Id,
			GoodsMark: bgd.GoodsMark,
			Title: bgd.Title,
			Logo: image_path + bgd.Logo,
			GoodsPrice: bgd.GoodsPrice,
			GoodsDisPrice: bgd.GoodsDisPrice,
			IsDiscount: bgd.IsDiscount,
			IsIgSend: bgd.IsIgSend,
			IsGroup: bgd.IsGroup,
			IsIntegral: bgd.IsIntegral,
		}
		best_goods_list = append(best_goods_list, bst_gd_ret)
	}
	data := map[string]interface{}{
		"banner": banner_ret_list,
		"category": cat_ret_list,
		"limit_time_buy": limit_buy_list,
		"hot_sell": hot_sell_list,
		"best_goods": best_goods_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取商品成功")
	this.ServeJSON()
	return
}


// @Title AppIndexDown
// @Description 首页上面的商品列表接口 AppIndexDown
// @Success 200 status bool, data interface{}, msg string
// @router /index_down_list [post]
func (this *IndexController) AppIndexDown() {
	index_down_check := type_index.IndexDownCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &index_down_check); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err.Error(), "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := index_down_check.IndexDownCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	goods_list, total, err := models.GetIndexDownGoodsList(index_down_check.Page, index_down_check.PageSize, index_down_check.IndexCatId)
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetGoodsListFail, nil, err.Error())
		this.ServeJSON()
		return
	}
	image_path := beego.AppConfig.String("img_root_path")
	var goods_list_down []type_index.IndexDownGoodsListRet
	for _, value := range goods_list {
		gds_down := type_index.IndexDownGoodsListRet{
			GoodsId: value.Id,
			GoodsMark: value.GoodsMark,
			Title: value.Title,
			Logo: image_path + value.Logo,
			GoodsPrice: value.GoodsPrice,
			GoodsDisPrice: value.GoodsDisPrice,
			GoodsIntegral: value.GoodsIntegral,
			IsDiscount: value.IsDiscount,
			IsIgSend: value.IsIgSend,
			IsGroup: value.IsGroup,
			IsIntegral: value.IsIntegral,
		}
		goods_list_down = append(goods_list_down, gds_down)
	}
	data := map[string]interface{}{
		"total":     total,
		"gds_lst": goods_list_down,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取商品首页下方商品列表成功")
	this.ServeJSON()
	return
}