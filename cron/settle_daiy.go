package cron

import (
	"ganji/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

const (
	OrderUnPaid       = 0      // 未支付
	OrderPaySuccess   = 2      // 支付成功
	OrderPayFail      = 3      // 支付失败
	OrderShipGds      = 4      // 已发货
	OrderFinished     = 5      // 已完成
	OrderCancleNormal = 0      // 正常订单
	OrderCancling     = 1      // 退货
	OrderChanging     = 2      // 换货
	OrderCancleSucc   = 3      // 退货成功
	OrderChangeSucc   = 4      // 换货成功
)

type SettleData struct {
	MerchantId          int64    `json:"merchant_id"`           // 商户ID
	OrderNum            int64    `json:"order_num"`             // 订单总数量
	OrderAmount         float64  `json:"order_amount"`          // 订单总金额
	ValidOrderNum       int64    `json:"valid_order_num"`       // 有效订单数量
	ValidOrderAmount    float64  `json:"valid_order_amount"`    // 有效订单金额
	UnpayOrderNum       int64    `json:"unpay_order_num"`       // 未支付订单数量
	UnpayOrderAmount    float64  `json:"unpay_order_amount"`    // 未支付订单金额
	PayfailOrderNum     int64    `json:"payfail_order_num"`     // 支付失败订单数量
	PayfailOrderAmount  float64  `json:"payfail_order_amount"`  // 支付失败订单金额
	PaySuccessOdrNum    int64    `json:"pay_success_odr_num"`   // 支付成功订单数量
	PaySuccessOdrAmount float64  `json:"pay_success_odr_amount"`// 支付成功订单金额
	SendGoodsOdrNum     int64    `json:"send_goods_odr_num"`    // 已经发货订单数量
	SendGoodsOdrAmount  float64  `json:"send_goods_odr_amount"` // 已经发货订单金额
	FinishedOdrNum      int64    `json:"finished_odr_num"`      // 完成订单数量
	FinishedOdrAmount   float64  `json:"finished_odr_amount"`   // 完成订单金额
	ReturnOrderNum      int64    `json:"return_order_num"`      // 退款订单数量
	ReturnOrderAmount   float64  `json:"return_order_amount"`   // 退款订单金额
	ChangeOrderNum      int64    `json:"change_order_num"`      // 换货订单数量
	ChangeOrderAmount   float64  `json:"change_order_amount"`	// 换货订单金额
	SettleAmount        float64  `json:"settle_amount"`         // 结算金额
}


func MerchantSettleDaiy(day string) (err error) {
	db := orm.NewOrm()
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
			err = errors.Wrap(err, "rollback db transaction error in IntergralInviteReward")
		} else {
			err = errors.Wrap(db.Commit(), "commit db transaction error in IntergralInviteReward")
		}
	}()
	merchant_list, err := models.GetMerchantSettleList(db)
	if err != nil {
		logs.Error(err)
		return err
	}
	for _, merchant := range merchant_list {
		var order_num, valid_order_num, unpay_order_num, payfail_order_num, pay_success_odr_num,
		send_goods_odr_num, finished_odr_num, return_order_num, change_order_num int64
		var order_amount, valid_order_amount, unpay_order_amount, payfail_order_amount, pay_success_odr_amount,
		send_goods_odr_amount, finished_odr_amount, return_order_amount, change_order_amount, settle_amount float64
		met_order_list, err := models.GetOrdetListByMet(merchant.Id, db)
		if err != nil {
			logs.Error(err)
			return err
		}
		for _, met_order := range met_order_list{
			order_num += 1
			order_amount += met_order.PayAmount
			// 退货的订单
			if met_order.IsCancle ==  OrderCancleSucc || met_order.IsCancle == OrderCancling {
				return_order_num += 1
				return_order_amount += met_order.PayAmount
			}
			// 换货的订单
			if met_order.IsCancle ==  OrderChanging || met_order.IsCancle == OrderChangeSucc {
				change_order_num += 1
				change_order_amount += met_order.PayAmount
			}
			// 未支付的订单
			if met_order.OrderStatus == OrderUnPaid {
				unpay_order_num += 1
				unpay_order_amount += met_order.PayAmount
			}
			//支付失败的订单
			if  met_order.OrderStatus == OrderPayFail {
				payfail_order_num += 1
				payfail_order_amount += met_order.PayAmount
			}
			// 支付成功并没有退换货的订单
			if met_order.OrderStatus == OrderPaySuccess && met_order.IsCancle == OrderCancleNormal {
				pay_success_odr_num += 1
				pay_success_odr_amount += met_order.PayAmount
			}
			// 发货并没有退换货的订单
			if met_order.OrderStatus == OrderShipGds && met_order.IsCancle == OrderCancleNormal {
				send_goods_odr_num += 1
				send_goods_odr_amount += met_order.PayAmount
			}
			// 完成并没有退换货的订单
			if met_order.OrderStatus == OrderFinished && met_order.IsCancle == OrderCancleNormal {
				finished_odr_num += 1
				valid_order_num += 1
				finished_odr_amount += met_order.PayAmount
				valid_order_amount += met_order.PayAmount
			}
			met_order.IsStatic = 1
			_, err := db.Update(&met_order)
			if err != nil {
				return err
			}
		}
		settle_amount = valid_order_amount
		settle_data := SettleData{
			MerchantId: merchant.Id,
			OrderNum: order_num,
			OrderAmount: order_amount,
			ValidOrderNum: valid_order_num,
			ValidOrderAmount: valid_order_amount,
			UnpayOrderNum: unpay_order_num,
			UnpayOrderAmount: unpay_order_amount,
			PayfailOrderNum: payfail_order_num,
			PayfailOrderAmount: payfail_order_amount,
			PaySuccessOdrNum: pay_success_odr_num,
			PaySuccessOdrAmount: pay_success_odr_amount,
			SendGoodsOdrNum: send_goods_odr_num,
			SendGoodsOdrAmount: send_goods_odr_amount,
			FinishedOdrNum: finished_odr_num,
			FinishedOdrAmount: finished_odr_amount,
			ReturnOrderNum: return_order_num,
			ReturnOrderAmount: return_order_amount,
			ChangeOrderNum: change_order_num,
			ChangeOrderAmount: change_order_amount,
			SettleAmount: settle_amount,
		}
		ok, err := settle_data.StatDataInsertDb(db, day)
		if !ok {
			return err
		}
	}
	return nil
}


func (this SettleData)StatDataInsertDb(db orm.Ormer, day string)(bool, error){
	settle_daily := models.MerchantSettleDaily{
		MerchantId: this.MerchantId,
		OrderNum:  this.OrderNum,
		OrderAmount: this.OrderAmount,
		ValidOrderNum: this.ValidOrderNum,
		ValidOrderAmount: this.ValidOrderAmount,
		UnpayOrderNum: this.UnpayOrderNum,
		UnpayOrderAmount:this.UnpayOrderAmount,
		PayfailOrderNum: this.PayfailOrderNum,
		PayfailOrderAmount: this.PayfailOrderAmount,
		PaySuccessOdrNum: this.PaySuccessOdrNum,
		PaySuccessOdrAmount: this.PaySuccessOdrAmount,
		SendGoodsOdrNum: this.SendGoodsOdrNum,
		SendGoodsOdrAmount: this.SendGoodsOdrAmount,
		FinishedOdrNum: this.FinishedOdrNum,
		FinishedOdrAmount: this.FinishedOdrAmount,
		ReturnOrderNum:this.ReturnOrderNum,
		ReturnOrderAmount:this.ReturnOrderAmount,
		ChangeOrderNum: this.ChangeOrderNum,
		ChangeOrderAmount: this.ChangeOrderAmount,
		SettleAmount: this.SettleAmount,
		StaticTime: day,
	}
	_, err := db.Insert(&settle_daily)
	if err != nil {
		return false, err
	}
	return true, nil
}
