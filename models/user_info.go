package models

import (
	"ganji/common"
	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	BaseModel
	Id         int64     `json:"id"`
	UserId     int64     `orm:"index" json:"user_id"`
	RealName   string    `orm:"default(ganji);size(15);index" json:"real_name"`
	WeiChat    string 	 `orm:"default(ganji);size(15);index" json:"wei_chat"`
	QQ         string 	 `orm:"default(ganji);size(15);index" json:"qq"`
	Sex        int8      `orm:"default(0);index"` // 0: 男，1: 女  3:未知
}

func (this *UserInfo) TableName() string {
	return common.TableName("user_info")
}

func (this *UserInfo) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *UserInfo) Read(fields ...string) error {
	return nil
}

func (this *UserInfo) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *UserInfo) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *UserInfo) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}


func (u *UserInfo) GetUserInfoByUserId(user_id int64) (UserInfo, error) {
	var user_info UserInfo
	err := user_info.Query().Filter("UserId", user_id).Limit(1).One(&user_info)
	if err != nil {
		return user_info, err
	}
	return user_info, nil
}



