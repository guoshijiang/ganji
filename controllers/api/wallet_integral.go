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
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取我的积分成功")
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
	wlt_rd := models.WalletRecord{
		UserId: user_token.Id,
		Amount: ig_excg.ExchangeCny,
		OrderNumber: order_nmb.String(),
		Type: 2,
		Source: 3,
	}
	err = wlt_rd.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "数据库操作失败")
		this.ServeJSON()
		return
	}
	integral, err := models.GetIntegralByUserId(user_token.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "获取我的积分余额失败")
		this.ServeJSON()
		return
	} else {
		integral.TotalIg = integral.TotalIg - ig_excg.IntegralAmount
		err = integral.Update()
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "数据库操作失败")
			this.ServeJSON()
			return
		}
	}
	wlt, err := models.GetWalletByUserId(user_token.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "数据库操作失败")
		this.ServeJSON()
		return
	}
	wlt.TotalAmount =  wlt.TotalAmount + ig_excg.ExchangeCny
	err = wlt.Update()
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
// @router /integral_exrd_list [post]
func (this *WalletIntegralController) IntegralExchangeRecordList () {
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
	var ig_ex_record wallet_integral.IntegralExchangeRecordCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ig_ex_record); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := ig_ex_record.IntegralExchangeRecordCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if user_token.Id != ig_ex_record.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	var ig_record_r []wallet_integral.IntegralRecordListRet
	ig_trade_list, total, err := models.GetIntegralTradetList(ig_ex_record.Page, ig_ex_record.PageSize, ig_ex_record.UserId)
	for _, value := range ig_trade_list {
		ig_r := wallet_integral.IntegralRecordListRet{
			IntegralId: value.Id,
			IntegralType: "积分兑换",
			IntegralAmount: value.IntegralSize,
			CnyAmount: value.CnySize,
			CreateTime: value.CreatedAt,
		}
		ig_record_r = append(ig_record_r, ig_r)
	}
	data := map[string]interface{}{
		"total": total,
		"ig_record_r": ig_record_r,
	}
	this.Data["json"] = RetResource(false, types.ReturnSuccess, data, "获取积分兑换记录成功")
	this.ServeJSON()
	return
}


// @Title IntegralExchangeList
// @Description 积分兑换详情 IntegralExchangeList
// @Success 200 status bool, data interface{}, msg string
// @router /integral_record_detail [post]
func (this *WalletIntegralController) IntegralRecordDetail () {
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
	var ig_exdtl wallet_integral.IntegralExchangeDetailCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ig_exdtl); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	igt_db, code, err := models.GetIntegralTradetDetail(ig_exdtl.IntegralId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "数据库查询错误")
		this.ServeJSON()
		return
	}
	idt := wallet_integral.IntegralRecordDetailRet {
		IntegralId: igt_db.Id,
		IntegralType: "积分兑换",
		IntegralAmount: igt_db.IntegralSize,
		CnyAmount: igt_db.CnySize,
		Fee: igt_db.Fee,
		OrderNumber: igt_db.OrderNumber,
		Status: igt_db.Status,
		CreateTime: igt_db.CreatedAt,
	}
	this.Data["json"] = RetResource(false, types.ReturnSuccess, idt, "获取积分兑换详情成功")
	this.ServeJSON()
	return
}


// @Title IntegralExchangeRecordList
// @Description 积分来源记录 IntegralExchangeRecordList
// @Success 200 status bool, data interface{}, msg string
// @router /integral_source_record_list [post]
func (this *WalletIntegralController) IntegralSourceRecordList () {
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
	var ig_ex_record wallet_integral.IntegralExchangeRecordCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ig_ex_record); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := ig_ex_record.IntegralExchangeRecordCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if user_token.Id != ig_ex_record.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	var ig_source_r []wallet_integral.IntegralRecordListRet
	ig_source_list, total, err := models.GetIntegralRecordList(ig_ex_record.Page, ig_ex_record.PageSize, ig_ex_record.UserId)
	for _, value := range ig_source_list {
		ig_s := wallet_integral.IntegralRecordListRet{
			IntegralId: value.Id,
			IntegralType: value.IntegralSource,
			IntegralAmount: value.IntegralAmount,
			CreateTime: value.CreatedAt,
		}
		ig_source_r = append(ig_source_r, ig_s)
	}
	data := map[string]interface{}{
		"total": total,
		"ig_source_r": ig_source_r,
	}
	this.Data["json"] = RetResource(false, types.ReturnSuccess, data, "获取积分来源记录成功")
	this.ServeJSON()
	return
}


// @Title IntegralSourceRecordDetail
// @Description 积分来源记录详情 IntegralSourceRecordDetail
// @Success 200 status bool, data interface{}, msg string
// @router /integral_source_detail [post]
func (this *WalletIntegralController) IntegralSourceRecordDetail () {
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
	var ig_exdtl wallet_integral.IntegralExchangeDetailCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ig_exdtl); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	igt, code, err := models.GetIntegralRecordDetail(ig_exdtl.IntegralId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "数据库查询错误")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(false, types.ReturnSuccess, igt, "获取积分来源详情成功")
	this.ServeJSON()
	return
}


// @Title MyWallet
// @Description 我的钱包余额 MyWallet
// @Success 200 status bool, data interface{}, msg string
// @router /my_wallet [post]
func (this *WalletIntegralController) MyWallet() {
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
	wallet, err := models.GetWalletByUserId(user_token.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, "获取我的钱包余额失败")
		this.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"asset_name": wallet.AssetName,
		"total_amount": wallet.TotalAmount,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取钱包余额成功")
	this.ServeJSON()
	return
}


// @Title WalletRecordList
// @Description 钱包充提记录列表 WalletRecordList
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_reocrd_list  [post]
func (this *WalletIntegralController) WalletRecordList () {
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
	var wallet_record wallet_integral.WalletRecordCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &wallet_record); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := wallet_record.WalletRecordCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	if user_token.Id != wallet_record.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	var wallet_record_list []wallet_integral.WalletRecordListRet
	wr_list, total, err := models.GetWalletRecordList(wallet_record.Page, wallet_record.PageSize, wallet_record.UserId)
	for _, value := range wr_list {
		// 0:充值；1:提现 2:积分兑换 3:消费
		var type_str  string
		if value.Type == 0 {
			type_str = "充值"
		} else if value.Type == 1 {
			type_str = "提现"
		} else if value.Type == 2 {
			type_str = "积分兑换"
		} else {
			type_str = "消费"
		}
		wl_r := wallet_integral.WalletRecordListRet{
			RecordId: value.Id,
			IntegralType: type_str,
			TotalAmount: value.Amount,
			CreateTime: value.CreatedAt,
		}
		wallet_record_list = append(wallet_record_list, wl_r)
	}
	data := map[string]interface{}{
		"total": total,
		"ig_source_r": wallet_record_list,
	}
	this.Data["json"] = RetResource(false, types.ReturnSuccess, data, "获取钱包记录列表成功")
	this.ServeJSON()
	return
}


// @Title MyWalletRecord
// @Description 钱包充提记录详情 MyWalletRecord
// @Success 200 status bool, data interface{}, msg string
// @router /wallet_reocrd_detail  [post]
func (this *WalletIntegralController) WalletRecordDetail () {
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
	var w_dtl wallet_integral.WalletDetailCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &w_dtl); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	wt_dtl, code, err := models.GetWalletRecordDetail(w_dtl.WalletId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, err.Error(), "数据库查询错误")
		this.ServeJSON()
		return
	}
	// 0:充值；1:提现 2:积分兑换 3:消费
	var type_str, source_str string
	if wt_dtl.Type == 0 {
		type_str = "充值"
	} else if wt_dtl.Type == 1 {
		type_str = "提现"
	} else if wt_dtl.Type == 2 {
		type_str = "积分兑换"
	} else {
		type_str = "消费"
	}
	// 0：支付宝 1:微信; 2:积分兑换
	if wt_dtl.Source == 0 {
		source_str = "支付宝"
	} else if wt_dtl.Source == 1 {
		source_str = "微信"
	} else {
		source_str = "积分兑换"
	}
	wlt_dtl := wallet_integral.WalletRecordDetailRet {
		RecordId: wt_dtl.Id,
		IntegralType: type_str,
 		IntegralSource: source_str,
		TotalAmount:wt_dtl.Amount,
		OrderNumber: wt_dtl.OrderNumber,
		CreateTime: wt_dtl.CreatedAt,
	}
	this.Data["json"] = RetResource(false, types.ReturnSuccess, wlt_dtl, "获取钱包记录详情成功")
	this.ServeJSON()
	return
}

