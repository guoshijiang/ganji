package controllers

import (
	"ganji/form_validate"
	"ganji/global"
	"ganji/global/response"
	"ganji/services"
	"github.com/gookit/validate"
	"log"
	"strconv"
)

type SettleAccountController struct {
	baseController
}

func (Self *SettleAccountController) Index() {
	var srv  services.SettleAccountService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "settle/index.html"
}


func (Self *SettleAccountController) Add() {
	var srv services.MerchantService
	merchants := srv.GetMerchants()
	Self.Data["merchants"] = merchants

	Self.Layout = "public/base.html"
	Self.TplName = "settle/add.html"
}

func (Self *SettleAccountController) Create() {
	var frm form_validate.SettleAccountForm
	var srv services.SettleAccountService
	if err := Self.ParseForm(&frm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(frm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	//上传LOGO
	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "qrcode")
	if err != nil {
		log.Println("upload--err",err)
	}
	frm.Qrcode = imgPath

	insertId := srv.Create(&frm)
	url := global.URL_BACK

	if frm.IsCreate == 1 {
		url = global.URL_RELOAD
	}
	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SettleAccountController) Edit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var srv services.SettleAccountService

	data := srv.GetById(id)
	if data == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}
	var msrv services.MerchantService
	merchants := msrv.GetMerchants()
	Self.Data["merchants"] = merchants

	Self.Data["data"] = data
	Self.Layout = "public/base.html"
	Self.TplName = "settle/edit.html"
}

func (Self *SettleAccountController) Update() {
	var frm form_validate.SettleAccountForm
	if err := Self.ParseForm(&frm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}

	if frm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", Self.Ctx)
	}

	v := validate.Struct(frm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	//上传LOGO
	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "qrcode")
	if err != nil {
		log.Println("upload--err",err)
	}
	frm.Qrcode = imgPath

	var srv services.SettleAccountService

	if srv.Update(&frm) > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SettleAccountController) Delete() {
	idStr := Self.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		Self.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}
	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", Self.Ctx)
	}
	var srv services.SettleAccountService
	if srv.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}
