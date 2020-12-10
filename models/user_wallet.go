package models

import (
	"ganji/common"
	"github.com/astaxie/beego/orm"
)

type UserWallet struct {
	BaseModel
	Id          int64     `json:"id"`
	UserId      int64     `orm:"index" json:"user_id"`
	AssetName   string    `orm:"size(128);index" json:"asset_name"` // 资产名称
	TotalAmount float64   `orm:"default(150);digits(22);decimals(8)" json:"total_amount"`
}

func (this *UserWallet) TableName() string {
	return common.TableName("user_wallet")
}

func (this *UserWallet) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *UserWallet) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *UserWallet) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

