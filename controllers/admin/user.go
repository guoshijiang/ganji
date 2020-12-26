package controllers

import (
	"ganji/form_validate"
	"ganji/global"
	"ganji/global/response"
	"ganji/models"
	"ganji/services"
	"github.com/gookit/validate"
	"strconv"
)

type UserController struct {
	baseController
}

func (Self *UserController) Index() {
	var srv services.UserService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "user/index.html"
}

func (Self *UserController) Add() {
	Self.Data["levels"] = []models.Select{{Id:0,Name: "普通会员"},{Id: 1,Name: "白银会员"},{Id: 2,Name: "白金会员"},{Id: 3,Name: "黄金会员"},{Id: 4,Name: "砖石会有"},{Id: 5,Name: "皇冠会员"}}

	Self.Layout = "public/base.html"
	Self.TplName = "user/add.html"
}

//用户创建页面
func (Self *UserController) Create() {
	var vForm form_validate.UserForm
	var srv services.UserService
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(srv)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	var user models.User
	ok := user.ExistByPhone(vForm.Phone)
	if ok {
		response.ErrorWithMessage("号码已存在", Self.Ctx)
	}

	if srv.IsExistName(vForm.UserName,0) {
		response.ErrorWithMessage("用户名已存在", Self.Ctx)
	}

	url := global.URL_BACK
	if vForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	vForm.Avator,_ = new(services.UploadService).Upload(Self.Ctx,"avator")

	if err := srv.Create(&vForm);err == nil {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

//用户编辑页面
func (Self *UserController) Edit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var srv services.UserService

	user := srv.GetUserById(id)
	if user == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["levels"] = []models.Select{{Id:0,Name: "普通会员"},{Id: 1,Name: "白银会员"},{Id: 2,Name: "白金会员"},{Id: 3,Name: "黄金会员"},{Id: 4,Name: "砖石会有"},{Id: 5,Name: "皇冠会员"}}
	Self.Data["data"] = user
	Self.Layout = "public/base.html"
	Self.TplName = "user/edit.html"
}

//用户更新
func (Self *UserController) Update(){
	var (
		vForm 	form_validate.UserForm
		srv 	services.UserService
	)
	if err := Self.ParseForm(&vForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}

	vForm.Avator,_ = new(services.UploadService).Upload(Self.Ctx,"avator")
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

//用户删除
func (Self *UserController) Del() {
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
	var srv services.UserService
	if srv.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *UserController) Wallet() {
	var srv services.UserService
	id, _ := Self.GetInt("id", -1)
	if id < 0 {
		response.ErrorWithMessage("用户不存在",Self.Ctx)
	}
	gQueryParams.Set("user_id",strconv.Itoa(id))
	data, pagination := srv.GetPaginateDataWalletRaw(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "user/wallet.html"
}

func (Self *UserController) Integral() {
	var srv services.UserService
	id, _ := Self.GetInt("id", -1)
	if id < 0 {
		response.ErrorWithMessage("用户不存在",Self.Ctx)
	}
	gQueryParams.Set("user_id",strconv.Itoa(id))
	data, pagination := srv.GetPaginateDataIntegralRaw(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "user/integral.html"
}

func (Self *UserController) Address() {
	var srv services.UserService
	id, _ := Self.GetInt("id", -1)
	if id < 0 {
		response.ErrorWithMessage("用户不存在",Self.Ctx)
	}
	gQueryParams.Set("_user_id",strconv.Itoa(id))
	data, pagination := srv.GetPaginateAddressData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "user/address.html"
}


func (Self *UserController) Coupon() {
	var srv services.UserService
	id, _ := Self.GetInt("id", -1)
	if id < 0 {
		response.ErrorWithMessage("用户不存在",Self.Ctx)
	}
	gQueryParams.Set("user_id",strconv.Itoa(id))
	data, pagination := srv.GetPaginateDataCouponRaw(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "user/coupon.html"
}