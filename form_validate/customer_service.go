package form_validate


type CustomerServiceForm struct {
	Id			   int64		 `form:"id"`
	UserName       string        `form:"user_name"`
	Phone          string        `form:"phone"`
	WeiChat        string        `form:"wei_chat"`
	WcQrcode       string        `form:"wc_qrcode"`
	Type           int8          `form:"type"`
	IsCreate 	   int    	 	 `form:"_create"`
}

func (*CustomerServiceForm) Messages() {
	//todo
}