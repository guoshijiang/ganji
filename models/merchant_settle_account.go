package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type MerchantSettleAccount struct {
	BaseModel
	Id         int64  `json:"id"`
	MerchantId int64  `orm:"index" json:"merchant_id"` // 商户 ID
	AcctSeq    string `orm:"size(150);index" json:"acct_seq"`    // 账户序号：填写示范：支付宝一
	AcctType   int8   `orm:"default(0);index" json:"acct_type"` // 0:支付宝； 1:微信
	AcctName   string `orm:"size(150);index" json:"acct_name"`   // 账号名称
	RealName   string `orm:"size(150);index" json:"real_name"`   // 账号持有人真实名字
	Qrcode     string `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"qrcode"` // 账号收款码
}

func (this *MerchantSettleAccount) TableName() string {
	return common.TableName("merchant_settle_account")
}

func (this *MerchantSettleAccount) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *MerchantSettleAccount) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *MerchantSettleAccount) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *MerchantSettleAccount) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *MerchantSettleAccount) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

