package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type GoodsCat struct {
	BaseModel
	Id           int64     `orm:"column(id);auto;size(11)" json:"id" form:"id"`
	CatLevel     int8      `orm:"default(1)" json:"cat_level"` // 分类级别
	FatherCatId  int64     `json:"father_cat_id"`              // 父级分类 ID
	Name         string    `orm:"size(512);index" json:"name"` // 分类标题
	Icon         string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"icon"` // 分类Icon
	IsDispay     int8      `orm:"default(0)" json:"is_dispay"`   // 0 显示 1 不显示
}

func (this *GoodsCat) TableName() string {
	return common.TableName("goods_cat")
}

func (this *GoodsCat) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsCat)SearchField() []string{
	return []string{"name"}
}

func (this *GoodsCat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCat) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCat) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsCat) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func GetOneLevelCategoryList() ([]*GoodsCat, int, error) {
	var goods_cat_list []*GoodsCat
	if _, err := orm.NewOrm().QueryTable(GoodsCat{}).
		Filter("IsDispay", 1).
		Filter("CatLevel", 1).
		OrderBy("-id").All(&goods_cat_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_cat_list, types.ReturnSuccess, nil
}

func GetSecodLevelCategoryList(my_id int64) ([]*GoodsCat, int, error) {
	var goods_cat_list []*GoodsCat
	if _, err := orm.NewOrm().QueryTable(GoodsCat{}).
		Filter("FatherCatId", my_id).
		Filter("IsDispay", 1).
		Filter("CatLevel", 2).
		OrderBy("CreatedAt").All(&goods_cat_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_cat_list, types.ReturnSuccess, nil
}