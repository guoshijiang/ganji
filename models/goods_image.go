package models

import (
	"errors"
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type GoodsImage struct {
	BaseModel
	Id           int64     `orm:"pk;column(id);auto'" json:"id" form:"id"`
	GoodsId      int64     `json:"goods_id"`
	Image        string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品图片
	IsDispay     int8      `orm:"default(1)" json:"is_dispay"`   // 0 不显示 1 显示
}

func (this *GoodsImage) TableName() string {
	return common.TableName("goods_image")
}


func (this *GoodsImage) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (*GoodsImage) SearchField() []string {
	return []string{"goods_name"}
}

func (this *GoodsImage) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsImage) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsImage) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsImage) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}


func GetGoodsImgList(goods_id int64) ([]*GoodsImage, int, error) {
	var goods_img_list []*GoodsImage
	if _, err := orm.NewOrm().QueryTable(GoodsImage{}).Filter("GoodsId", goods_id).All(&goods_img_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_img_list, types.ReturnSuccess, nil
}

