package api

import (
	"encoding/base64"
	"encoding/json"
	"ganji/common/utils"
	"ganji/models"
	"ganji/types"
	"ganji/types/w_or_d"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

type DepositWithdrawController struct {
	beego.Controller
}


// @Title Deposit finished
// @Description 充值 Deposit
// @Success 200 status bool, data interface{}, msg string
// @router /deposit [post]
func (this *DepositWithdrawController) Deposit() {
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
	var deposit w_or_d.DepositCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &deposit); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := deposit.DepositCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != deposit.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	order_nmb := uuid.NewV4()
	hex_uuid := base64.RawURLEncoding.EncodeToString(order_nmb.Bytes())
	deposit_order_number := "deposit-" + hex_uuid
	w_r := models.WalletRecord{
		UserId: requestUser.Id,
		Amount: deposit.Amount,
		OrderNumber: deposit_order_number,
		Type: 0,
		Source: deposit.PayWay,
		IsHanle: 0,
		DealUser: "市集",
		Status:0,
	}
	err = w_r.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "充值失败")
		this.ServeJSON()
		return
	}
	if deposit.PayWay == 0 {  // 支付宝
		pay_amount := strconv.FormatFloat(deposit.Amount,'f',-1,64)
		notify_url := beego.AppConfig.String("ali_pay_notify_url")
		return_url := beego.AppConfig.String("ali_dw_return_url")
		zhifubao_config := utils.AliPayZfb(notify_url, return_url, deposit_order_number, pay_amount)
		this.Data["json"] = RetResource(true, types.ReturnSuccess, zhifubao_config, "支付宝充值中")
		this.ServeJSON()
		return
	} else if deposit.PayWay == 1 { // 微信
		ret_data, err := utils.WxPayOrder(deposit_order_number, deposit.Amount)
		if err != nil {
			this.Data["json"] = RetResource(false, types.DepositException, err.Error(), "微信充值异常，请联系客服处理")
			this.ServeJSON()
			return
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, ret_data, "微信充值中")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(true, types.InvalidVerifyWay, nil, "暂时不支持该充值方式")
		this.ServeJSON()
		return
	}
}


// @Title Withdraw finished
// @Description 提现 Withdraw
// @Success 200 status bool, data interface{}, msg string
// @router /withdraw [post]
func (this *DepositWithdrawController) Withdraw() {
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
	var with_draw w_or_d.WithdrawCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &with_draw); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := with_draw.WithdrawCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != with_draw.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	s_or_f := models.AccountExist(requestUser.Id, with_draw.PayWay)
	if s_or_f == false  {
		this.Data["json"] = RetResource(false, types.NoBindAccount, nil, "您没有绑定账号, 请去绑定账号")
		this.ServeJSON()
		return
	}
	order_nmb := uuid.NewV4()
	w_r := models.WalletRecord{
		UserId: requestUser.Id,
		Amount: with_draw.Amount,
		OrderNumber:order_nmb.String(),
		Type: 1,
		Source: with_draw.PayWay,
		IsHanle: 0,
		DealUser: "市集",
		Status:0,
	}
	err = w_r.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "提现失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "提现已经提交，请等待审核")
	this.ServeJSON()
	return
}

