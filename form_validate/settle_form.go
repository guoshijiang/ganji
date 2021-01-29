package form_validate

import "time"

type SettleForm struct {
	Id                 int64      `form:"id"`
	SettleAccountId    int64      `form:"settle_account_id"`  		// 结算账号ID
	StartSettleTime    time.Time `form:"start_settle_time"`  		// 结算开始日期
	EndSettleTime  	   time.Time `form:"end_settle_time"`    		// 结算结束日期
	MerchantId         int64      `form:"merchant_id"`        		// 商户 ID
	SettleAmount       float64    `form:"settle_amount"` 			// 结算金额
	HandUser           string     `form:"hand_user"`               	// 平台处理账户
	Status             int8       `form:"status"`

}