package goods

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type GoodsCategoryCheck struct {
	types.PageSizeData
	FirstLevetCatId int64 `json:"first_levet_cat_id"`
	LastLevelCatId  int64 `json:"last_level_cat_id"`
}

func (this GoodsCategoryCheck) GoodsCategoryCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.FirstLevetCatId < 0 {
		return types.ParamLessZero, errors.New("商品一级分类ID小于0")
	}
	if this.LastLevelCatId < 0 {
		return types.ParamLessZero, errors.New("商品二级分类ID小于0")
	}
	return types.ReturnSuccess, nil
}

type GoodsDetailCheck struct {
	GoodsId  int64 `json:"goods_id"`
}

func (this GoodsDetailCheck) GoodsDetailCheckParamValidate() (int, error) {
	if this.GoodsId < 0 {
		return types.ParamLessZero, errors.New("商品ID不能小于0")
	}
	return types.ReturnSuccess, nil
}


type MerchantGoodsListCheck struct {
	types.PageSizeData
	MerchantId  int64 `json:"merchant_id"`
	QueryWay    int8 `json:"query_way"` // 0:全部；1:活动优选；2:爆款产品
}

func (this MerchantGoodsListCheck) MerchantGoodsListCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.MerchantId <= 0 {
		return types.ParamLessZero, errors.New("商家ID不能小于等于0")
	}
	if this.QueryWay < 0 || this.QueryWay > 2 {
		return types.InvalidVerifyWay, errors.New("无效的查询方式")
	}
	return types.ReturnSuccess, nil
}