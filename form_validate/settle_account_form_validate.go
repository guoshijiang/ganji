package form_validate

type SettleAccountForm struct {
	Id             	int64     		`form:"id"`
	MerchantId     	int64     		`form:"merchant_id"` 	// 商品ID
	IsCreate 	   	int    	 		`form:"_create"`
	AcctSeq    		string 			`form:"acct_seq"`    	// 账户序号：填写示范：支付宝一
	AcctType   		int8   			`form:"acct_type"`   	// 0:支付宝； 1:微信
	AcctName   		string 			`form:"acct_name"`   	// 账号名称
	RealName   		string 			`form:"real_name"`   	// 账号持有人真实名字
	Qrcode     		string 			`form:"qrcode"`     	// 账号收款码
}