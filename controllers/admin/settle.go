package controllers

import (
	"ganji/form_validate"
	"ganji/global/response"
	"ganji/models"
	"ganji/services"
	"github.com/gookit/validate"
	"strings"
	"time"
)

type SettleController struct {
	baseController
}

func (Self *SettleController) Daily() {
	var srv  services.SettleAccountService
	data, pagination := srv.DailyData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination
	Self.Data["super"] = loginUser

	Self.Layout = "public/base.html"
	Self.TplName = "settle/daily.html"
}

func (Self *SettleController) BillSettle() {
	var srv  services.SettleAccountService
	data, pagination := srv.SettleData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination
	Self.Data["super"] = loginUser

	Self.Layout = "public/base.html"
	Self.TplName = "settle/bill.html"
}

func (Self *SettleController) SearchSettle() {
	tm := Self.GetString("create_time")
	var msrv services.MerchantService
	if len(tm) > 0 {
		var (
			srv services.SettleAccountService
			end *time.Time
			data []*models.DailySettleData
		)
		if loginUser.MerchantId == 0 {
			data, _ = srv.SettleDailyData(gQueryParams,end)
			for _, v := range data {
				v.Date = tm
				v.MerchantName = msrv.GetMerchantById(v.MerchantId).MerchantName
			}
			Self.Data["data"] = data
		} else {
			data, _ = srv.SettleDailyData(gQueryParams,end)
			for _, v := range data {
				v.Date = tm
			}
			Self.Data["data"] = data
		}
	}
	Self.Data["super"] = loginUser
	if loginUser.MerchantId == 0 {
		Self.Data["merchants"] = msrv.GetMerchants()
	}

	Self.Layout = "public/base.html"
	Self.TplName = "settle/search.html"
}


func (Self *SettleController) BillSettleCreate() {
	var frm form_validate.SettleForm
	var srv services.SettleAccountService
	if err := Self.ParseForm(&frm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(frm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}
	date := Self.GetString("date")
	dateSlice := strings.Split(date," - ")

	start,_ :=time.ParseInLocation("2006-01-02", dateSlice[0],time.Local)
	end,_ := time.ParseInLocation("2006-01-02", dateSlice[1],time.Local)
	//时间判断
	settles,_ := srv.GetSettleNew(int(frm.MerchantId))
	if len(settles) > 0 {
		if settles[0].EndSettleTime.Unix() > end.Unix() { //已经结算过了
			response.ErrorWithMessage("不能重复结算", Self.Ctx)
		}
	}
	frm.StartSettleTime = start
	frm.EndSettleTime = end
	frm.Status = 0
	insertId := srv.CreateSettle(&frm)

	if insertId > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SettleController) BillSettleUpdate() {
	var frm form_validate.SettleForm
	var srv services.SettleAccountService
	frm.Id,_ = Self.GetInt64("id")
	frm.Status = 1
	num := srv.UpdateSettle(&frm)
	if num > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

//结算确认
func (Self *SettleController) Configure() {
	var srv services.SettleAccountService

	Self.Data["accounts"] = srv.GetList()
	Self.Data["date"] = Self.GetString("date")
	Self.Data["merchant_id"],_ = Self.GetInt64("id")
	Self.Data["amounts"],_ = Self.GetFloat("amounts")

	Self.Layout = "public/base.html"
	Self.TplName = "settle/configure.html"
}