package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Banner struct {
	BaseModel
	Id           int64     `json:"id" form:"id"`
	Avator       string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"avator"`
	Url          string    `orm:"size(512);index" json:"url"`
	IsDispay     int8      `orm:"default(0)" json:"is_dispay"`   // 0 不显示 1 显示
}

func (this *Banner) TableName() string {
	return common.TableName("goods_order")
}

func (this *Banner) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *Banner) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Banner) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Banner) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Banner) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func GetBannerList() ([]*Banner, int, error) {
	var banner_list []*Banner
	if _, err := orm.NewOrm().QueryTable(Banner{}).Filter("IsDispay", 1).All(&banner_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return banner_list, types.ReturnSuccess, nil
}
