package group_order

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type CreateGroupOrderCheck struct {
	GoodsId       int64   `json:"goods_id"`
	AddressId     int64   `json:"address_id"`
	UserId        int64   `json:"user_id"`
}

func (this CreateGroupOrderCheck) CreateGroupOrderCheckParamValidate() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品ID小于等于 0")
	}
	if this.AddressId <= 0 {
		return types.ParamLessZero, errors.New("地址ID小于等于 0")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}

type GroupOrderListCheck struct {
	types.PageSizeData
	UserId    int64   `json:"user_id"`
	QueryWay  int8  `json:"query_way"`  // 1:我的助力订单 2 好友的助力订单
}

func (this GroupOrderListCheck) GroupOrderListCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.QueryWay <= 0 ||  this.QueryWay >= 3 {
		return types.InvalidVerifyWay, errors.New("无效的查询方式")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type OrderDetailCheck struct {
	GroupOrderId int64 `json:"group_order_id"`
}

func (this OrderDetailCheck) OrderDetailCheckParamValidate() (int, error) {
	if this.GroupOrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}


type HelpOrderDCheck struct {
	GroupOrderId int64 `json:"group_order_id"`
	BuyUserId    int64 `json:"buy_user_id"`
	SelfUserId  int64  `json:"self_user_id"`
}

func (this HelpOrderDCheck) HelpOrderDCheckParamValidate() (int, error) {
	if this.GroupOrderId <= 0 {
		return types.ParamLessZero, errors.New("订单 ID小于等于 0")
	}
	if this.BuyUserId <= 0 {
		return types.ParamLessZero, errors.New("购买用户的ID小于等于 0")
	}
	if this.SelfUserId <= 0 {
		return types.ParamLessZero, errors.New("助力用户的ID小于等于 0")
	}
	return types.ReturnSuccess, nil
}