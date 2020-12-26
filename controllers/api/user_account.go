package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_account "ganji/types/user_account"
	"github.com/astaxie/beego"
	"strings"
)

type UserAccountController struct {
	beego.Controller
}


// @Title AddAccount finished
// @Description 添加新账号 AddAccount
// @Success 200 status bool, data interface{}, msg string
// @router /add_account [post]
func (this *UserAccountController) AddAccount() {
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
	var add_acct type_account.UserAccountAddCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &add_acct); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := add_acct.UserAddressAddCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != add_acct.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加用户账户")
		this.ServeJSON()
		return
	}
	if add_acct.AcountType == 2 {
		this.Data["json"] = RetResource(false, types.InvalidVerifyWay, nil, "暂时不支持银行卡绑定")
		this.ServeJSON()
		return
	}
	exist := models.AccountExist(requestUser.Id,  add_acct.AcountType)
	if exist == false {
		account := models.UserAccount{
			UserId: add_acct.UserId,
			AcountType: add_acct.AcountType,
			AccountName: add_acct.AccountName,
			UserName: add_acct.UserName,
			CardNum: add_acct.CardNum,
			Address: add_acct.Address,
			IsInvalid: 0,
		}
		if err := account.Insert(); err != nil {
			this.Data["json"] = RetResource(false, types.CreateAddressFail, nil, "创建用户账户失败，请联系客服处理")
			this.ServeJSON()
			return
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "添加用户账号成功")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(false, types.AlreadyBindAccount, nil, "已经绑定改类型的账号，请勿重复绑定")
		this.ServeJSON()
		return
	}
}


// @Title UpdAccount finished
// @Description 修改地址和手机好码 UpdAccount
// @Success 200 status bool, data interface{}, msg string
// @router /upd_account [post]
func (this *UserAccountController) UpdAccount() {
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
	var upd_acct type_account.UserAccountUpdCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_acct); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := upd_acct.UserAccountUpdCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	usr_acct, code, err := models.GetAccountById(upd_acct.UserAccountId)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	if upd_acct.AccountName != "" {
		usr_acct.AccountName = upd_acct.AccountName
	}
	if upd_acct.UserName != "" {
		usr_acct.UserName = upd_acct.UserName
	}
	if upd_acct.CardNum != "" {
		usr_acct.CardNum = upd_acct.CardNum
	}
	if upd_acct.Address != "" {
		usr_acct.Address = upd_acct.Address
	}
	if upd_acct.IsInvalid == 0 || upd_acct.IsInvalid ==1 {
		usr_acct.IsInvalid = upd_acct.IsInvalid
	}
	err = usr_acct.Update()
	if err != nil {
		this.Data["json"] = RetResource(false, types.UpdateAccountFail, nil, err.Error())
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "更新账号成功")
		this.ServeJSON()
		return
	}
}

// @Title DelAccount finished
// @Description 删除账号 DelAccount
// @Success 200 status bool, data interface{}, msg string
// @router /del_account [post]
func (this *UserAccountController) DelAccount () {
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
	var del_acct type_account.UserAccountDelCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &del_acct); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := del_acct.UserAdddressDelParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	var usr_acct models.UserAccount
	usr_acct.Id = del_acct.UserAccountId
	err = usr_acct.Delete()
	if err != nil {
		this.Data["json"] = RetResource(false, types.AddressIsEmpty, err, "删除账户失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "删除账户成功")
	this.ServeJSON()
	return
}


// @Title AccountList a finished
// @Description 获取地址列表 AccountList
// @Success 200 status bool, data interface{}, msg string
// @router /account_list [post]
func (this *UserAccountController) AccountList () {
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
	account_list, code, err := models.GetUserAccountList(user_token.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, account_list, "获取地址用户账户列表成功")
	this.ServeJSON()
	return
}
