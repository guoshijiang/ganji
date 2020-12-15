package controllers

import (
	"ganji/form_validate"
	"ganji/global"
	"ganji/global/response"
	"ganji/models"
	"ganji/services"
	"github.com/gookit/validate"
	"log"
	"strconv"
	"strings"
)

type GoodsController struct {
	baseController
}

func (Self *GoodsController) Index() {
	var goodsService services.GoodsServices
	data, pagination := goodsService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "goods/index.html"
}

func (Self *GoodsController) Add() {
	Self.Data["calcway"] = []models.Select{{Id: 0,Name: "件"},{Id: 1,Name: "斤"}}
	Self.Data["cats"] = (&services.GoodsCateService{}).GetGoodsCats()
	Self.Layout = "public/base.html"
	Self.TplName = "goods/add.html"
}


func (Self *GoodsController) Create() {
	var goodsForm form_validate.GoodsForm
	var goodsService services.GoodsServices
	if err := Self.ParseForm(&goodsForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}
	v := validate.Struct(goodsForm)
	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	//上传LOGO
	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "logo")
	if err != nil {
		log.Println("upload--err",err)
	}
	goodsForm.Logo = imgPath

	insertId := goodsService.Create(&goodsForm)
	url := global.URL_BACK

	if goodsForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}
	if insertId > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *GoodsController) Edit() {
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var goodsService services.GoodsServices

	goods := goodsService.GetGoodsById(id)
	if goods == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["calcway"] = []models.Select{{Id: 0,Name: "件"},{Id: 1,Name: "斤"}}
	Self.Data["cats"] = (&services.GoodsCateService{}).GetGoodsCats()
	Self.Data["data"] = goods
	Self.Layout = "public/base.html"
	Self.TplName = "goods/edit.html"
}

func (Self *GoodsController) Update(){
	var goodsForm form_validate.GoodsForm
	if err := Self.ParseForm(&goodsForm); err != nil {
		response.ErrorWithMessage(err.Error(), Self.Ctx)
	}

	if goodsForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", Self.Ctx)
	}

	v := validate.Struct(goodsForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), Self.Ctx)
	}

	//上传LOGO
	imgPath, err := new(services.UploadService).Upload(Self.Ctx, "logo")
	if err != nil {
		log.Println("upload--err",err)
	}
	goodsForm.Logo = imgPath
	//商家验重
	var goodsService services.GoodsServices
	if goodsService.IsExistName(strings.TrimSpace(goodsForm.GoodsName), goodsForm.Id) {
		response.ErrorWithMessage("账号已经存在", Self.Ctx)
	}
	if goodsService.Update(&goodsForm) > 0 {
		response.Success(Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}

func (Self *GoodsController) Del() {
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
	var goodsService services.GoodsServices
	if goodsService.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}
