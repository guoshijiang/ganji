package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type MerchantSettle struct {
	BaseModel
	Id                 int64      `json:"id"`
	MerchantId         int64      `orm:"size(64);index" json:"merchant_id"`  // 商户 ID
	OrderNum           int64      `orm:"size(64);index" json:"order_num"`    // 订单总数量
	OrderAmount        float64    `orm:"default(0);digits(22);decimals(8)" json:"order_amount"`  // 订单总金额
	ValidOrderNum      int64      `orm:"size(64);index" json:"valid_order_num"`    // 有效订单数量
	ValidOrderAmount   float64    `orm:"default(0);digits(22);decimals(8)" json:"valid_order_amount"`  // 有效订单金额
	InvalidOrderNum    int64      `orm:"size(64);index" json:"invalid_order_num"`    // 退款订单数量
	InvalidOrderAmount float64    `orm:"default(0);digits(22);decimals(8)" json:"invalid_order_amount"` // 退款订单金额
	SettleTime    *time.Time `orm:"type(datetime);null" json:"settle_time"`  // 统计的时间
}


func (this *MerchantSettle) TableName() string {
	return common.TableName("merchant_settle")
}

func (this *MerchantSettle) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *MerchantSettle) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantSettle) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantSettle) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantSettle) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

