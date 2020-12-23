package wallet_integral

import "time"

type IntegralRecordListRet struct {
	IntegralId     int64     `json:"integral_id"`
	IntegralType   string    `json:"integral_type"`
	IntegralAmount float64   `json:"integral_amount"`
	CnyAmount      float64   `json:"cny_amount"`
	CreateTime     time.Time `json:"create_time"`
}

type IntegralRecordDetailRet struct {
	IntegralId     int64     `json:"integral_id"`
	IntegralType   string    `json:"integral_type"`
	IntegralAmount float64   `json:"integral_amount"`
	CnyAmount      float64   `json:"cny_amount"`
	Fee            float64   `json:"fee"`
	OrderNumber    string    `json:"order_number"`
	Status         int8      `json:"status"`
	CreateTime     time.Time `json:"create_time"`
}


type WalletRecordListRet struct {
	RecordId     int64     `json:"record_id"`
	IntegralType string    `json:"integral_type"`
	TotalAmount  float64   `json:"total_amount"`
	CreateTime   time.Time `json:"create_time"`
}


type WalletRecordDetailRet struct {
	RecordId     int64     `json:"record_id"`
	IntegralType string    `json:"integral_type"`
	IntegralSource string  `json:"integral_source"`
	TotalAmount  float64   `json:"total_amount"`
	OrderNumber  string    `json:"order_number"`
	CreateTime   time.Time `json:"create_time"`
}

