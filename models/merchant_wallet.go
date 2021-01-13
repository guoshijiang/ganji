package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MerchantWallet struct {
	BaseModel
	Id                 int64      `json:"id"`
	MerchantId         int64      `orm:"size(64);index" json:"merchant_id"`  // 商户 ID
	TotalAmount        float64    `orm:"default(0);digits(22);decimals(8)" json:"order_amount"`  // 钱包总金额
	AvailableAmount    float64    `orm:"default(0);digits(22);decimals(8)" json:"valid_order_amount"`  // 钱包可用金额
	WithdrawAmount     float64    `orm:"default(0);digits(22);decimals(8)" json:"invalid_order_amount"` // 钱包提现金额
}


func (this *MerchantWallet) TableName() string {
	return common.TableName("merchant_wallet")
}

func (this *MerchantWallet) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *MerchantWallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantWallet) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantWallet) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantWallet) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

