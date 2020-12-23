package models


import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type IntegralTrade struct {
	BaseModel
	Id              int64     `orm:"index" json:"id"`
	UserId          int64     `orm:"size(64);index" json:"user_id"`
	OrderNumber     string    `orm:"size(64);index" json:"order_number"`
	IntegralSize    float64   `orm:"default(0);digits(32);decimals(8)" json:"integral_size"`
	CnySize         float64   `orm:"default(0);digits(32);decimals(8)" json:"cny_size"`
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


func GetIntegralTradetList(page, pageSize int, user_id int64) ([]*IntegralTrade, int64, error) {
	offset := (page - 1) * pageSize
	ig_trade_list := make([]*IntegralTrade, 0)
	query := orm.NewOrm().QueryTable(IntegralTrade{}).Filter("UserId", user_id)
	total, _ := query.Count()
	_, err := query.OrderBy("-CreatedAt").Limit(pageSize, offset).All(&ig_trade_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return ig_trade_list, total, nil
}


func GetIntegralTradetDetail(id int64) (*IntegralTrade, int, error) {
	var integral IntegralTrade
	if err := orm.NewOrm().QueryTable(IntegralTrade{}).Filter("Id", id).RelatedSel().One(&integral); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &integral, types.ReturnSuccess, nil
}


