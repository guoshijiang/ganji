package models

import "ganji/common"

type GoodsImage struct {
	BaseModel
	Id           int64     `json:"id" form:"id"`
	GoodsId      int64     `json:"goods_id"`
	Image        string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品图片
}

func (this *GoodsImage) TableName() string {
	return common.TableName("goods_image")
}

