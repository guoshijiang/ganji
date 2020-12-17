package controllers

import (
	"ganji/global"
	"ganji/global/response"
	"ganji/services"
	"strconv"
)

type OrderController struct {
	baseController
}


//订单列表
func (Self *OrderController) Index() {
	var orderService services.OrderService
	data, pagination := orderService.GetPaginateDataRaw(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "order/index.html"
}

//订单编辑
func (Self *OrderController) Edit() {
	Self.Layout = "public/base.html"
	Self.TplName = "order/edit.html"
}


func (Self *OrderController) Del() {
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
	var orderService services.OrderService
	if orderService.Del(idArr) > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, Self.Ctx)
	} else {
		response.Error(Self.Ctx)
	}
}