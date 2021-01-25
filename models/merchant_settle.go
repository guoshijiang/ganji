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
	StartSettleTime    *time.Time `orm:"type(datetime);null" json:"start_settle_time"`                   // 结算开始日期
	EndSettleTime  	   *time.Time `orm:"type(datetime);null" json:"end_settle_time"`                     // 结算结束日期
	MerchantId         int64      `orm:"size(64);index" json:"merchant_id"`                              // 商户 ID
	SettleAmount       float64    `orm:"default(0);digits(22);decimals(8)" json:"settle_amount"`         // 结算金额
	HandUser           string     `orm:"size(64);index" json:"hand_user"`                                 // 平台处理账户
	Status             int8       `orm:"default(0);index" json:"is_settled"`                             // 0:商家已确认； 1:平台已确认； 2：已付款
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


