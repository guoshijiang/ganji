package models

import (
	"ganji/common"
	"github.com/astaxie/beego/orm"
)

type MerchantWithdraw struct {
	BaseModel
	Id           int64       `json:"id"`
	MerchantId   int64      `orm:"size(64);index" json:"merchant_id"`  // 商户 ID
	Amount       float64     `orm:"default(0);digits(32);decimals(8)" json:"amount"`  // 金额
	OrderNumber  string      `orm:"size(256)" json:"order_number"`                    // 交易订单号
	IsHanle      int8        `orm:"default(0);index" json:"is_hanle"`                 // 0:审核中；1:审核通过；2:已打款; 3:审核拒绝
	DealUser   	 string      `orm:"size(256)" json:"deal_user"`                       // 处理人
}


func (this *MerchantWithdraw) TableName() string {
	return common.TableName("merchant_withdraw")
}

func (this *MerchantWithdraw) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantWithdraw) Update(fields  ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

