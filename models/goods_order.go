package models

import (
"ganji/common"
	"ganji/types"
	type_order "ganji/types/order"
	"github.com/astaxie/beego/logs"
"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type GoodsOrder struct {
	BaseModel
	Id            int64      `json:"id"`
	GoodsId       int64      `orm:"size(64);index" json:"goods_id"`                         // 商品 ID
	MerchantId    int64      `orm:"size(64);index" json:"merchant_id"`                      // 商户 ID
	AddressId     int64      `orm:"size(64);index" json:"address_id"`                       // 地址 ID
	GoodsTypes    string     `orm:"size(512)" json:"goods_types"`                           // 商品属性
	GoodsTitle    string     `orm:"size(64)" json:"goods_title"`                            // 商品标题
	GoodsName     string     `orm:"size(512);index" json:"goods_name" form:"goods_name"`    // 产品名称
	Logo          string     `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品Logo
	UserId        int64      `orm:"size(64);index" json:"user_id"`                          // 购买用户
	BuyNums       int64      `orm:"default(0)" json:"buy_nums"`                             // 购买数量
	PayWay        int8       `orm:"index" json:"pay_way"`                                   // 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付; 4:未知支付方式
	PayAmount     float64    `orm:"default(0);digits(22);decimals(8)" json:"pay_amount"`    // 支付金额
	PayCoupon	  float64  	 `orm:"default(0);digits(22);decimals(8)" json:"pay_coupon"`    // 优惠券抵扣金额
	PayIntegral   float64  	 `orm:"default(0);digits(22);decimals(8)" json:"pay_integral"`  // 支付的积分
	SendIntegral  float64    `orm:"default(1);digits(22);decimals(8)" json:"send_integral"` // 赠送积分
	OrderNumber   string     `orm:"size(64);index" json:"order_number"`                     // 订单号
	Logistics	  string     `orm:"size(64);index;default('')" json:"logistics"`            // 物流公司
	ShipNumber    string     `orm:"size(64);index;default('')" json:"ship_number"`          // 运单号
	OrderStatus   int8       `orm:"index" json:"order_status"`                              // 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已完成
	FailureReason string     `json:"failure_reason"`                                        // 失败原因
	PayAt         *time.Time `orm:"type(datetime);null" json:"pay_at"`                      // 支付时间
	DealMerchant  string     `orm:"default('')" json:"deal_user"`                           // 处理商家
	DealAt        time.Time  `orm:"null;type(datetime);" json:"deal_at"`                    // 处理时间
	IsCancle      int8       `orm:"default(0);index" json:"is_cancle"`                      // 0 正常；1.退货; 2:换货; 3:退货成功; 4:换货成功
	IsComment     int8       `orm:"default(0);index" json:"is_comment"`                     // 0 正常；1.已评价
	BatchId       string     `orm:"size(128);index" json:"batch_id"`                        // 订单批 ID
	IsReward      int8       `orm:"default(0);index" json:"is_reward"`                      // 0 正常；1.已计算直间邀请奖励奖励；2:已计算管理奖励
	IsBuild       int8       `orm:"default(0);index" json:"is_build"`                       // 0 正常；1.已构建
	IsStatic      int8 		 `orm:"default(0);index" json:"is_static"`                      // 0 正常；1.已统计
}

func (this *GoodsOrder) TableName() string {
	return common.TableName("goods_order")
}

func (this *GoodsOrder) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsOrder) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsOrder) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsOrder) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsOrder) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

func (this *GoodsOrder) SearchField() []string {
	return []string{"order_num"}
}


func GetGoodsOrderList(page, pageSize int, user_id int64, status int8) ([]*GoodsOrder, int64, error) {
	offset := (page - 1) * pageSize
	gds_order_list := make([]*GoodsOrder, 0)
	query := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("UserId", user_id).OrderBy("-id")
	if status >= 0  && status <= 5 {
		query = query.Filter("OrderStatus", status)
	}
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&gds_order_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gds_order_list, total, nil
}

func GetGoodsOrderBatchList(batch_id string, user_id int64) ([]*GoodsOrder, int, error) {
	gds_order_list := make([]*GoodsOrder, 0)
	_, err := orm.NewOrm().QueryTable(GoodsOrder{}).
		Filter("UserId", user_id).
		Filter("BatchId", batch_id).
		Filter("OrderStatus", 0).All(&gds_order_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gds_order_list, types.ReturnSuccess, nil
}

func GetGoodsOrderDetail(id int64) (*GoodsOrder, int, error) {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("Id", id).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_dtl, types.ReturnSuccess, nil
}


func GetGoodsOrderByOrderNumber(order_number string) (*GoodsOrder, int, error) {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("OrderNumber", order_number).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_dtl, types.ReturnSuccess, nil
}


// 1.退货,资金返回钱包账号; 2:退货,资金原路返回; 3:换货
func ReturnGoodsOrder(oret type_order.ReturnGoodsOrderCheck) (*GoodsOrder, int, error) {
	var order_dtl GoodsOrder
	if err := orm.NewOrm().QueryTable(GoodsOrder{}).Filter("Id", oret.OrderId).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	if oret.FundRet == 1 || oret.FundRet == 2 {
		order_dtl.IsCancle = 1
	}
	if oret.FundRet == 3 {
		order_dtl.IsCancle = 2
	}
	err := order_dtl.Update()
	if err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	order_p := OrderProcess{
		OrderId: order_dtl.Id,
		MerchantId: order_dtl.MerchantId,
		UserId: order_dtl.UserId,
		AddressId: order_dtl.AddressId,
		GoodsId: order_dtl.GoodsId,
		RetGoodsRs: oret.RetGoodsRs,
		QsDescribe: oret.QsDescribe,
		QsImgOne: oret.QsImgOne,
		QsImgTwo: oret.QsImgTwo,
		QsImgThree: oret.QsImgThree,
		Process: 0,
		LeftTime: 604800,
		IsRecvGoods: oret.IsRecvGoods,    // 0:未收到货物，1:已经收到货物
		FundRet: oret.FundRet,
	}
	err, _ = order_p.Insert()
	if err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_dtl, types.ReturnSuccess, nil
}


func GetOrdetListByMet(merchant_id int64, db orm.Ormer) ([]*GoodsOrder, error) {
	var order_list []*GoodsOrder
	if _, err := db.QueryTable(GoodsOrder{}).Filter("MerchantId", merchant_id).Filter("IsStatic", 0).All(&order_list); err != nil {
		return nil, errors.New("数据库操作错误")
	}
	return order_list, nil
}


func GetRewardGoodsOrderList(db orm.Ormer, is_reward int8) ([]*GoodsOrder, error) {
	gds_order_list := make([]*GoodsOrder, 0)
	_, err := db.QueryTable(GoodsOrder{}).Filter("IsReward", is_reward).All(&gds_order_list)
	if err != nil {
		return nil, errors.New("查询数据库失败")
	}
	return gds_order_list, nil
}


func GetUnbuildTreeGoodsOrders(db orm.Ormer) ([]*GoodsOrder, error) {
	gds_order_list := make([]*GoodsOrder, 0)
	_, err := db.QueryTable(GoodsOrder{}).Filter("IsBuild", 0).All(&gds_order_list)
	if err != nil {
		return nil, errors.New("查询数据库失败")
	}
	return gds_order_list, nil
}