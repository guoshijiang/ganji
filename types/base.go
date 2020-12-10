package types

import "github.com/pkg/errors"

// 错误码定义
const (
	ReturnSuccess                 = 2000  // 成功返回
	SystemDbErr                   = 3000  // 数据库错误
	InvalidFormatError            = 3001  // 无效的参数格式
	InvalidVerifyWay              = 3002  // 无效的验证方式
	ParamEmptyError               = 3003  // 传入参数为空
	UserToKenCheckError           = 3004  // 用户 Token 校验失败
	PageIsZero                    = 4000  // 页码 0
	PageSizeIsZero                = 4001  // 每页数量 0
	PhoneEmptyError               = 4002  // 手机号为空
	PhoneFormatError              = 4003  // 手机号码格式不正确
	PhoneVerifyCodeEmptyError     = 4004  // 手机号码验证码为空
	PhoneVerifyCodeError          = 4005  // 手机号码验证码不正确
	EmailEmptyError               = 4006  // 邮箱为空
	EmailFormatError              = 4007  // 邮箱格式不正确
	EmailVerifyCodeEmptyError     = 4008  // 邮箱码验证码为空
	EmailVerifyCodeError          = 4009  // 手机号码验证码不正确
	UserAlreadyRegister           = 4010  // 用户已经注册
	UserNotRegister               = 4011  // 用户还没有注册
	NoThisLoginRisgerWay          = 4012  // 没有这种登陆注册验证方式
	UserIsNotExist                = 4013  // 没有这个用户
	UserIsExist                   = 4014  // 用户已经存在
	InviteCodeNotExist            = 4015  // 没有这个邀请码
	RegisterFail                  = 4016  // 注册失败
	CreateUserFail                = 4017  // 创建用户失败
	InsertIntegralFail            = 4018  // 插入积分失败
	CreateUserWalletFail          = 4019  // 创建用户钱包失败
	PasswordError                 = 4020  // 输入的密码错误
	PasswordIsEmpty               = 4021  // 输入密码为空
	NewOldPasswordEqual           = 4022  // 新旧密码相等
	TwicePasswordNotEqual         = 4023  // 新旧密码相等
	AddLoginTimesError            = 4024  // 添加登陆次数错误
	OldNewPhoneIsEqual            = 4025  // 新旧手机号码一样
	PhoneIsAlreadyBind            = 4026  // 手机号码已经绑定
	OldNewEmailEqual              = 4027  // 新旧邮箱一样
	EamilAlreadyBind              = 4028  // 邮箱已经绑定
	BindPhoneError                = 4029  // 没有绑定手机号
	BindEamilError                = 4030  // 没有绑定邮箱


)


type PageSizeData struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (this PageSizeData) PageSizeDataParamValidate() (int, error) {
	if this.Page == 0 {
		return PageIsZero, errors.New("页码数量不能为 0")
	}
	if this.PageSize == 0 {
		return PageSizeIsZero, errors.New("每页显示的数量不能为 0")
	}
	return ReturnSuccess, nil
}
