package controllers

import "ganji/services"

type SettleController struct {
	baseController
}

func (Self *SettleController) Daily() {
	var srv  services.SettleAccountService
	gQueryParams.Add("t0.valid_order_num",Self.GetString("_keywords"))
	data, pagination := srv.DailyData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "settle/daily.html"
}
