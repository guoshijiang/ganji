package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_user "ganji/types/user"
	"github.com/astaxie/beego"
	"strings"
)

type UserInfoController struct {
	beego.Controller
}


// @Title GetUserInfo
// @Description 获取用户信息 GetUserInfo
// @Success 200 status bool, data interface{}, msg string
// @router /user_info [post]
func (this *UserInfoController) GetUserInfo() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var useinfo models.UserInfo
	uinf, err := useinfo.GetUserInfoByUserId(user.Id)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, nil, "用户不存在, 请联系客服处理")
		this.ServeJSON()
		return
	}
	var user_integral float64
	user_ig, _ := models.GetIntegralByUserId(user.Id)
	if user_ig != nil {
		user_integral = user_ig.TotalIg
	} else {
		user_integral = 0
	}
	var user_wallet_cny float64
	user_w, _ := models.GetWalletByUserId(user.Id)
	if user_w != nil {
		user_wallet_cny = user_w.TotalAmount
	} else {
		user_wallet_cny = 0
	}
	user_infos := type_user.UserInfoRet {
		UserId: user.Id,
		Token: user.Token,
		UserName: user.UserName,
		IgAmount: user_integral,
		CnyAmount: user_wallet_cny,
		Phone:  user.Phone,
		Eamil: user.Email,
		Sex: uinf.Sex,
		MemberLevel: user.MemberLevel,
		InviteCode: user.MyInviteCode,
		Avator: user.Avator,
		RealName: uinf.RealName,
		WeiChat: uinf.WeiChat,
		QQ: uinf.QQ,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, user_infos, "获取用户信息成功")
	this.ServeJSON()
	return
}


// @Title UpdateUserInfo
// @Description 修改用户信息 UpdateUserInfo
// @Success 200 status bool, data interface{}, msg string
// @router /update_user_info [post]
func (this *UserInfoController) UpdateUserInfo() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var user_info type_user.UpdateUserInfoCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user_info); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	success, code, err := models.UpdateUserInfo(user.Id, user_info)
	if success {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, user_info, "修改用户信息成功")
	} else {
		this.Data["json"] = RetResource(false, code, err, err.Error())
	}
	this.ServeJSON()
	return
}

// @Title GetMyCoupon
// @Description 获取我的优惠券 GetMyCoupon
// @Success 200 status bool, data interface{}, msg string
// @router /my_coupon [post]
func (this *UserInfoController) GetMyCoupon() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	user, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	user_cp, err := models.GetMyCoupon(user.Id)
	if err != nil && user_cp != nil {
		ucp := type_user.UserConponRet {
			ConponId: user_cp.Id,
			ConponName: user_cp.ConponName,
			IsUsed: user_cp.IsUsed,
			TotalAmount: user_cp.TotalAmount,
			StartTime: user_cp.StartTime,
			EndTime: user_cp.EndTime,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, ucp, "获取我的优惠券成功")
	} else {
		this.Data["json"] = RetResource(false, types.GetConponFail, err, err.Error())
	}
	this.ServeJSON()
	return
}
