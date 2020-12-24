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
		PayWay: 5,
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
		cmt := models.GoodsOrder{
			GoodsId: goods_dtl.Id,
			MerchantId: goods_dtl.MerchantId,
			AddressId: gds_car_dtl.AddresId,
			GoodsTitle: goods_dtl.Title,
			GoodsName: goods_dtl.GoodsName,
			Logo: goods_dtl.Logo,
			UserId: create_order_gdscar.UserId,
			BuyNums: gds_car_dtl.BuyNums,
			PayWay: 5,
			PayAmount: gds_car_dtl.PayAmount,
			SendIntegral: send_integral,
			OrderNumber: order_nmb.String(),
			OrderStatus: 0,
			FailureReason: "未支付",
		}
		err, _ = cmt.Insert()
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
		}
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "创建订单成功")
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
	for _, value := range ols {
		m, _, _ := models.GetMerchantDetail(value.MerchantId)
		gds, _, _ := models.GetGoodsDetail(value.GoodsId)
		ordr := type_order.OrderListRet {
			MerchantId: m.Id,
			MerchantName: m.MerchantName,
			OrderId:value.Id,
			GoodsName: value.GoodsName,
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
	var addr models.UserAddress
	addr.Id = ord_dtl.AddressId
	addrs, _, _ := addr.GetAddressById()
	mct, _, _ := models.GetMerchantDetail(ord_dtl.MerchantId)
	gdsdtl, _, _ := models.GetGoodsDetail(ord_dtl.GoodsId)
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
	ord_dtl, code, err := models.GetGoodsOrderDetail(order_ret.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	ord_dtl.IsCancle = order_ret.IsCancle
	err = ord_dtl.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "退换货失败")
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
