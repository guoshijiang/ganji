package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_order "ganji/types/order"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type OrderController struct {
	beego.Controller
}


// @Title CreateOrder finished
// @Description 直接创建订单 CreateOrder
// @Success 200 status bool, data interface{}, msg string
// @router /create_order [post]
func (this *OrderController) CreateOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	requestUser, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var create_order type_order.CreateOrderCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &create_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := create_order.CreateOrderCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != create_order.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	gds, _, _ := models.GetGoodsDetail(create_order.GoodsId)
	order_nmb := uuid.NewV4()
	if create_order.IsDis == 0 {  // 不打折
		if gds.GoodsPrice * float64(create_order.BuyNums) != create_order.PayAmount {
			this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, err, "无效的商品价格")
			this.ServeJSON()
			return
		}
	} else if create_order.IsDis == 1 { //打折活动产品
		if gds.GoodsDisPrice * float64(create_order.BuyNums) != create_order.PayAmount {
			this.Data["json"] = RetResource(false, types.InvalidGoodsPirce, err, "无效的商品价格")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(false, types.InvalidVerifyWay, err, "无效的验证方式")
		this.ServeJSON()
		return
	}
	cmt := models.GoodsOrder{
		GoodsId: gds.Id,
		MerchantId: gds.MerchantId,
		AddressId: create_order.AddressId,
		GoodsTitle: gds.Title,
		GoodsName: gds.GoodsName,
		Logo: gds.Logo,
		UserId: create_order.UserId,
		BuyNums: create_order.BuyNums,
		PayWay: 4,
		PayAmount: create_order.PayAmount,
		SendIntegral: create_order.SendIntegral,
		OrderNumber: order_nmb.String(),
		OrderStatus: 0,
		FailureReason: "未支付",
	}
	err, id := cmt.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "创建订单失败")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, map[string]interface{}{"id": id}, "创建订单成功")
		this.ServeJSON()
		return
	}
}


// @Title CreateOrderByGoodsCar finished
// @Description 通过购物车 IDS 创建订单 CreateOrderByGoodsCar
// @Success 200 status bool, data interface{}, msg string
// @router /create_order_by_gdscar [post]
func (this *OrderController) CreateOrderByGoodsCar() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	requestUser, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var create_order_gdscar type_order.CreateOrderGoodsCarCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &create_order_gdscar); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := create_order_gdscar.CreateOrderGoodsCarCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != create_order_gdscar.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	ids_list := create_order_gdscar.GoodsCarIds
	var order_ids []int64
	var total_pay_amount float64
	total_pay_amount = 0
	for i := 0; i < len(ids_list); i++ {
		gds_car_dtl, code, err := models.GetGoodsCarDetail(ids_list[i])
		if err != nil {
			this.Data["json"] = RetResource(false, code, err.Error(), "获取购物车详细信息失败")
			this.ServeJSON()
			return
		}
		goods_dtl, _, _ := models.GetGoodsDetail(gds_car_dtl.GoodsId)
		order_nmb := uuid.NewV4()
		send_integral := gds_car_dtl.PayAmount
		total_pay_amount = total_pay_amount + gds_car_dtl.PayAmount
		cmt := models.GoodsOrder{
			GoodsId: goods_dtl.Id,
			MerchantId: goods_dtl.MerchantId,
			AddressId: gds_car_dtl.AddresId,
			GoodsTitle: goods_dtl.Title,
			GoodsName: goods_dtl.GoodsName,
			Logo: goods_dtl.Logo,
			UserId: create_order_gdscar.UserId,
			BuyNums: gds_car_dtl.BuyNums,
			PayWay: 4,
			PayAmount: gds_car_dtl.PayAmount,
			SendIntegral: send_integral,
			OrderNumber: order_nmb.String(),
			OrderStatus: 0,
			FailureReason: "未支付",
		}
		err, id := cmt.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "创建订单失败")
			this.ServeJSON()
			return
		} else {
			err = gds_car_dtl.Delete()
			if err != nil {
				this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "数据库操作错误")
				this.ServeJSON()
				return
			}
			order_ids =append(order_ids, id)
		}
	}
	data := map[string]interface{}{
		"order_ids": order_ids,
		"total_pay_amount": total_pay_amount,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "创建订单成功")
	this.ServeJSON()
	return
}


// @Title PayOrder finished
// @Description 支付订单 PayOrder
// @Success 200 status bool, data interface{}, msg string
// @router /pay_order [post]
func (this *OrderController) PayOrder() {

}


// @Title OrderList finished
// @Description 订单列表 OrderList
// @Success 200 status bool, data interface{}, msg string
// @router /order_list [post]
func (this *OrderController) OrderList() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	u_tk, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_lst type_order.OrderListCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_lst); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_lst.OrderListCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	ols, total, err := models.GetGoodsOrderList(order_lst.Page, order_lst.PageSize, u_tk.Id, order_lst.OrderStatus)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		this.ServeJSON()
		return
	}
	var olst_ret []type_order.OrderListRet
	img_path := beego.AppConfig.String("img_root_path")
	for _, value := range ols {
		m, _, _ := models.GetMerchantDetail(value.MerchantId)
		gds, _, _ := models.GetGoodsDetail(value.GoodsId)
		ordr := type_order.OrderListRet {
			MerchantId: m.Id,
			MerchantName: m.MerchantName,
			MerchantPhone: m.Phone,
			OrderId:value.Id,
			GoodsName: value.GoodsName,
			GoodsLogo: img_path + value.Logo,
			GoodsPrice: gds.GoodsPrice,
			OrderStatus: value.OrderStatus,
			BuyNums: value.BuyNums,
			PayAmount: value.PayAmount,
			IsCancle: value.IsCancle,
			IsComment: value.IsComment,
		}
		olst_ret = append(olst_ret, ordr)
	}
	data := map[string]interface{}{
		"total":     total,
		"order_lst": olst_ret,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取订单列表成功")
	this.ServeJSON()
	return
}


// @Title OrderDetail finished
// @Description 订单详情 OrderDetail
// @Success 200 status bool, data interface{}, msg string
// @router /order_detail [post]
func (this *OrderController) OrderDetail() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_dtl type_order.OrderDetailCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_dtl); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_dtl.OrderDetailCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	ord_dtl, code, err := models.GetGoodsOrderDetail(order_dtl.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	img_path := beego.AppConfig.String("img_root_path")
	var addr models.UserAddress
	addr.Id = ord_dtl.AddressId
	addrs, _, _ := addr.GetAddressById()
	mct, _, _ := models.GetMerchantDetail(ord_dtl.MerchantId)
	gdsdtl, _, _ := models.GetGoodsDetail(ord_dtl.GoodsId)
	var ret_ordr *type_order.ReturnOrderProcess
	if ord_dtl.IsCancle != 0 {
		order_process, _, err := models.GetOrderProcessDetail(ord_dtl.Id)
		if err != nil && order_process != nil {
			ret_ordr = &type_order.ReturnOrderProcess{
				ProcessId: order_process.Id,
				ReturnUser: mct.ContactUser,
				ReturnPhone: mct.Phone,
				ReturnAddress: mct.Address,
				// 0:等待卖家确认; 1:卖家已同意; 2:卖家拒绝; 3:等待买家邮寄; 4:等待卖家收货; 5:卖家已经发货; 6:等待买家收货; 7:已完成
				Process: order_process.Process,
				LeftTime: order_process.LeftTime,
			}
		} else {
			ret_ordr = nil
		}
	} else {
		ret_ordr = nil
	}
	odl := type_order.OrderDetailRet{
		OrderId: ord_dtl.Id,
		Logistics: ord_dtl.Logistics,
		ShipNumber: ord_dtl.ShipNumber,
		RecUser: addrs.UserName,
		RecPhone: addrs.Phone,
		RecAddress:addrs.Address,
		MerchantId: mct.Id,
		MerchantName: mct.MerchantName,
		GoodsName: gdsdtl.GoodsName,
		GoodsLogo: img_path + gdsdtl.Logo,
		GoodsPrice: gdsdtl.GoodsPrice,
		OrderStatus: ord_dtl.OrderStatus,
		BuyNums: ord_dtl.BuyNums,
		PayAmount: ord_dtl.PayAmount,
		ShipFee: 0,
		Coupons: ord_dtl.PayCoupon,
		PayWay: ord_dtl.PayWay,
		OrderNumber: ord_dtl.OrderNumber,
		PayTime: ord_dtl.PayAt,
		CreateTime: ord_dtl.CreatedAt,
		IsCancle: ord_dtl.IsCancle,
		IsComment: ord_dtl.IsComment,
		RetrurnOrder: ret_ordr,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, odl, "获取订单详情成功")
	this.ServeJSON()
	return
}


// @Title ReturnGoodsOrder finished
// @Description 订单换退货 ReturnGoodsOrder
// @Success 200 status bool, data interface{}, msg string
// @router /return_goods_order [post]
func (this *OrderController) ReturnGoodsOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_ret type_order.ReturnGoodsOrderCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order_ret); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := order_ret.ReturnGoodsOrderCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	odr, _, _ := models.GetGoodsOrderDetail(order_ret.OrderId)
	if odr.IsCancle != 0 {
		this.Data["json"] = RetResource(false, types.AlreadyCancleOrder, nil, "该订单已经发起退换货")
		this.ServeJSON()
		return
	}
	ord_dtl, code, err := models.ReturnGoodsOrder(order_ret)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"order_id": ord_dtl.Id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "退换货成功")
	this.ServeJSON()
	return
}


// @Title CancleReturnGoodsOrder finished
// @Description 撤销换退货 CancleReturnGoodsOrder
// @Success 200 status bool, data interface{}, msg string
// @router /cancle_return_goods_order [post]
func (this *OrderController) CancleReturnGoodsOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var cancle_order type_order.CancleReturnGoodsOrderCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cancle_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := cancle_order.CancleReturnGoodsOrderCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl, code, err := models.GetGoodsOrderDetail(cancle_order.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl.IsCancle = 0
	err = order_dtl.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"order_id": order_dtl.Id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "撤销退换货成功")
	this.ServeJSON()
	return
}



// @Title ConfirmRecvGoods finished
// @Description 确认收货 ConfirmRecvGoods
// @Success 200 status bool, data interface{}, msg string
// @router /confirm_revc_goods [post]
func (this *OrderController) ConfirmRecvGoods() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var cancle_order type_order.CancleReturnGoodsOrderCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &cancle_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := cancle_order.CancleReturnGoodsOrderCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl, code, err := models.GetGoodsOrderDetail(cancle_order.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	order_dtl.OrderStatus = 5
	err = order_dtl.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err, err.Error())
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"order_id": order_dtl.Id,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "确认收货成功")
	this.ServeJSON()
	return
}