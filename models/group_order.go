package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type GroupOrder struct {
	BaseModel
	Id                int64      `json:"id"`
	GoodsId           int64      `orm:"size(64);index" json:"goods_id"`                         // 商品 ID
	AddressId         int64      `orm:"size(64);index" json:"address_id"`                       // 地址 ID
	GoodsTitle        string     `orm:"size(64)" json:"goods_title"`                            // 商品标题
	Logo              string     `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品Logo
	UserId            int64      `orm:"size(64);index" json:"user_id"`                          // 购买用户
	BuyNums           int64      `orm:"default(0)" json:"buy_nums"`                             // 购买数量
	OrderAmount       float64    `orm:"default(0);digits(22);decimals(8)" json:"order_amount"`  // 订单金额
	NeedBargainTimes  int64      `orm:"default(0)" json:"need_bargain_times"`                   // 需要砍价次数
	BargainTimes      int64      `orm:"default(0)" json:"bargain_times"`                        // 砍价次数
	BargainPrice      float64    `orm:"default(0);digits(22);decimals(8)" json:"bargain_price"` // 价格
	OrderNumber       string     `orm:"size(64);index" json:"order_number"`                     // 订单号
	ShipNumber        string     `orm:"size(64);index;default('')" json:"ship_number"`          // 运单号
	OrderStatus       int8       `orm:"index" json:"order_status"`                              // 0: 砍价中，1: 砍价成功，2:已发货 3：已收货
	FailureReason     string     `json:"failure_reason"`                                        // 失败原因
	PayAt             *time.Time `orm:"type(datetime);null" json:"pay_at"`                      // 支付时间
	DealMerchant      string     `orm:"default('')" json:"deal_user"`                           // 处理商家
	DealAt            time.Time  `orm:"null;type(datetime);" json:"deal_at"`                    // 处理时间
}

func (this *GroupOrder) TableName() string {
	return common.TableName("group_order")
}

func (this *GroupOrder) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GroupOrder) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GroupOrder) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GroupOrder) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GroupOrder) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

