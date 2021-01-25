package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type MerchantSettleDaily struct {
	BaseModel
	Id                 int64      `json:"id"`
	MerchantId         int64      `orm:"size(64);index" json:"merchant_id"`                              // 商户 ID
	OrderNum           int64      `orm:"size(64);index" json:"order_num"`                                // 订单总数量
	OrderAmount        float64    `orm:"default(0);digits(22);decimals(8)" json:"order_amount"`          // 订单总金额
	ValidOrderNum      int64      `orm:"size(64);index" json:"valid_order_num"`                          // 有效订单数量
	ValidOrderAmount   float64    `orm:"default(0);digits(22);decimals(8)" json:"valid_order_amount"`    // 有效订单金额
	UnpayOrderNum      int64      `orm:"size(64);index" json:"unpay_order_num"`                          // 未支付订单数量
	UnpayOrderAmount   float64    `orm:"default(0);digits(22);decimals(8)" json:"unpay_order_amount"`    // 未支付订单金额
	PayfailOrderNum    int64      `orm:"size(64);index" json:"payfail_order_num"`                        // 未支付订单数量
	PayfailOrderAmount float64    `orm:"default(0);digits(22);decimals(8)" json:"payfail_order_amount"`  // 未支付订单金额
	ReturnOrderNum     int64      `orm:"size(64);index" json:"return_order_num"`                         // 退款订单数量
	ReturnOrderAmount  float64    `orm:"default(0);digits(22);decimals(8)" json:"invalid_order_amount"`  // 退款订单金额
	SettleAmount       float64    `orm:"default(0);digits(22);decimals(8)" json:"settle_amount"`         // 可结算金额
	IsSettled          int8       `orm:"default(0);index" json:"is_settled"`                             // 0:未结算； 1:已结算
	StaticTime         *time.Time `orm:"type(datetime);null" json:"settle_time"`                         // 统计的时间
}


func (this *MerchantSettleDaily) TableName() string {
	return common.TableName("merchant_settle_daily")
}

func (this *MerchantSettleDaily) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *MerchantSettleDaily) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantSettleDaily) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantSettleDaily) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantSettleDaily) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

