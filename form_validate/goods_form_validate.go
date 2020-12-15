package form_validate

import "github.com/gookit/validate"

type GoodsForm struct {
	Id             int64     `form:"id"`
	GoodsCatId     int64     `form:"goods_cat_id"`					// 商品所属一级分类ID
	GoodsLastCatId int64     `form:"goods_level_cat_id"`            // 商品所属最后一级分类ID
	GoodsMark      string    `form:"goods_mark"`    				// 商品备注
	Serveice       string    `form:"serveice"`      				// 服务说明
	CalcWay        int8      `form:"calc_way"`     					// 0:按件计量 1:按近计量
	MerchantId     int64     `form:"merchant_id"`                   // 商品所属商家ID
	Title          string    `form:"title"`         				// 商品标题
	Logo           string    `form:"logo"`   						// 商品封面
	TotalAmount    int64     `form:"total_amount"`                 	// 商品总量
	LeftAmount     int64     `form:"left_amount"`                   // 剩余商品总量
	GoodsPrice     float64   `form:"goods_price"`                   // 商品价格
	GoodsDisPrice  float64   `form:"goods_discount_price"`          // 商品折扣价格
	GoodsIntegral  float64   `form:"goods_integral"`                // 购买商品获得积分
	GoodsName      string    `form:"goods_name"`  					// 产品名称
	GoodsParams    string    `form:"goods_params"`   				// 产品参数
	GoodsDetail    string    `form:"goods_detail"`   				// 产品详细介绍
	Discount       float64   `form:"discount"`       				// 折扣，取值 0.1-9.9；0代表不打折
	Sale           int8      `form:"sale"`             				// 0:上架 1:下架
	IsDisplay      int8      `form:"is_display"` 					// 0:首页不展示, 1:首页展示
	IsHot          int8      `form:"is_hot"`                       	// 0:非爆款产品 1:爆款产品
	IsDiscount     int8      `form:"is_discount"`                  	// 0:不打折，1:打折活动产品
	IsIgExchange   int8      `form:"is_ig_exchange"`               	// 0:正常，1:可以积分兑换
	IsGroup        int8      `form:"is_group"`                     	// 0:非拼购产品 1:拼购产品
	IsIntegral     int8      `form:"is_integral"`                  	// 0:非积分兑换产品 1:积分兑换产品
	LeftTime       int64     `form:"left_time"`                    	// 限时产品剩余时间
	IsLimitTime    int8      `form:"is_limit_time"`                	// 0:不是限时产品 1:是限时
	IsCreate 	   int    	 `form:"_create"`

}

func (*GoodsForm) Messages() map[string]string {
	return validate.MS{
		"MerchantName.required":        "名称不能为空.",
		"MerchantIntro.required": 		"介绍不能为空.",
		"Address.required": 			"地址不能为空.",
		"ShopLevel.int8":          		"请填写等级",
	}
}