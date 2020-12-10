package models

import "ganji/common"

type Merchant struct {
	BaseModel
	Id           int64     `json:"id" form:"id"`
	MerchantName string    `orm:"size(512);index" json:"merchant_name"`  // 商家名称
	Logo         string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"`   // 商家 Logo
	Address      string    `orm:"size(512);index" json:"address"`  // 店铺地址
	GoodsNum     int64     `json:"goods_num"`  // 商品总数
	ShopLevel    int8      `json:"shop_level"` // 店铺等级
	ShopServer   int8      `json:"shop_server"` // 店铺服务
}

func (this *Merchant) TableName() string {
	return common.TableName("merchant")
}
