package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type CustomerService struct {
	BaseModel
	Id             int64         `json:"id"`
	UserName       string        `orm:"size(128)" json:"user_name"`
	Phone          string        `orm:"size(64);index" json:"phone"`
	WeiChat        string        `orm:"size(64);index" json:"wei_chat"`
	WcQrcode       string        `orm:"size(150);default(/static/upload/default/user-default-60x60.png)"`
	Type           int8          `orm:"default(1);index" json:"type"`   // 1:申请入住  2:客户服务
}

func (this *CustomerService) TableName() string {
	return common.TableName("customer_service")
}

func (this *CustomerService) SearchField() []string {
	return []string{"user_name"}
}

func (this *CustomerService) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *CustomerService) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *CustomerService) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *CustomerService) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func GetCustomerServicel(type_w int8) (*CustomerService, int, error) {
	var customer_s CustomerService
	if err := orm.NewOrm().QueryTable(CustomerService{}).Filter("Type", type_w).OrderBy("-id").Limit(1).RelatedSel().One(&customer_s); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &customer_s, types.ReturnSuccess, nil
}
