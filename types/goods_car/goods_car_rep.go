package goods_car

type GoodsCarList struct {
	MerchantId    int64  `json:"merchant_id"`
	MerchantName  string `json:"merchant_name"`
	GoodsCarId    int64  `json:"goods_car_id"`
	GoodsId       int64  `json:"goods_id"`
	GoodsLogo     string `json:"logo"`
	GoodsTitle    string `json:"goods_title"`
	GoodsMark     string `json:"goods_marks"`
	GoodsName     string `json:"goods_name"`
	GoodsPrice    float64 `json:"goods_price"`
	UserId        int64  `json:"user_id"`
	BuyNums       int64  `json:"buy_nums"`
	PayAmount     float64 `json:"pay_amount"`
}