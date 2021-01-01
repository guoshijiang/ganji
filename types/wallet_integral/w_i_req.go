package wallet_integral

import (
	"ganji/types"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
	"strconv"
)

type IntegralExchangeCheck struct {
	UserId          int64  `json:"user_id"`
	IntegralAmount  float64 `json:"integral_amount"`
	ExchangeCny     float64 `json:"exchange_cny"`
	IntegralTradeFee     float64 `json:"integral_trade_fee"`
}

func (this IntegralExchangeCheck) IntegralExchangeCheckParamValidate() (int, error) {
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	if this.IntegralAmount <= 0 {
		return types.ParamLessZero, errors.New("兑换积分数量小于等于 0")
	}
	if this.ExchangeCny <= 0 {
		return types.ParamLessZero, errors.New("兑换积分人民币小于等于 0")
	}

	integral_trade_rate := beego.AppConfig.String("integral_trade_rate")
	integral_trade_fee := beego.AppConfig.String("integral_trade_fee")
	integral_trade_rate_f, _ := strconv.ParseFloat(integral_trade_rate,64)
	integral_trade_fee_f, _ := strconv.ParseFloat(integral_trade_fee,64)
	if this.ExchangeCny * integral_trade_rate_f != this.IntegralAmount {
		return types.ExchangeAmountError, errors.New("兑换金额和积分数量不匹配")
	}
	if this.IntegralTradeFee != integral_trade_fee_f {
		return types.ParamLessZero, errors.New("交易手续费应为" + integral_trade_fee)
	}
	return types.ReturnSuccess, nil
}

type IntegralExchangeRecordCheck struct {
	types.PageSizeData
	UserId  int64 `json:"user_id"`
}


func (this IntegralExchangeRecordCheck) IntegralExchangeRecordCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type IntegralExchangeDetailCheck struct {
	IntegralId     int64   `json:"integral_id"`
}

type WalletRecordCheck struct {
	types.PageSizeData
	UserId int64 `json:"user_id"`
}

func (this WalletRecordCheck) WalletRecordCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type WalletDetailCheck struct {
	WalletId int64 `json:"wallet_id"`
}

func (this WalletDetailCheck) WalletDetailCheckParamValidate() (int, error) {
	if this.WalletId <= 0 {
		return types.ParamLessZero, errors.New("钱包ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}


