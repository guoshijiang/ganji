package api

import (
	"encoding/json"
	"ganji/common/utils"
	"ganji/models"
	"ganji/types"
	"ganji/types/pay"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type PayController struct {
	beego.Controller
}


// @Title SingleOrderPay finished
// @Description 单个订单支付 SingleOrderPay
// @Success 200 status bool, data interface{}, msg string
// @router /single_order_pay [post]
func (this *PayController) SingleOrderPay() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	u_t, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	single_order := pay.SingleOrderPayCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &single_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := single_order.SingleOrderPayCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	ordr, code, err := models.GetGoodsOrderDetail(single_order.OrderId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	ordr.PayAmount = single_order.PayAmount
	ordr.PayCoupon = single_order.CouponPay
	ordr.PayWay = single_order.PayWay
	ordr.OrderStatus = 1
	if single_order.PayWay == 0 { // 积分兑换
		i_g, _ := models.GetIntegralByUserId(u_t.Id)
		if i_g.TotalIg < single_order.PayAmount {
			this.Data["json"] = RetResource(false, types.IntegralNotEnogh, nil, "积分余额不足, 请去赚取更多的积分")
			this.ServeJSON()
			return
		}
		i_g_r := models.IntegralRecord{
			UserId: u_t.Id,
			IntegralName: "积分",
			IntegralType: 4,
			IntegralSource: "积分消费",
			IntegralAmount: single_order.PayAmount,
			OrderNumber: ordr.OrderNumber,
			SourceUserId: 0,
		}
		err := i_g_r.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "生成积分消费记录失败")
			this.ServeJSON()
			return
		}
		i_g.TotalIg = i_g.TotalIg - single_order.PayAmount
		err = i_g.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "更新积分余额失败")
			this.ServeJSON()
			return
		}
		data := map[string]interface{}{
			"order_id": ordr.Id,
		}
		ordr.OrderStatus = 2
		err = ordr.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "支付失败")
			this.ServeJSON()
			return
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "支付成功")
		this.ServeJSON()
		return
	}
	if single_order.PayWay == 1 { // 账户余额支付
		u_w, _ := models.GetWalletByUserId(u_t.Id)
		if u_w.TotalAmount < single_order.PayAmount {
			this.Data["json"] = RetResource(false, types.UserWalletNotEnogh, nil, "账户余额不足, 请去充值")
			this.ServeJSON()
			return
		}
		u_w_r := models.WalletRecord{
			UserId: u_t.Id,
			Amount: single_order.PayAmount,
			OrderNumber: ordr.OrderNumber,
			Type: 3,
			Source: 2,
			IsHanle: 2,
			DealUser: "市集APP",
			Status: 1,
		}
		err := u_w_r.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "生成钱包消费记录失败")
			this.ServeJSON()
			return
		}
		u_w.TotalAmount = u_w.TotalAmount - single_order.PayAmount
		err = u_w.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "更新钱包余额失败")
			this.ServeJSON()
			return
		}
		data := map[string]interface{}{
			"order_id": ordr.Id,
		}
		ordr.OrderStatus = 2
		err = ordr.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "支付失败")
			this.ServeJSON()
			return
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "支付成功")
		this.ServeJSON()
		return
	}
	if single_order.PayWay == 2 { // 微信支付
		this.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "暂时不支持该支付方式")
		this.ServeJSON()
		return
	}
	if single_order.PayWay == 3 {  // 支付宝支付
		err = ordr.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "支付失败")
			this.ServeJSON()
			return
		}
		pay_amount := strconv.FormatFloat(ordr.PayAmount,'E',-1,64)
		notify_url := beego.AppConfig.String("pay_notify_url")
		return_url := beego.AppConfig.String("dw_return_url")
		zhifubao_config := utils.AliPayZfb(notify_url, return_url, ordr.OrderNumber, pay_amount)
		this.Data["json"] = RetResource(true, types.ReturnSuccess, zhifubao_config, "支付进入支付中状态")
		this.ServeJSON()
		return
	}
}



// @Title BatchOrderPay finished
// @Description 批量订单支付 BatchOrderPay
// @Success 200 status bool, data interface{}, msg string
// @router /batch_order_pay [post]
func (this *PayController) BatchOrderPay() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	u_t, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	batch_order := pay.BatchOrderPayCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &batch_order); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := batch_order.BatchOrderPayCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	var order_total_pay_amount float64
	order_total_pay_amount = 0
	batch_ordr, code, err := models.GetGoodsOrderBatchList(batch_order.BatchOrderId, u_t.Id)
	for _, v := range batch_ordr {
		order_total_pay_amount = order_total_pay_amount + v.PayAmount
		v.PayWay = batch_order.PayWay
		err = v.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, code, nil, "支付失败")
			this.ServeJSON()
			return
		}
	}
	if order_total_pay_amount != batch_order.TotalPayAmount {
		this.Data["json"] = RetResource(false, types.PayAmountError, nil, "您的支付金额不对，请联系客服处理")
		this.ServeJSON()
		return
	}
	if batch_order.PayWay == 0 { // 积分兑换
		i_g, _ := models.GetIntegralByUserId(u_t.Id)
		if i_g.TotalIg < batch_order.TotalPayAmount {
			this.Data["json"] = RetResource(false, types.IntegralNotEnogh, nil, "积分余额不足, 请去赚取更多的积分")
			this.ServeJSON()
			return
		}
		i_g_r := models.IntegralRecord{
			UserId: u_t.Id,
			IntegralName: "积分",
			IntegralType: 4,
			IntegralSource: "积分消费",
			IntegralAmount: batch_order.TotalPayAmount,
			OrderNumber: batch_order.BatchOrderId,
			SourceUserId: 0,
		}
		err := i_g_r.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "生成积分消费记录失败")
			this.ServeJSON()
			return
		}
		i_g.TotalIg = i_g.TotalIg - batch_order.TotalPayAmount
		err = i_g.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "更新积分余额失败")
			this.ServeJSON()
			return
		}
		for _, v := range batch_ordr {
			v.OrderStatus = 2
			err = v.Update()
			if err != nil {
				this.Data["json"] = RetResource(false, code, nil, "支付失败")
				this.ServeJSON()
				return
			}
		}
		data := map[string]interface{}{
			"batch_order_id": batch_order.BatchOrderId,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "支付成功")
		this.ServeJSON()
		return
	}

	if batch_order.PayWay == 1 { // 账户余额支付
		u_w, _ := models.GetWalletByUserId(u_t.Id)
		if u_w.TotalAmount < batch_order.TotalPayAmount {
			this.Data["json"] = RetResource(false, types.UserWalletNotEnogh, nil, "账户余额不足, 请去充值")
			this.ServeJSON()
			return
		}
		u_w_r := models.WalletRecord{
			UserId: u_t.Id,
			Amount: batch_order.TotalPayAmount,
			OrderNumber: batch_order.BatchOrderId,
			Type: 3,
			Source: 2,
			IsHanle: 2,
			DealUser: "市集APP",
			Status: 1,
		}
		err := u_w_r.Insert()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "生成钱包消费记录失败")
			this.ServeJSON()
			return
		}
		u_w.TotalAmount = u_w.TotalAmount - batch_order.TotalPayAmount
		err = u_w.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "更新钱包余额失败")
			this.ServeJSON()
			return
		}
		data := map[string]interface{}{
			"batch_order_id": batch_order.BatchOrderId,
		}
		for _, v := range batch_ordr {
			v.OrderStatus = 2
			err = v.Update()
			if err != nil {
				this.Data["json"] = RetResource(false, code, nil, "支付失败")
				this.ServeJSON()
				return
			}
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "支付成功")
		this.ServeJSON()
		return
	}
	if batch_order.PayWay == 2 { // 微信支付
		this.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "暂时不支持该支付方式")
		this.ServeJSON()
		return
	}
	if batch_order.PayWay == 3 {  // 支付宝支付
		zhifubao_config, _ := beego.AppConfig.GetSection("zhifubu")
		this.Data["json"] = RetResource(true, types.ReturnSuccess, zhifubao_config, "支付进入支付中状态")
		this.ServeJSON()
		return
	}
}
