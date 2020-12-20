package order

import "time"

type OrderListRet struct {
	MerchantId int64    `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	OrderId  int64      `json:"order_id"`
	GoodsName string    `json:"goods_name"`
	GoodsPrice float64  `json:"goods_price"`
	OrderStatus int8    `json:"order_status"`
	BuyNums     int64   `json:"buy_nums"`
	PayAmount   float64 `json:"pay_amount"`
}


type OrderDetailRet struct {
	OrderId  int64      `json:"order_id"`
	ShipLogo   string   `json:"ship_logo"`
	ShipInfo   string   `json:"ship_info"`
	RecUser    string   `json:"rec_user"`
	RecPhone   string   `json:"rec_phone"`
	RecAddress string   `json:"rec_address"`
	MerchantId int64    `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	GoodsName string    `json:"goods_name"`
	GoodsPrice float64  `json:"goods_price"`
	OrderStatus int8    `json:"order_status"`
	BuyNums     int64   `json:"buy_nums"`
	PayAmount   float64 `json:"pay_amount"`
	ShipFee     float64 `json:"ship_fee"`
	Coupons     float64 `json:"coupons"`
	PayWay      int8    `json:"pay_way"`
	OrderNumber string  `json:"order_number"`
	PayTime     *time.Time `json:"pay_time"`
	CreateTime  time.Time `json:"create_time"`
}
