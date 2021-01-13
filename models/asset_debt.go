package models

import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type AssetDebt struct {
	BaseModel
	Id            int64      `json:"id"`
	RecAmount     float64    `orm:"default(0);digits(22);decimals(8)" json:"rec_amount"`     // 收到的金额
	ActualAmount  float64    `orm:"default(0);digits(22);decimals(8)" json:"actual_amount"`  // 平台实际收入
	UnpaidAmount  float64    `orm:"default(0);digits(22);decimals(8)" json:"unpaid_amount"`  // 应付给商家的钱
	PaidAmount    float64    `orm:"default(0);digits(22);decimals(8)"   json:"paid_amount"`  // 已付给商家的钱
	AssetAmount   float64    `orm:"default(0);digits(22);decimals(8)" json:"asset_amount"`   // 总资产
	DebtAmount    float64    `orm:"default(0);digits(22);decimals(8)" json:"debt_amount"`    // 总负债
	StaticTime    *time.Time `orm:"type(datetime);null" json:"static_time"`                  // 生成时间
}


func (this *AssetDebt) TableName() string {
	return common.TableName("asset_debt")
}

func (this *AssetDebt) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *AssetDebt) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *AssetDebt) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *AssetDebt) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *AssetDebt) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

