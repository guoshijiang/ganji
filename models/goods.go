package models


import (
	"ganji/common"
)

type Goods struct {
	BaseModel
	Id           int64     `json:"id" form:"id"`
	GoodsCatId   int64     `json:"goods_cat_id"`
	Title        string    `orm:"size(512);index" json:"title" form:"title"`                                                  // 商品标题
	Logo         string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"`   // 商品封面
	TotalAmount  int64     `orm:"default(150000)" json:"total_amount" form:"total_amount"`                                    // 商品总量
	LeftAmount   int64     `orm:"default(150000)" json:"left_amount" form:"left_amount"`                                      // 剩余商品总量
	GoodsPrice   float64   `orm:"default(1);digits(22);decimals(8)" json:"goods_price" form:"goods_price"`
	GoodsIntegral float64  `orm:"default(1);digits(22);decimals(8)" json:"goods_integral"`
	GoodsName    string    `orm:"size(512);index" json:"goods_name" form:"goods_name"`  // 产品名称
	GoodsType    string    `orm:"size(512);index" json:"goods_type" form:"goods_type"`  // 产品型号
	GoodsSize    string    `orm:"size(512);index" json:"goods_size" form:"goods_size"`  // 产品尺寸
	GoodsParams  string    `orm:"type(text)" json:"goods_params" form:"goods_params"`   // 产品参数
	GoodsDetail  string    `orm:"type(text)" json:"goods_detail" form:"goods_detail"`   // 产品详细介绍
	IsDisplay    int8      `orm:"default(0);index" json:"is_display" form:"is_display"` // 0:首页不展示, 1:首页展示
	IsHot        int8      `orm:"default(1);index" json:"is_hot"`                       // 0:热卖 1:非热卖
	IsDiscount   int8      `orm:"default(0);index" json:"is_discount"`                  // 0:不打折，1:打折
	IsIgExchange int8      `orm:"default(0);index" json:"is_ig_exchange"`               // 0:正常，1:可以积分兑换
	Discount     float64   `orm:"default(0);index" json:"discount"`                     // 折扣，取值 0.1-9.9；0代表不打折
	Sale         int8      `orm:"default(1);index" json:"sale" form:"sale"`             // 0:上架 1:下架
}

func (this *Goods) TableName() string {
	return common.TableName("goods")
}
