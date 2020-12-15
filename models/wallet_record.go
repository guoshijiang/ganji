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
	Type         int8        `orm:"index" json:"type"`                                // 0:充值； 1:提现 2:消费
	Source       int8        `orm:"default(0);index" json:"source"`                   // 0：支付宝 1:微信; 2:银行卡
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

func (this *WalletRecord) GetWalletDepositList(asset_name string, page int64, page_size int64) ([]*WalletRecord, int, error) {
	var wdt []*WalletRecord
	filter := orm.NewOrm().QueryTable(&WalletRecord{}).Filter("UserId", this.UserId)
	if this.Status > 0 {
		filter = filter.Filter("Status", this.Status)
	}
	total, err := filter.Count()
	_, err = filter.Limit(page_size, page_size*(page-1)).All(&wdt)
	if err != nil {
		return nil, types.ReturnSuccess, errors.Wrap(err, "query deposit list fail")
	}
	return wdt, int(total), nil
}

func (this *WalletRecord) GetWalletDepositById(deposit_id int64) (*WalletRecord, int, error) {
	var wdd WalletRecord
	err := orm.NewOrm().QueryTable(&WalletRecord{}).Filter("Id", deposit_id).One(&wdd)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("Query database error")
	}
	return &wdd, types.ReturnSuccess, nil
}

