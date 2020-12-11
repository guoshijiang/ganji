package models

import (
	"ganji/common"
)

type GoodsCat struct {
	BaseModel
	Id           int64     `json:"id" form:"id"`
	CatLevel     int8      `orm:"default(1)" json:"cat_level"` // 分类级别
	FatherCatId  int64     `json:"father_cat_id"`              // 父级分类 ID
	Name         string    `orm:"size(512);index" json:"name"` // 分类标题
	Icon         string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"` // 分类Icon
	IsDispay     int8      `orm:"default(0)" json:"is_dispay"`   // 0 不显示 1 显示
}

func (this *GoodsCat) TableName() string {
	return common.TableName("goods_cat")
}