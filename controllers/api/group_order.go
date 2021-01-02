package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	"ganji/types/group_order"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type GroupOrderController struct {
	beego.Controller
}


// @Title CreateGroupOrder finished
// @Description 创建助力订单 CreateOrder
// @Success 200 status bool, data interface{}, msg string
// @router /create_group_order [post]
func (this *GroupOrderController) CreateGroupOrder() {
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
	var create_g_order group_order.CreateGroupOrderCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &create_g_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := create_g_order.CreateGroupOrderCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != create_g_order.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	exist := models.ExistOrderByUser(create_g_order.GoodsId, create_g_order.UserId)
	if exist {
		this.Data["json"] = RetResource(false, types.GroupOrderExist, nil, "您已经发起助力订单")
		this.ServeJSON()
		return
	} else {
		gds, _, _ := models.GetGoodsDetail(create_g_order.GoodsId)
		u_addr, _, _ := models.GetUserAddressDefault(create_g_order.UserId)
		currentTime := time.Now()
		days3_later_time := currentTime.AddDate(0, 0, 3).Format("2016-07-09 12:00")
		order_nmb := uuid.NewV4()
		cmt := models.GroupOrder{
			GoodsId: gds.Id,
			MerchantId: gds.MerchantId,
			AddressId: u_addr.Id,
			GoodsTitle: gds.Title,
			Logo: gds.Logo,
			UserId: requestUser.Id,
			BuyNums: 1,
			OrderAmount: 0,
			GroupNumber: gds.GroupNumber,
			HelpNumber: 0,
			OrderNumber: order_nmb.String(),
			OrderStatus: 0,
			FailureReason: "助力中",
			DeadLime: days3_later_time,
		}
		err = cmt.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "创建订单失败")
			this.ServeJSON()
			return
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "创建助力订单成功")
			this.ServeJSON()
			return
		}
	}
}


// @Title GroupOrderList finished
// @Description 助力订单列表 CreateOrder
// @Success 200 status bool, data interface{}, msg string
// @router /group_order_list [post]
func (this *GroupOrderController) GroupOrderList () {
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
	var g_odr_lst group_order.GroupOrderListCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &g_odr_lst); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := g_odr_lst.GroupOrderListCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != g_odr_lst.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	order_list, total, err := models.GroupOrderList(g_odr_lst.Page, g_odr_lst.PageSize,  requestUser.InviteMeUserId, g_odr_lst.UserId, g_odr_lst.QueryWay)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	img_path := beego.AppConfig.String("img_root_path")
	var golst []group_order.GroupOrderListRet
	for _, value := range order_list {
		m, _, _ := models.GetMerchantDetail(value.MerchantId)
		gds, _, _ := models.GetGoodsDetail(value.GoodsId)
		ordr := group_order.GroupOrderListRet {
			MerchantId: m.Id,
			MerchantName: m.MerchantName,
			OrderId: value.Id,
			GoodsName: gds.GoodsName,
			GoodsLogo: img_path + value.Logo,
			GoodsPrice: gds.GoodsPrice,
			OrderStatus: value.OrderStatus,
			GroupNumber: value.GroupNumber,
			HelpNumber: value.HelpNumber,
			IsValid: value.IsValid,
			DeadLime: value.DeadLime,
		}
		golst = append(golst, ordr)
	}
	data := map[string]interface{}{
		"total":     total,
		"order_lst": golst,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取订单列表成功")
	this.ServeJSON()
	return
}


// @Title GroupOrderDetail finished
// @Description 助力订单详情 GroupOrderDetail
// @Success 200 status bool, data interface{}, msg string
// @router /group_order_detail [post]
func (this *GroupOrderController) GroupOrderDetail () {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_t, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var order_dtl group_order.OrderDetailCheck
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
	order_detail, code, err := models.GetGroupOrderDetail(order_dtl.GroupOrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	img_path := beego.AppConfig.String("img_root_path")
	var addr models.UserAddress
	addr.Id = order_detail.AddressId
	addrs, _, _ := addr.GetAddressById()
	mct, _, _ := models.GetMerchantDetail(order_detail.MerchantId)
	gdsdtl, _, _ := models.GetGoodsDetail(order_detail.GoodsId)
	is_help := models.ExistOrderByOdrUser(order_dtl.GroupOrderId, user_t.Id)
	odl := group_order.GroupOrderDetailRet{
		OrderId: order_detail.Id,
		Logistics: order_detail.Logistics,
		ShipNumber: order_detail.ShipNumber,
		RecUser: addrs.UserName,
		RecPhone: addrs.Phone,
		RecAddress:addrs.Address,
		MerchantId: mct.Id,
		MerchantName: mct.MerchantName,
		GoodsName: gdsdtl.GoodsName,
		GoodsLogo: img_path + gdsdtl.Logo,
		GoodsPrice: gdsdtl.GoodsPrice,
		OrderStatus: order_detail.OrderStatus,
		GroupNumber: order_detail.GroupNumber,
		HelpNumber: order_detail.HelpNumber,
		ShipFee: 20,
		OrderNumber: order_detail.OrderNumber,
		PayAt: order_detail.PayAt,
		CreateTime: order_detail.CreatedAt,
		IsHelp: is_help,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, odl, "获取助力订单详情成功")
	this.ServeJSON()
	return
}


// @Title HelpOrder finished
// @Description 帮好友助力 HelpOrder
// @Success 200 status bool, data interface{}, msg string
// @router /help_order [post]
func (this *GroupOrderController) HelpOrder () {
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
	var help_odr group_order.HelpOrderDCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &help_odr); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := help_odr.HelpOrderDCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	success_fail, code, err := models.HelpOrder(help_odr.GroupOrderId, help_odr.BuyUserId, help_odr.SelfUserId)
	if err != nil {
		this.Data["json"] = RetResource(success_fail, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "帮助好友助力成功")
	this.ServeJSON()
	return
}