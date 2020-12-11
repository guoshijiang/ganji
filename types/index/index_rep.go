package index

type IndexBannerRet struct {
	BannerId int64   `json:"banner_id"`
	BannerImg string `json:"banner_img"`
	BannerUrl string `json:"banner_url"`
}

type IndexCatRet struct {
	CatId   int64   `json:"cat_id"`
	CatName string `json:"cat_name"`
	CatIcon string `json:"cat_icon"`
}

type IndexGoodsBuyRet struct {
	GoodsId       int64  `json:"goods_id"`
	GoodsMark     string `json:"goods_mark"`
	Title         string `json:"title"`
	Logo          string `json:"logo"`
	GoodsPrice    float64 `json:"goods_price"`
	GoodsDisPrice float64 `json:"goods_discount_price"`
}

type IndexLimitTimeGoodsRet struct {
	GoodsId       int64  `json:"goods_id"`
	GoodsMark     string `json:"goods_mark"`
	Title         string `json:"title"`
	Logo          string `json:"logo"`
	GoodsPrice    float64 `json:"goods_price"`
	GoodsDisPrice float64 `json:"goods_discount_price"`
	LeftTime      int64   `json:"left_time"`
}

type IndexDownGoodsListRet struct {
	GoodsId       int64  `json:"goods_id"`
	GoodsMark     string `json:"goods_mark"`
	Title         string `json:"title"`
	Logo          string `json:"logo"`
	GoodsPrice    float64 `json:"goods_price"`
	GoodsDisPrice float64 `json:"goods_discount_price"`
	GoodsIntegral  float64 `json:"goods_integral"`
}



