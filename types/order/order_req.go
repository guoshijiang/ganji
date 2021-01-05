package order

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type CreateOrderCheck struct {
	GoodsId       int64   `json:"goods_id"`
	AddressId     int64   `json:"address_id"`
	UserId        int64   `json:"user_id"`
	BuyNums       int64   `json:"buy_nums"`
	PayWay        int8    `json:"pay_way"`         // 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付
	PayAmount     float64 `json:"pay_amount"`      // 支付金额
	PayIntegral   float64 `json:"pay_integral"`    // 支付积分
	SendIntegral  float64 `json:"send_integral"`   // 购买商品赠送积分
	IsIntegral    int8    `json:"is_integral"`     // 0:非积分兑换产品 1:积分兑换产品
	IsDis         int8    `json:"is_dis"`          // 0:不打折，1:打折活动产品
}

func (this CreateOrderCheck) CreateOrderCheckParamValidate() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品ID小于等于 0")
	}
	if this.AddressId <= 0 {
		return types.ParamLessZero, errors.New("您没有选择地址，或者您还没有添加地址，请去选择或者添加")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	if this.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量小于等于 0")
	}
	if this.PayWay < 0 || this.PayWay > 3 {
		return types.InvalidVerifyWay, errors.New("无效的付款方式")
	}
	if this.PayAmount <= 0 {
		return types.ParamLessZero, errors.New("支付金额小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type CreateOrderGoodsCarCheck struct {
	UserId       int64   `json:"user_id"`
	GoodsCarIds  []int64  `json:"goods_car_ids"`
}

func (this CreateOrderGoodsCarCheck) CreateOrderGoodsCarCheckParamValidate() (int, error) {
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	if this.GoodsCarIds == nil {
		return types.ParamLessZero, errors.New("商品购物车IDS数组不能为空")
	}
	return types.ReturnSuccess, nil
}


type OrderListCheck struct {
	types.PageSizeData
	UserId int64 `json:"user_id"`
	OrderStatus int8 `json:"order_status"`  // 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已经收货; 6: 全部
}

func (this OrderListCheck) OrderListCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	if this.OrderStatus < 0 || this.OrderStatus > 6 {
		return types.InvalidFormatError, errors.New("查看的订单状态无效")
	}
	return types.ReturnSuccess, nil
}


type OrderDetailCheck struct {
	OrderId int64  `json:"order_id"`
	IsCancle int8  `json:"is_cancle"` //0:正常； 1.退换货

}

func (this OrderDetailCheck) OrderDetailCheckParamValidate() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID 小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type ReturnGoodsOrderCheck struct {
	OrderId       int64  `json:"order_id"`
	RetGoodsRs    string `json:"ret_goods_rs"`    // 退货原因
	QsDescribe    string `json:"qs_describe"`     // 问题描述
	QsImgOne      string `json:"qs_img_one"`
	QsImgTwo      string `json:"qs_img_two"`
	QsImgThree    string `json:"qs_img_three"`
	IsRecvGoods   int8   `json:"is_recv_goods"`   // 0:未收到货物，1:已经收到货物
	FundRet       int8   `json:"fund_ret"`        // 1.退货,资金返回钱包账号; 2:退货,资金原路返回; 3:换货
}

func (this ReturnGoodsOrderCheck) ReturnGoodsOrderCheckParamValidate() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID 小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type CancleReturnGoodsOrderCheck struct {
	OrderId  int64 `json:"order_id"`
}

func (this CancleReturnGoodsOrderCheck) CancleReturnGoodsOrderCheckParamValidate() (int, error) {
	if this.OrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID 小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type BatchOrderListCheck struct {
	BatchOrderId  string `json:"batch_order_id"`
}
