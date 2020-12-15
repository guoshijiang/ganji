package models


import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type IntegralRecord struct {
	BaseModel
	Id             int64      `json:"id"`
	UserId         int64      `orm:"index" json:"user_id"`
	IntegralName   string     `orm:"size(128);index" json:"integral_name"`
	IntegralType   int8       `orm:"index" json:"integral_type"` // 1:邀请积分; 2:购买积分; 3: 管理奖励
	IntegralSource string     `orm:"size(128);index" json:"integral_source"`
	IntegralAmount float64    `orm:"default(0);digits(22);decimals(8)" json:"integral_amount"`
	OrderNumber    string     `orm:"size(128);index" json:"order_number"`
	SourceUserId   int64      `orm:"index" json:"source_user_id"`
}

func (this *IntegralRecord) TableName() string {
	return common.TableName("integral_record")
}

func (this *IntegralRecord) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *IntegralRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *IntegralRecord) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *IntegralRecord) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *IntegralRecord) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *IntegralRecord) InsertDb (db orm.Ormer) error {
	if _, err := db.Insert(this); err != nil {
		return err
	}
	return nil
}

