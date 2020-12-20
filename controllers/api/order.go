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
// @Description 创建订单 CreateOrder
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
	_, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
}


// @Title OrderDetail finished
// @Description 订单详情 OrderDetail
// @Success 200 status bool, data interface{}, msg string
// @router /order_detail [post]
func (this *OrderController) OrderDetail() {

}