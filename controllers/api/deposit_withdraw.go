package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	"ganji/types/w_or_d"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
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
	w_r := models.WalletRecord{
		UserId: requestUser.Id,
		Amount: deposit.Amount,
		OrderNumber:order_nmb.String(),
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
		zhifubao_config, _ := beego.AppConfig.GetSection("zhifubu")
		this.Data["json"] = RetResource(true, types.ReturnSuccess, zhifubao_config, "充值成功")
		this.ServeJSON()
		return
	} else if deposit.PayWay == 1 { // 微信
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "暂时不支持该充值方式, 不久的将来将会上线")
		this.ServeJSON()
		return
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

