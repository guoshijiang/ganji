package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Merchant struct {
	BaseModel
	Id             int64     `json:"id" form:"id"`
	Logo           string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"`   // 商家 Logo
	MerchantName   string    `orm:"size(512);index" json:"merchant_name"`   // 商家名称
	MerchantIntro  string    `orm:"size(512);index" json:"merchant_intro"`  // 商家简介
	MerchantDetail string    `orm:"type(text)" json:"merchant_detail"`      // 商家详情
	Address        string    `orm:"size(512);index" json:"address"`         // 店铺地址
	GoodsNum       int64     `json:"goods_num"`                             // 商品总数
	MerchantWay    int8      `orm:"default(0);index" json:"merchant_way"`   // 0:自营商家； 1:认证商家  2:普通商家
	ShopLevel      int8      `json:"shop_level"`                            // 店铺等级
	ShopServer     int8      `json:"shop_server"`                           // 店铺服务
}

func (this *Merchant) TableName() string {
	return common.TableName("merchant")
}


func GetMerchantList(page, pageSize int, merct_name string, address string) ([]*Merchant, int64, error) {
	offset := (page - 1) * pageSize
	merchant_list := make([]*Merchant, 0)
	cond := orm.NewCondition()
	query := orm.NewOrm().QueryTable(Merchant{})
	if merct_name != "" || address != "" {
		cond_merct := cond.And("MerchantName__contains", merct_name).Or("Address__contains", address)
		query =  query.SetCond(cond_merct)
	}
	total, _ := query.Count()
	_, err := query.OrderBy("-GoodsNum").Limit(pageSize, offset).All(&merchant_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return merchant_list, total, nil
}


func GetMerchantDetail(id int64) (*Merchant, int, error) {
	var merchant Merchant
	if err := orm.NewOrm().QueryTable(Merchant{}).Filter("Id", id).RelatedSel().One(&merchant); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &merchant, types.ReturnSuccess, nil
}


func (*Merchant) SearchField() []string {
	return []string{"merchant_name"}
}
