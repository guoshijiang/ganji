package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type GoodsType struct {
	BaseModel
	Id             int64     `orm:"column(id);auto;size(11)" json:"id"`
	GoodsId        int64     `json:"goods_id"`                         // 商品ID
	TypeDef        int8      `json:"type_def"`                         // 属性类别：0:颜色； 1:大小属性; 2:其他属性
	TypeName       string    `orm:"size(512);index" json:"type_name"`  // 属性文字：商品大小配置等属性文字描述
	IsDispay       int8      `orm:"default(1)" json:"is_dispay"`       // 0 不显示 1 显示
}


func (this *GoodsType) TableName() string {
	return common.TableName("goods_type")
}

func (this *GoodsType) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsType) SearchField() []string {
	return []string{"type_name"}
}

func (this *GoodsType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsType) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsType) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsType) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}


func GetGoodsTypeList(goods_id int64, type_def int8) ([]*GoodsType, int64, error) {
	var type_list []*GoodsType
	if _, err := orm.NewOrm().QueryTable(GoodsType{}).Filter("GoodsId", goods_id).Filter("TypeDef", type_def).All(&type_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return type_list, types.ReturnSuccess, nil
}
