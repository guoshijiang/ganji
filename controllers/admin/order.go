package controllers

import (
	"ganji/global"
	"ganji/global/response"
	"ganji/models"
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
	id, _ := Self.GetInt64("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", Self.Ctx)
	}
	var srv services.OrderService

	data := srv.GetOrderById(id)
	if data == nil {
		response.ErrorWithMessage("Not Found Info By Id.", Self.Ctx)
	}

	Self.Data["data"] = data
	Self.Layout = "public/base.html"
	Self.TplName = "order/edit.html"
}



func (Self *OrderController) Update() {
	id,_ := Self.GetInt64("id",-1)
	ship_number := Self.GetString("ship_number","")
	if id < 0 {
		response.ErrorWithMessage("订单不存在",Self.Ctx)
	}
	order := models.GoodsOrder{Id: id,ShipNumber: ship_number}
	if new(services.OrderService).UpdateShipNumber(&order) > 0 {
		response.Success(Self.Ctx)
	}
	response.Error(Self.Ctx)
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