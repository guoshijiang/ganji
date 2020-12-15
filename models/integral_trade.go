package models


import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type IntegralTrade struct {
	BaseModel
	Id              int64     `orm:"index" json:"id"`
	UserId          int64     `orm:"size(64);index" json:"user_id"`
	OrderNumber     string    `orm:"size(64);index" json:"exchange_id"`
	IntegralSize    float64   `orm:"default(0);digits(32);decimals(8)" json:"ldst_size"`
	CnySize         float64   `orm:"default(0);digits(32);decimals(8)" json:"usdt_size"`
	Fee             float64   `orm:"default(0);digits(32);decimals(8)" json:"fee"`
	Status          int8      `orm:"default(0)"`  // 0:交易中；1: 交易成功 2: 交易失败
}

func (this *IntegralTrade) TableName() string {
	return common.TableName("integral_trade")
}

func (this *IntegralTrade) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *IntegralTrade) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *IntegralTrade) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *IntegralTrade) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *IntegralTrade) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

