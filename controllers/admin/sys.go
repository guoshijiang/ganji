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

type SysController struct {
	baseController
}

func (Self *SysController) BannerIndex() {
	var srv services.BannerService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "banner/index.html"
}

func (Self *SysController) BannerAdd() {
	Self.Layout = "public/base.html"
	Self.TplName = "banner/add.html"
}


func (Self *SysController) BannerCreate() {
	var vForm form_validate.BannerForm
	var srv services.BannerService
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(vForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "avator")
	if err != nil {
		log.Println("upload--err",err)
	}
	vForm.Avator = imgPath
	insertId := srv.Create(&vForm)
	url := global.URL_BACK

	if vForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}
	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SysController) BannerEdit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var srv services.BannerService

	banner := srv.GetBannerById(id)
	if banner == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["data"] = banner
	Self.Layout = "public/base.html"
	Self.TplName = "banner/edit.html"
}

func (Self *SysController) BannerUpdate(){
	var (
		vForm form_validate.BannerForm
		srv services.BannerService
	)
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}

	if vForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", Self.Ctx)
	}

	v := validate.Struct(vForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "avator")
	if err != nil {
		log.Println("upload--err",err.Error())
	}
	vForm.Avator = imgPath
	if srv.Update(&vForm) > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SysController) BannerDel() {
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
	var srv services.BannerService
	if srv.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}


func (Self *SysController) VerIndex() {
	var srv services.VersionService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "version/index.html"
}

func (Self *SysController) VerAdd() {
	Self.Layout = "public/base.html"
	Self.TplName = "version/add.html"
}


func (Self *SysController) VerCreate() {
	var (
		vForm form_validate.VersionForm
		srv services.VersionService
	)
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(vForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	insertId := srv.Create(&vForm)
	url := global.URL_BACK

	if vForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}
	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SysController) VerEdit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var srv services.VersionService

	ver := srv.GetVersionById(id)
	if ver == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["data"] = ver
	Self.Layout = "public/base.html"
	Self.TplName = "version/edit.html"
}

func (Self *SysController) VerUpdate(){
	var (
		vForm form_validate.VersionForm
		srv services.VersionService
	)
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	if vForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", Self.Ctx)
	}

	v := validate.Struct(vForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}
	if srv.Update(&vForm) > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *SysController) VerDel() {
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
	var srv services.VersionService
	if srv.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

//充值记录
func (Self *SysController) WalletRecord() {
	var srv services.SysService
	data, pagination := srv.GetPaginateDataWalletRecordRaw(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "user/wallet_record.html"
}
