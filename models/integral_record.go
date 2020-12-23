package models


import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type IntegralRecord struct {
	BaseModel
	Id             int64      `json:"id"`
	UserId         int64      `orm:"index" json:"user_id"`
	IntegralName   string     `orm:"size(128);index" json:"integral_name"`
	IntegralType   int8       `orm:"index" json:"integral_type"` // 1:邀请积分; 2:购买积分; 3: 管理奖励
	IntegralSource string     `orm:"size(128);index" json:"integral_source"`
	IntegralAmount float64    `orm:"default(0);digits(22);decimals(8)" json:"integral_amount"`
	OrderNumber    string     `orm:"size(128);index" json:"order_number"`
	SourceUserId   int64      `orm:"index" json:"source_user_id"`
}

func (this *IntegralRecord) TableName() string {
	return common.TableName("integral_record")
}

func (this *IntegralRecord) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *IntegralRecord) SearchField() []string {
	return []string{"user_id"}
}

func (this *IntegralRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *IntegralRecord) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *IntegralRecord) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *IntegralRecord) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *IntegralRecord) InsertDb (db orm.Ormer) error {
	if _, err := db.Insert(this); err != nil {
		return err
	}
	return nil
}

func GetIntegralRecordList(page, pageSize int, user_id int64) ([]*IntegralRecord, int64, error) {
	offset := (page - 1) * pageSize
	ig_trade_list := make([]*IntegralRecord, 0)
	query := orm.NewOrm().QueryTable(IntegralRecord{}).Filter("UserId", user_id)
	total, _ := query.Count()
	_, err := query.OrderBy("-CreatedAt").Limit(pageSize, offset).All(&ig_trade_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return ig_trade_list, total, nil
}


func GetIntegralRecordDetail(id int64) (*IntegralRecord, int, error) {
	var integral IntegralRecord
	if err := orm.NewOrm().QueryTable(IntegralRecord{}).Filter("Id", id).RelatedSel().One(&integral); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &integral, types.ReturnSuccess, nil
}

