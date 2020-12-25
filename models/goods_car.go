package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type GoodsCar struct {
	BaseModel
	Id           int64      `orm:"column(id);auto;size(11)" json:"id" form:"id"`
	GoodsId      int64      `orm:"default(1)" json:"goods_id"`  // 商品ID
	MerchantId   int64      `json:"merchant_id"`                // 商品所属商家ID
	Logo         string     `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"` // 商品LOGO
	AddresId     int64      `orm:"default(1)" json:"addres_id"`  // 地址ID
	GoodsTitle   string     `orm:"size(64)" json:"goods_title"`                              // 商品标题
	GoodsName    string     `orm:"size(512);index" json:"goods_name" form:"goods_name"`      // 产品名称
	UserId       int64      `orm:"size(64);index" json:"user_id"`                            // 购买用户
	BuyNums      int64      `orm:"default(0)" json:"buy_nums"`                               // 购买数量
	PayAmount    float64    `orm:"default(0);digits(22);decimals(8)" json:"pay_amount"`      // 支付金额
}

func (this *GoodsCar) TableName() string {
	return common.TableName("goods_car")
}


func (this *GoodsCar) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GoodsCar) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCar) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsCar) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsCar) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func GetGoodsCarList(page, pageSize int, user_id int64) ([]*GoodsCar, int64, error) {
	offset := (page - 1) * pageSize
	gds_car_list := make([]*GoodsCar, 0)
	query := orm.NewOrm().QueryTable(GoodsCar{}).Filter("UserId", user_id)
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&gds_car_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gds_car_list, total, nil
}

func GetGoodsCarDetail(id int64) (*GoodsCar, int, error) {
	goods_car := GoodsCar{}
	if err := orm.NewOrm().QueryTable(GoodsCar{}).Filter("Id", id).RelatedSel().One(&goods_car); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &goods_car, types.ReturnSuccess, nil
}


func GetGoodsCarDetailByGoodsId(user_id, goods_id int64) (*GoodsCar, int, error) {
	goods_car := GoodsCar{}
	if err := orm.NewOrm().QueryTable(GoodsCar{}).
		Filter("UserId", user_id).
		Filter("GoodsId", goods_id).RelatedSel().One(&goods_car); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &goods_car, types.ReturnSuccess, nil
}


