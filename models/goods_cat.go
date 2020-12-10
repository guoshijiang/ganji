package models

import (
	"ganji/common"
)

type GoodsCat struct {
	BaseModel
	Id           int64     `json:"id" form:"id"`
	Name         string    `orm:"size(512);index" json:"name"` // 分类标题
}

func (this *GoodsCat) TableName() string {
	return common.TableName("goods_cat")
}