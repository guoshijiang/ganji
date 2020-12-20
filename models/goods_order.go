package models

import (
"ganji/common"
"github.com/astaxie/beego/logs"
"github.com/astaxie/beego/orm"
"time"
)

type GoodsOrder struct {
	BaseModel
	Id            int64      `json:"id"`
	GoodsId       int64      `orm:"size(64);index" json:"goods_id"`                         // 商品 ID
	AddressId     int64      `orm:"size(64);index" json:"address_id"`                       // 地址 ID
	GoodsTitle    string     `orm:"size(64)" json:"goods_title"`                            // 商品标题
	GoodsName    string    `orm:"size(512);index" json:"goods_name" form:"goods_name"`      // 产品名称
	Logo          string     `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品Logo
	UserId        int64      `orm:"size(64);index" json:"user_id"`                          // 购买用户
	BuyNums       int64      `orm:"default(0)" json:"buy_nums"`                             // 购买数量
	PayWay        int8       `orm:"index" json:"pay_way"`                                   // 0:积分兑换，1:账户余额支付，2:微信支付；3:支付宝支付
	PayAmount     float64    `orm:"default(0);digits(22);decimals(8)" json:"pay_amount"`    // 支付金额
	SendIntegral  float64    `orm:"default(1);digits(22);decimals(8)" json:"send_integral"` // 赠送积分
	OrderNumber   string     `orm:"size(64);index" json:"order_number"`                     // 订单号
	ShipNumber    string     `orm:"size(64);index;default('')" json:"ship_number"`          // 运单号
	OrderStatus   int8       `orm:"index" json:"order_status"`                              // 0: 未支付，1: 支付中，2：支付成功；3：支付失败 4：已发货；5：已经收货
	FailureReason string     `json:"failure_reason"`                                        // 失败原因
	PayAt         *time.Time `orm:"type(datetime);null" json:"pay_at"`                      // 支付时间
	DealMerchant  string     `orm:"default('')" json:"deal_user"`                           // 处理商家
	DealAt        time.Time  `orm:"null;type(datetime);" json:"deal_at"`                    // 处理时间
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
