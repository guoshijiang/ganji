package models


import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type OrderProcess struct {
	BaseModel
	Id            int64      `json:"id"`
	OrderId       int64      `orm:"size(64);index" json:"order_id"`                       // 商品 ID
	MerchantId    int64      `orm:"size(64);index" json:"merchant_id"`                    // 商户 ID
	AddressId     int64      `orm:"size(64);index" json:"address_id"`                     // 地址 ID
	GoodsId       int64      `orm:"size(64);index" json:"goods_id"`                       // 地址 ID
	// 0:等待卖家确认; 1:卖家已同意; 2:卖家拒绝; 3:等待买家邮寄; 4:等待卖家收货; 5:卖家已经发货; 6:等待买家收货; 7:已完成
	Process       int8       `orm:"default(0);index" json:"process"`                      // 订单退换货情况
}

func (this *OrderProcess) TableName() string {
	return common.TableName("order_process")
}

func (this *OrderProcess) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *OrderProcess) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *OrderProcess) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *OrderProcess) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *OrderProcess) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

func (this *OrderProcess) SearchField() []string {
	return []string{"order_num"}
}

