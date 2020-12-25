package user_account

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type UserAccountAddCheck struct {
	UserId      int64  `json:"user_id"`
	AcountType  int8   `json:"acount_type"`  // 0 支付宝； 1:微信; 2:银行卡
	AccountName string `json:"account_name"` // 账号名称; 银行名称
	UserName    string `json:"user_name"`    // 用户名字; 银行开户名字
	CardNum     string `json:"card_num"`     // 账号；银行卡号
	Address     string `json:"address"`      // 开户行地址
}

func (this UserAccountAddCheck) UserAddressAddCheckParamValidate() (int, error) {
	if this.UserId <= 0 {
		return types.UserIsNotExist, errors.New("用户不存在, 请联系客服处理")
	}
	if this.AcountType < 0 || this.AcountType > 2 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	if this.AccountName == "" || this.UserName == "" || this.CardNum == ""{
		return types.ParamEmptyError, errors.New("您填写的账号名称，用户名字或者账号为空")
	}
	return types.ReturnSuccess, nil
}

type UserAccountUpdCheck struct {
	UserId        int64  `json:"user_id"`
	UserAccountId int64  `json:"user_account_id"`
	AcountType    int8   `json:"acount_type"`  // 0 支付宝； 1:微信; 2:银行卡
	AccountName   string `json:"account_name"` // 账号名称; 银行名称
	UserName      string `json:"user_name"`    // 用户名字; 银行开户名字
	CardNum       string `json:"card_num"`     // 账号；银行卡号
	Address       string `json:"address"`      // 开户行地址
	IsInvalid     int8   `json:"is_invalid"`   // 0 激活； 1:禁用
}

func (this UserAccountUpdCheck) UserAccountUpdCheckParamValidate() (int, error) {
	if this.UserId <= 0 || this.UserAccountId <= 0{
		return types.UserIsNotExist, errors.New("用户不存在或者用户账号不存在, 请联系客服处理")
	}
	if this.AcountType < 0 || this.AcountType > 2 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	if this.AccountName == "" || this.UserName == "" || this.CardNum == ""{
		return types.ParamEmptyError, errors.New("您填写的账号名称，用户名字或者账号为空")
	}
	if this.IsInvalid < 0 || this.IsInvalid > 1{
		return types.InvalidVerifyWay, errors.New("无效的修改状态")
	}
	return types.ReturnSuccess, nil
}


type UserAccountDelCheck struct {
	UserAccountId int64  `json:"user_account_id"`
}

func (this UserAccountDelCheck) UserAdddressDelParamValidate() (int, error) {
	if this.UserAccountId <= 0 {
		return types.ParamLessZero, errors.New("用户账号不存在, 请联系客服处理")
	}
	return types.ReturnSuccess, nil
}
