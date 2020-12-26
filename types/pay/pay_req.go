package pay

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type SingleOrderPayCheck struct {
	OrderId   int64   `json:"order_id"`
	PayAmount float64 `json:"pay_amount"`    // 付款金额或者付款积分
	CouponPay float64 `json:"coupon_pay"`    // 优惠券抵扣
	PayWay    int8    `json:"pay_way"`       // 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付
}

func (this SingleOrderPayCheck) SingleOrderPayCheckParamValidate() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单的 ID 不能小于等于 0")
	}
	if this.PayAmount <= 0 {
		return types.ParamLessZero, errors.New("付款金额不能小于等于 0")
	}
	if this.PayWay <0  || this.PayWay > 3 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	return types.ReturnSuccess, nil
}

type BatchOrderPayCheck struct {
	BatchOrderId   string   `json:"batch_order_id"`
	TotalPayAmount float64  `json:"total_pay_amount"` // 总的付款金额或者付款积分
	PayWay         int8     `json:"pay_way"`          // 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付
}

func (this BatchOrderPayCheck) BatchOrderPayCheckParamValidate() (int, error) {
	if this.BatchOrderId == "" {
		return types.RealNameEmpty, errors.New("订单batch id为空")
	}
	if this.TotalPayAmount <= 0 {
		return types.ParamLessZero, errors.New("付款金额不能小于等于 0")
	}
	if this.PayWay <0  || this.PayWay > 3 {
		return types.InvalidVerifyWay, errors.New("无效的支付方式")
	}
	return types.ReturnSuccess, nil
}
