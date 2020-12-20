package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	"ganji/types/wallet_integral"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	uuid "github.com/satori/go.uuid"
)

type WalletIntegralController struct {
	beego.Controller
}


// @Title MyIntegral
// @Description 我的积分余额 MyIntegral
// @Success 200 status bool, data interface{}, msg string
// @router /my_integral [post]
func (this *WalletIntegralController) MyIntegral() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_token, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	integral, err := models.GetIntegralByUserId(user_token.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "获取我的积分余额失败")
		this.ServeJSON()
		return
	}
	integral_trade_rate := beego.AppConfig.String("integral_trade_rate")
	integral_trade_fee := beego.AppConfig.String("integral_trade_fee")
	integral_fee_rate_f, _ := strconv.ParseFloat(integral_trade_rate,64)
	integral_trade_fee_f, _ := strconv.ParseFloat(integral_trade_fee,64)
	data := map[string]interface{}{
		"total_integral": integral.TotalIg,
		"integral_trade_rate": integral_fee_rate_f,
		"integral_trade_fee": integral_trade_fee_f,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取订单列表成功")
	this.ServeJSON()
	return
}


// @Title IntegralExchange
// @Description 积分兑换 IntegralExchange
// @Success 200 status bool, data interface{}, msg string
// @router /integral_exchange [post]
func (this *WalletIntegralController) IntegralExchange () {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_token, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var ig_excg wallet_integral.IntegralExchangeCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ig_excg); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := ig_excg.IntegralExchangeCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if user_token.Id != ig_excg.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	order_nmb := uuid.NewV4()
	ig_trade := models.IntegralTrade{
		UserId: user_token.Id,
		OrderNumber:order_nmb.String(),
		IntegralSize: ig_excg.IntegralAmount,
		CnySize: ig_excg.ExchangeCny,
		Fee: ig_excg.ExchangeFee,
		Status: 1,
	}
	err = ig_trade.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "数据库操作失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "积分兑换成功")
	this.ServeJSON()
	return
}


// @Title IntegralExchangeRecordList
// @Description 积分兑换记录 IntegralExchangeRecordList
// @Success 200 status bool, data interface{}, msg string
// @router /integral_record_list [post]
func (this *WalletIntegralController) IntegralRecordList () {

}


// @Title IntegralExchangeList
// @Description 积分兑换详情 IntegralExchangeList
// @Success 200 status bool, data interface{}, msg string
// @router /integral_record_detail [post]
func (this *WalletIntegralController) IntegralRecordDetail () {

}


// @Title IntegralExchangeRecordList
// @Description 积分来源记录 IntegralExchangeRecordList
// @Success 200 status bool, data interface{}, msg string
// @router /integral_record_list [post]
func (this *WalletIntegralController) IntegralSourceRecordList () {

}


// @Title IntegralExchangeList
// @Description 积分来源记录详情 IntegralExchangeList
// @Success 200 status bool, data interface{}, msg string
// @router /integral_record_detail [post]
func (this *WalletIntegralController) IntegralSourceRecordDetail () {

}


// @Title MyIntegral
// @Description 我的钱包余额 MyIntegral
// @Success 200 status bool, data interface{}, msg string
// @router /my_integral [post]
func (this *WalletIntegralController) MyWallet() {

}


// @Title WalletRecordList
// @Description 钱包充提记录列表 WalletRecordList
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_reocrd_list  [post]
func (this *WalletIntegralController) WalletRecordList () {

}


// @Title MyWalletRecord
// @Description 钱包充提记录详情 MyWalletRecord
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_reocrd_detail  [post]
func (this *WalletIntegralController) WalletRecordDetail () {

}

