package merchant

type MerchantListRet struct {
	MctName       string `json:"mct_name"`
	MctIntroduce  string `json:"mct_introduce"`
	MctLogo       string `json:"mct_logo"`
	MctWay        int8   `json:"mct_way"`      // 0:自营商家； 1:认证商家  2:普通商家
	ShopLevel     int8   `json:"shop_level"`   // 店铺等级
	ShopServer    int8   `json:"shop_server"`  // 店铺服务
}
