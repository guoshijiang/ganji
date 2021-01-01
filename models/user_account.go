package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type UserAccount struct {
	BaseModel
	Id          int64       `json:"id"`
	UserId      int64       `orm:"index" json:"user_id"`
	AcountType  int8        `orm:"default(0)" json:"acount_type"`          // 0 支付宝； 1:微信; 2:银行卡; 3:平台账户
	AccountName string      `orm:"size(128);index" json:"account_name"`    // 账号名称; 银行名称
	UserName    string 		`orm:"size(128);index" json:"user_name"`       // 用户名字; 银行开户名字
	CardNum     string      `orm:"size(128);index" json:"card_num"`        // 账号；银行卡号
	Address     string      `orm:"size(128);index" json:"address"`         // 开户行地址
	IsInvalid   int8        `orm:"default(0)" json:"is_invalid"`           // 0 激活； 1:禁用
}

func (this *UserAccount) TableName() string {
	return common.TableName("user_account")
}

func (this *UserAccount) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *UserAccount) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *UserAccount) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *UserAccount) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func AccountExist(user_id int64, account_type int8) bool {
	return  orm.NewOrm().QueryTable(UserAccount{}).Filter("UserId", user_id).Filter("AcountType", account_type).Exist()
}

func GetAccountById(account_id int64) (*UserAccount, int, error) {
	var account UserAccount
	if err := orm.NewOrm().QueryTable(UserAccount{}).Filter("Id", account_id).RelatedSel().One(&account); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &account, types.ReturnSuccess, nil
}

func GetUserAccountList(user_id int64) ([]*UserAccount, int, error) {
	var account_list []*UserAccount
	if _, err := orm.NewOrm().QueryTable(UserAccount{}).Filter("UserId", user_id).All(&account_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return account_list, types.ReturnSuccess, nil
}