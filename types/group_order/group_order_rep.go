package group_order

import "time"

type GroupOrderListRet struct {
	MerchantId   int64   `json:"merchant_id"`
	MerchantName string  `json:"merchant_name"`
	BuyUserId    int64   `json:"buy_user_id"`
	OrderId      int64   `json:"order_id"`
	GoodsName    string  `json:"goods_name"`
	GoodsLogo    string  `json:"goods_logo"`
	GoodsPrice   float64 `json:"goods_price"`
	OrderStatus  int8    `json:"order_status"`
	GroupNumber  int64   `json:"group_number"`
	HelpNumber   int64   `json:"help_number"`
	IsValid      int8    `json:"is_valid"`
	DeadLime     string  `json:"dead_lime"`
}


type GroupOrderDetailRet struct {
	OrderId    int64    `json:"order_id"`
	GoodsId   int64     `json:"goods_id"`
	BuyUserId  int64   `json:"buy_user_id"`
	RecUser    string   `json:"rec_user"`
	RecPhone   string   `json:"rec_phone"`
	RecAddress string   `json:"rec_address"`
	MerchantId int64    `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	GoodsName string    `json:"goods_name"`
	GoodsLogo string `json:"goods_logo"`
	GoodsPrice float64  `json:"goods_price"`
	OrderStatus int8    `json:"order_status"`
	GroupNumber  int64   `json:"group_number"`
	HelpNumber   int64   `json:"help_number"`
	ShipFee     float64 `json:"ship_fee"`
	Logistics	string  `json:"logistics"`
	ShipNumber  string  `json:"ship_number"`
	OrderNumber string  `json:"order_number"`
	PayAt       *time.Time `json:"pay_at"`
	CreateTime  time.Time `json:"create_time"`
	IsHelp      bool `json:"is_help"`  // 0:没有助力；1:已经助力
}


type HelpUser struct {
	UserId    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	UserPhoto string `json:"user_photo"`
	UserPhone string `json:"user_phone"`
}

type HelpInfoDetailRet struct {
	GoodsId   int64  `json:"goods_id"`
	OrderId   int64  `json:"order_id"`
	GoodsName string `json:"goods_name"`
	GoodsLogo string `json:"goods_logo"`
	GoodsMark string `json:"goods_mark"`
	DeadLine  string `json:"dead_line"`
	OrderStatus int8 `json:"order_status"`
	HelpUsers  []HelpUser `json:"help_users"`
}
