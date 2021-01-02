package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type GroupHelper struct {
	BaseModel
	Id                int64   `json:"id"`
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


func GetGroupHelpUserListByOrderId(order_id int64) ([]*GroupHelper, int, error) {
	var group_hlp []*GroupHelper
	_, err := orm.NewOrm().QueryTable(GroupHelper{}).Filter("GroupOrderId", order_id).OrderBy("-CreatedAt").All(&group_hlp)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return group_hlp, types.ReturnSuccess, nil
}

