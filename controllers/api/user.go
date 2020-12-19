package api

import (
	"encoding/json"
	"fmt"
	"ganji/models"
	rds_conn "ganji/redis"
	"ganji/types"
	type_user "ganji/types/user"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type UserController struct {
	beego.Controller
}

// @Title SendPhoneCode
// @Description 发送手机号验证码 SendPhoneCode
// @Success 200 status bool, data interface{}, msg string
// @router /send_phone_code [post]
func (this *UserController) SendPhoneCode() {
	phone_number := type_user.PhoneNumberCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &phone_number); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err.Error(), "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := phone_number.PhoneNumberParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		// verify_code, _ := strconv.Atoi(common.GenValidateCode(6))
		verify_code := 666666
		rds_conn.RdsConn.Del(phone_number.Phone)
		rds_conn.RdsConn.Set(phone_number.Phone, fmt.Sprintf("%d", verify_code), time.Duration(1000)*time.Second).Err()
		// utils.SendMesseageCode(phone_number.Phone, verify_code)
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发送手机号验证码成功")
		this.ServeJSON()
		return
	}
}


// @Title PhoneCodeCheck
// @Description 手机号验证码校验 PhoneCodeCheck
// @Success 200 status bool, data interface{}, msg string
// @router /phone_code_check [post]
func (this *UserController) PhoneCodeCheck() {
	phone_code_check := type_user.PhoneCodeCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &phone_code_check); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := phone_code_check.ReqPhoneCodeCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "手机号验证码校验成功")
	this.ServeJSON()
	return
}


// @Title PhoneNumberRigsterCheck
// @Description 手机号是否注册校验 PhoneNumberRigsterCheck
// @Success 200 status bool, data interface{}, msg string
// @router /phone_number_check [post]
func (this *UserController) PhoneNumberRigsterCheck() {
	phone_reg_check := type_user.PhoneRigsterCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &phone_reg_check); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := phone_reg_check.PhoneRigsterCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		var user_m models.User
		success := user_m.ExistByPhone(phone_reg_check.Phone)
		if phone_reg_check.LoginRisger == 1 { // 登陆
			if success == true {
				this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "您已经注册，请继续登陆")
				this.ServeJSON()
				return
			} else {
				this.Data["json"] = RetResource(false, types.UserNotRegister, nil, "您还没有注册，请去注册")
				this.ServeJSON()
				return
			}
		} else { // 注册
			if success == true {
				this.Data["json"] = RetResource(false, types.UserAlreadyRegister, nil, "您还没有注册，请去登陆")
				this.ServeJSON()
				return
			} else {
				this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "您还没有注册，请继续注册")
				this.ServeJSON()
				return
			}
		}
	}
}

// @Title PostSendEmailCode
// @Description 发送邮箱验证码 PostSendEmailCode
// @Success 200 status bool, data interface{}, msg string
// @router /send_eamil_code [post]
func (this *UserController) PostSendEmailCode() {
	email_code := type_user.EmailNumberCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &email_code); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := email_code.EmailNumberCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		// verify_code, _ := strconv.Atoi(common.GenValidateCode(6))
		verify_code := 666666
		rds_conn.RdsConn.Del(email_code.Email)
		rds_conn.RdsConn.Set(email_code.Email, fmt.Sprintf("%d", verify_code), time.Duration(1000)*time.Second).Err()
		// utils.SendSSLEmaill(email_code.Email, verify_code)
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "发送邮箱验证码成功")
		this.ServeJSON()
		return
	}
}


// @Title EmailCodeCheck
// @Description 邮箱验证码校验 EmailCodeCheck
// @Success 200 status bool, data interface{}, msg string
// @router /eamil_code_check [post]
func (this *UserController) EmailCodeCheck() {
	email_code_check := type_user.EmailCodeCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &email_code_check); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := email_code_check.EmailCodeCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "邮箱验证码校验成功")
			this.ServeJSON()
			return
		}
	}
}


// @Title PostEmailCheck
// @Description 邮箱是否注册校验 PostEmailCheck
// @Success 200 status bool, data interface{}, msg string
// @router /email_check [post]
func (this *UserController) PostEmailCheck() {
	email_reg_check := type_user.EamilRigsterCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &email_reg_check); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := email_reg_check.EmailNumberCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		var user_m models.User
		success := user_m.ExistByEmail(email_reg_check.Email)
		if email_reg_check.LoginRisger == 1 {  // 登陆
			if success == true {
				this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "您已经注册，请继续登陆")
				this.ServeJSON()
				return

			} else {
				this.Data["json"] = RetResource(false, types.UserNotRegister, nil, "您还没有注册，请去注册")
				this.ServeJSON()
				return
			}
		} else {  // 注册
			if success == true {
				this.Data["json"] = RetResource(false, types.UserAlreadyRegister, nil, "您还没有注册，请去登陆")
				this.ServeJSON()
				return
			} else {
				this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "您还没有注册，请继续注册")
				this.ServeJSON()
				return
			}
		}
	}
}


// @Title UserRegister
// @Description 用户注册 UserRegister
// @Success 200 status bool, data interface{}, msg string
// @router /register [post]
func (this *UserController) UserRegister() {
	register_parm := type_user.UserRegisterCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &register_parm); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := register_parm.UserRegisterCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		success, code, err := models.RegisterByPhoneOrEmail(register_parm)
		if success {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "注册成功")
		} else {
			this.Data["json"] = RetResource(success, code, nil, err.Error())
		}
		this.ServeJSON()
		return
	}
}


// @Title UserLogin
// @Description 用户登陆 UserLogin
// @Success 200 status bool, data interface{}, msg string
// @router /login [post]
func (u *UserController) UserLogin() {
	login_parm := type_user.UserLoginCheck{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &login_parm); err != nil {
		u.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		u.ServeJSON()
		return
	} else {
		if code, err := login_parm.UserLoginCheckParamValidate(); err != nil {
			u.Data["json"] = RetResource(false, code, nil, err.Error())
			u.ServeJSON()
			return
		}
		_, user_data, code, err := models.LoginByPhoneOrEmail(login_parm)
		if code == types.ReturnSuccess {
			u.Data["json"] = RetResource(true, types.ReturnSuccess, user_data, "登陆成功")
		} else {
			u.Data["json"] = RetResource(false, code, nil, err.Error())
		}
		u.ServeJSON()
		return
	}
}


// @Title BindFundPassword
// @Description 绑定支付密码 BindFundPassword
// @Success 200 status bool, data interface{}, msg string
// @router /bind_fund_password [post]
func (this *UserController) BindFundPassword() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	usr_t, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	bind_pwd := type_user.BindFundPasswordCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &bind_pwd); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err.Error(), "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := bind_pwd.BindFundPasswordCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		success, code, err := models.BindFundPassword(bind_pwd, usr_t.Id)
		if code != types.ReturnSuccess {
			this.Data["json"] = RetResource(success, code, nil, err.Error())
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "绑定支付密码成功")
		}
		this.ServeJSON()
		return
	}
}


// @Title UpdatePassword
// @Description 修改登陆或者支付密码 UpdatePassword
// @Success 200 status bool, data interface{}, msg string
// @router /update_password [post]
func (this *UserController) UpdatePassword() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	usr_t, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	upd_pwd := type_user.UpdatePasswordCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_pwd); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := upd_pwd.UpdatePasswordCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		success, code, err := models.UpdatePassword(upd_pwd, usr_t.Id)
		if code != types.ReturnSuccess {
			this.Data["json"] = RetResource(success, code, nil, err.Error())
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "修改密码成功")
		}
		this.ServeJSON()
		return
	}
}


// @Title UpdateCreatePhoneEmail
// @Description 修改绑定手机号码邮箱 UpdateCreatePhoneEmail
// @Success 200 status bool, data interface{}, msg string
// @router /update_create_phone_email [post]
func (this *UserController) UpdateCreatePhoneEmail() {
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
	upd_phone_email := type_user.UpdateCreatePhoneEmailCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_phone_email); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := upd_phone_email.UpdateCreatePhoneEmailCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		success, code, err := models.UpdateOrCrearePhoneEmail(upd_phone_email, user_token.Id)
		if success {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, upd_phone_email, "绑定邮箱成功")
		} else {
			this.Data["json"] = RetResource(false, code, err, err.Error())
		}
		this.ServeJSON()
		return
	}
}


// @Title ForgetPassword
// @Description 找回密码 ForgetPassword
// @Success 200 status bool, data interface{}, msg string
// @router /forget_password [post]
func (this *UserController) ForgetPassword() {
	forget_password := type_user.ForgetPasswordCheck{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &forget_password); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	} else {
		if code, err := forget_password.ForgetPasswordCheckParamValidate(); err != nil {
			this.Data["json"] = RetResource(false, code, nil, err.Error())
			this.ServeJSON()
			return
		}
		success, code, err := models.ForgetPassword(forget_password)
		if err != nil {
			this.Data["json"] = RetResource(success, code, nil, err.Error())
		} else {
			this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "找回密码成功")
		}
		this.ServeJSON()
		return
	}
}


// @Title GetInviteCode
// @Description 获取我的邀请码 GetInviteCode
// @Success 200 status bool, data interface{}, msg string
// @router /get_invite_code [post]
func (u *UserController) GetInviteCode() {
	bearerToken := u.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		u.Data["json"] = RetResource(false, types.UserNotLogin, nil, "您还没有登陆，请登陆")
		u.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user_token, err := models.GetUserByToken(token)
	if err != nil {
		u.Data["json"] = RetResource(false, types.UserNotLogin, nil, "您还没有登陆，请登陆")
		u.ServeJSON()
		return
	}
	invite_code := make(map[string]string)
	invite_code["invite_code"] = user_token.MyInviteCode
	u.Data["json"] = RetResource(true, types.ReturnSuccess, invite_code, "获取我的邀请码成功")
	u.ServeJSON()
	return
}

