package controllers

import "ganji/services"

type IntegralController struct {
	baseController
}

func (Self *IntegralController) Index() {
	var srv services.IntegralService
	data, pagination := srv.GetPaginateDataRecordRaw(admin["per_page"].(int), gQueryParams)

	Self.Data["data"] = data
	Self.Data["paginate"] = pagination
	Self.Layout = "public/base.html"
	Self.TplName = "integral/record.html"
}

func (Self *IntegralController) Trade() {
	var srv services.IntegralService
	data, pagination := srv.GetPaginateDataTradeRaw(admin["per_page"].(int), gQueryParams)

	Self.Data["data"] = data
	Self.Data["paginate"] = pagination
	Self.Layout = "public/base.html"
	Self.TplName = "integral/trade.html"
}