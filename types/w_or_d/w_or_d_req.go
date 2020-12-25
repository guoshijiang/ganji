package w_or_d

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type DepositCheck struct {
	UserId  int64   `json:"user_id"`
	Amount  float64 `json:"amount"`
	PayWay  int8    `json:"pay_way"`  // 0 支付宝； 1:微信;  2:银行卡
}

func (this DepositCheck) DepositCheckParamValidate() (int, error) {
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("无效用户")
	}
	if this.Amount <= 0 {
		return types.ParamLessZero, errors.New("提现金额小于等于 0; 拒绝提现")
	}
	if this.PayWay < 0 || this.PayWay > 2 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	return types.ReturnSuccess, nil
}


type WithdrawCheck struct {
	UserId  int64   `json:"user_id"`
	Amount  float64 `json:"amount"`
	PayWay  int8    `json:"pay_way"`  // 0 支付宝； 1:微信;  2:银行卡
}

func (this WithdrawCheck) WithdrawCheckParamValidate() (int, error) {
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("无效用户")
	}
	if this.Amount <= 0 {
		return types.ParamLessZero, errors.New("提现金额小于等于 0; 拒绝提现")
	}
	if this.PayWay < 0 || this.PayWay > 2 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	return types.ReturnSuccess, nil
}

