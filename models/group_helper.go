package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GroupHelper struct {
	BaseModel
	Id                int64    `json:"id"`
	GroupOrderId      int64   `json:"group_order_id"`
	BuyUserId         int64   `json:"buy_user_id"`
	HelperUserId      int64   `json:"helper_user_id"`
}

func (this *GroupHelper) TableName() string {
	return common.TableName("group_helper")
}

func (this *GroupHelper) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GroupHelper) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GroupHelper) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GroupHelper) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GroupHelper) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}


func ExistOrderByOdrUser(order_id, user_id int64) bool {
	return orm.NewOrm().QueryTable(GroupHelper{}).Filter("GroupOrderId", order_id).Filter("HelperUserId", user_id).Exist()
}