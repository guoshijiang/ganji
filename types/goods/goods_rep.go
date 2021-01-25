package goods

type CategoryGoodsRet struct {
	GoodsId       int64  `json:"goods_id"`
	GoodsMark     string `json:"goods_mark"`
	Title         string `json:"title"`
	Logo          string `json:"logo"`
	GoodsPrice    float64 `json:"goods_price"`
	GoodsDisPrice float64 `json:"goods_discount_price"`
	GoodsIntegral float64 `json:"goods_integral"`
	SendIntegral  float64 `json:"send_integral"`
	LeftTime      int64   `json:"left_time"`
	IsHot          int8  `json:"is_hot"`
	IsDiscount     int8 `json:"is_discount"`
	IsIgSend       int8 `json:"is_ig_send"`
	IsGroup        int8 `json:"is_group"`
	IsIntegral     int8 `json:"is_integral"`
}

type GoodsImagesRet struct {
	GoodsImgId  int64  `json:"goods_img_id"`
	ImageUrl    string `json:"image_url"`
}

type GoodsTypeRet struct {
	GoodsTypeId   int64  `json:"goods_type_id"`
	GoodsTypeName string `json:"goods_type_name"`
}