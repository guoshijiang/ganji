package goods_car

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type AddGoodCarCheck struct {
	GoodsId    int64   `json:"goods_id"`
	UserId     int64   `json:"user_id"`
	BuyNums    int64   `json:"buy_nums"`
	PayAmount  float64 `json:"pay_amount"`
	IsDis      int8    `json:"is_dis"`   // 1:非打折商品  2:打折商品
}

func (this AddGoodCarCheck) AddGoodCarCheckParamValidate() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品 ID 不能小于等于 0")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户 ID 不能小于等于 0")
	}
	if this.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量不能小于等于 0")
	}
	if this.PayAmount <= 0 {
		return types.ParamLessZero, errors.New("支付金额不能小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type EditGoodCarCheck struct {
	GoodsId      int64    `json:"goods_id"`
	GoodsCarId   int64    `json:"goods_car_id"`
	UserId       int64    `json:"user_id"`
	BuyNums      int64    `json:"buy_nums"`
	PayAmount    float64  `json:"pay_amount"`
	IsDis        int8     `json:"is_dis"`   // 0:非打折商品  1:打折商品
}


func (this EditGoodCarCheck) EditGoodCarCheckParamValidate() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品 ID 不能小于等于 0")
	}
	if this.GoodsCarId <= 0 {
		return types.ParamLessZero, errors.New("购物车ID 不能小于等于 0")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户 ID 不能小于等于 0")
	}
	if this.BuyNums <= 0 {
		return types.ParamLessZero, errors.New("购买数量不能小于等于 0")
	}
	if this.PayAmount <= 0 {
		return types.ParamLessZero, errors.New("支付金额不能小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type DelGoodCarCheck struct {
	GoodsIds   string  `json:"goods_ids"`
}

func (this DelGoodCarCheck) DelGoodCarCheckParamValidate() (int, error) {
	if this.GoodsIds == "" {
		return types.ParamLessZero, errors.New("商品 ID 数组长度不能小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type GoodCarListCheck struct {
	types.PageSizeData
}

func (this GoodCarListCheck) GoodCarListCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	return types.ReturnSuccess, nil
}