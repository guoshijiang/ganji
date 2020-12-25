package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	"ganji/types/market"
	"github.com/astaxie/beego"
)

type MarketController struct {
	beego.Controller
}


// @Title ServiceContract
// @Description 获取客户联系方式 ServiceContract
// @Success 200 status bool, data interface{}, msg string
// @router /service_contact [post]
func (this *MarketController) ServiceContract() {
	c_ay := market.ContractWayCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &c_ay); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	cs, code, err := models.GetCustomerServicel(c_ay.QueryWay)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, cs, "获取客服信息成功")
	this.ServeJSON()
	return
}


// @Title QsList
// @Description 常见问题列表 QsList
// @Success 200 status bool, data interface{}, msg string
// @router /qs_list [post]
func (this *MarketController) QsList() {
	qs_lp := market.QsListCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &qs_lp); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	var qs_list_ret []market.QsListRet
	qs_list, total, err := models.GetQuestionsList(qs_lp.Page, qs_lp.PageSize)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "数据库错误")
		this.ServeJSON()
		return
	}
	for _, v := range qs_list {
		qs_ := market.QsListRet{
			QsId:v.Id,
			QsAuthor: v.Author,
			QsTitle: v.QsTitle,
			CreateTime: v.CreatedAt,
		}
		qs_list_ret = append(qs_list_ret, qs_)
	}
	data := map[string]interface{}{
		"total":     total,
		"qs_list": qs_list_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取常见问题列表成功")
	this.ServeJSON()
	return
}


// @Title QsDetail
// @Description 常见问题列表 QsDetail
// @Success 200 status bool, data interface{}, msg string
// @router /qs_detail [post]
func (this *MarketController) QsDetail() {
	qs_dtl_p := market.QsDetailCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &qs_dtl_p); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	qs, code, err := models.GetQuestionsDetail(qs_dtl_p.QsId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, "数据库错误")
		this.ServeJSON()
		return
	}
	qs_detail := market.QsDetailRet{
		QsId:qs.Id,
		QsAuthor: qs.Author,
		QsTitle: qs.QsTitle,
		QsDetail: qs.QsDetail,
		CreateTime: qs.CreatedAt,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, qs_detail, "获取常见问题列表成功")
	this.ServeJSON()
	return
}
