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
	ParamLessZero                 = 3005  // 参数小于 0
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
	UserNotLogin                  = 4031  // 用户没有登陆
	GetImagesFileFail             = 4032  // 获取文件失败
	FileFormatError               = 4033  // 文件格式不符合要求
	FileIsBig                     = 4034  // 文件太大
	CreateFilePathError           = 4035  // 创建文件路径失败
	SaveFileFail                  = 4036  // 保存文件失败
	UpdateUserInfoFail            = 4037  // 更新用户信息失败
	AddressIsEmpty                = 4038  // 地址为空
	AddressIdLessEqError          = 4039  // 地址 ID 为空
	UserIdEmptyError              = 4040  // 用户 ID 为空
	CreateAddressFail             = 4041  // 创建地址失败
	UpdateAddressFail             = 4042  // 修改地址失败
	GetConponFail                 = 4043  // 获取优惠券成功
	GetGoodsListFail              = 4044  // 获取商品列表失败
	GetMerchantListFail           = 4045  // 获取商家列表失败
	UserTokenUserIdNotEqual       = 4046  // 用户ID 和 Token 不符合
	InvalidGoodsPirce             = 4047  // 无效的商品价格
	RealNameEmpty                 = 4048  // 真实名字为空
	IdCardEmpty                   = 4049  // 身份证号为空
	IdCardFormatError             = 4050  // 身份证号格式错误
	UserAuthError                 = 4051  // 实名认证失败
	AlreadyBindPassword           = 4052  // 已经绑定支付密码
	GetGoodsCarListFail           = 4053  // 获取购物车失败
	ExchangeAmountError           = 4054  // 兑换金额不对
	AlreadyCancleOrder            = 4055  // 订单已经取消
	GroupOrderExist               = 4056  // 拼团订单已经存在
	AlreadyHelp                   = 4057  // 已经助力
	FileAlreadUpload              = 5058  // 图片已经上传过了
	NoBindAccount                 = 5059  // 没用绑定账号
	UpdateAccountFail             = 5060  // 更新账号失败
)


type PageSizeData struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (this PageSizeData) PageSizeDataParamValidate() (int, error) {
	if this.Page == 0 {
		return PageIsZero, errors.New("page 不能为 0")
	}
	if this.PageSize == 0 {
		return PageSizeIsZero, errors.New("pageSize 不能为 0")
	}
	return ReturnSuccess, nil
}
