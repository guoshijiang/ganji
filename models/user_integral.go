package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)


type UserIntegral struct {
	BaseModel
	Id             int64      `json:"id"`
	UserId         int64      `orm:"size(64);index" json:"user_id"`
	IntegralName   string     `orm:"size(128);index" json:"integral_name"`
	TotalIg        float64    `orm:"default(0);digits(22);decimals(8)" json:"total_ig"`     // 总的积分
	UsedIg         float64    `orm:"default(0);digits(22);decimals(8)" json:"used_ig"`      // 已赠送LSDT积分
	TodayIg        float64    `orm:"default(0);digits(22);decimals(8)" json:"today_ig"`     // 今日已赠送LSDT积分
}


func (this *UserIntegral) TableName() string {
	return common.TableName("user_integral")
}

func (this *UserIntegral) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *UserIntegral) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *UserIntegral) UpdateDb(db orm.Ormer) error {
	if _, err := db.Update(this); err != nil {
		return err
	}
	return nil
}

func (this *UserIntegral) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *UserIntegral) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (Self *UserIntegral) Insert() error {
	if _, err := orm.NewOrm().Insert(Self); err != nil {
		return err
	}
	return nil
}

func GetIntegralByUserId(user_id int64) (*UserIntegral, error) {
	var user_ig UserIntegral
	err := user_ig.Query().Filter("UserId", user_id).Limit(1).One(&user_ig)
	if err != nil {
		return nil, err
	}
	return &user_ig, nil
}


