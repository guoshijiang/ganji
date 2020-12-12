package goods

type CategoryGoodsRet struct {
	GoodsId       int64  `json:"goods_id"`
	GoodsMark     string `json:"goods_mark"`
	Title         string `json:"title"`
	Logo          string `json:"logo"`
	GoodsPrice    float64 `json:"goods_price"`
	GoodsDisPrice float64 `json:"goods_discount_price"`
	LeftTime      int64   `json:"left_time"`
}