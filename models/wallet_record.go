package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type WalletRecord struct {
	BaseModel
	Id           int64       `json:"id"`
	UserId       int64       `orm:"default(0);index" json:"user_id"`
	Amount       float64     `orm:"default(0);digits(32);decimals(8)" json:"amount"`  // 充值金额
	OrderNumber  string      `orm:"size(256)" json:"order_number"`                    // 交易订单号
	Type         int8        `orm:"index" json:"type"`                                // 0:充值；1:提现 2:积分兑换 3:消费
	Source       int8        `orm:"default(0);index" json:"source"`                   // 0：支付宝 1:微信; 2:积分兑换
	Status       int8        `orm:"default(0);index" json:"status"`                   // 0:入账成功；2: 入账失败
}

func (this *WalletRecord) TableName() string {
	return common.TableName("wallet_record")
}

func (this *WalletRecord) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *WalletRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func GetWalletRecordList(page, pageSize int, user_id int64) ([]*WalletRecord, int64, error) {
	offset := (page - 1) * pageSize
	ig_trade_list := make([]*WalletRecord, 0)
	query := orm.NewOrm().QueryTable(WalletRecord{}).Filter("UserId", user_id)
	total, _ := query.Count()
	_, err := query.OrderBy("-CreatedAt").Limit(pageSize, offset).All(&ig_trade_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return ig_trade_list, total, nil
}


func GetWalletRecordDetail(id int64) (*WalletRecord, int, error) {
	var wdtl WalletRecord
	if err := orm.NewOrm().QueryTable(WalletRecord{}).Filter("Id", id).RelatedSel().One(&wdtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &wdtl, types.ReturnSuccess, nil
}


