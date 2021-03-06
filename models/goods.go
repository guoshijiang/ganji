package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type Goods struct {
	BaseModel
	Id             int64     `orm:"column(id);auto;size(11)" json:"id" form:"id"`
	GoodsCatId     int64     `json:"goods_cat_id"`                        // 商品所属一级分类ID
	GoodsLastCatId int64     `json:"goods_level_cat_id"`                  // 商品所属最后一级分类ID
	GoodsMark      string    `orm:"size(512);index" json:"goods_mark"`    // 商品备注
	Serveice       string    `orm:"size(512);index" json:"serveice"`      // 服务说明
	CalcWay        int8      `orm:"default(0);index" json:"calc_way"`     // 0:按件计量 1:按近计量
	MerchantId     int64     `json:"merchant_id"`                         // 商品所属商家ID
	Title          string    `orm:"size(512);index" json:"title"`         // 商品标题
	Logo           string    `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"logo" form:"logo"`   // 商品封面
	TotalAmount    int64     `orm:"default(150000)" json:"total_amount" form:"total_amount"`                                    // 商品总量
	LeftAmount     int64     `orm:"default(150000)" json:"left_amount" form:"left_amount"`                                      // 剩余商品总量
	GoodsPrice     float64   `orm:"default(1);digits(22);decimals(8)" json:"goods_price"`                                       // 商品价格
	GoodsDisPrice  float64   `orm:"default(1);digits(22);decimals(8)" json:"goods_discount_price"`                              // 商品折扣价格
	GoodsIntegral  float64   `orm:"default(1);digits(22);decimals(8)" json:"goods_integral"`                                    // 购买需要的积分数量
	SendIntegral   float64   `orm:"default(1);digits(22);decimals(8)" json:"send_integral"` // 购买商品赠送积分
	GoodsName      string    `orm:"size(512);index" json:"goods_name" form:"goods_name"`    // 产品名称
	GoodsParams    string    `orm:"type(text)" json:"goods_params" form:"goods_params"`     // 产品参数
	GoodsDetail    string    `orm:"type(text)" json:"goods_detail" form:"goods_detail"`     // 产品详细介绍
	Discount       float64   `orm:"default(0);index" json:"discount"`                       // 折扣，取值 0.1-9.9；0代表不打折
	Sale           int8      `orm:"default(0);index" json:"sale" form:"sale"`               // 0:上架 1:下架
	IsDisplay      int8      `orm:"default(0);index" json:"is_display" form:"is_display"`   // 0:首页不展示, 1:首页展示
	SellNums       int64     `orm:"default(0);index" json:"sell_nums"`                      // 售出数量
	IsHot          int8      `orm:"default(0);index" json:"is_hot"`                         // 0:非爆款产品 1:爆款产品
	IsDiscount     int8      `orm:"default(0);index" json:"is_discount"`                    // 0:不打折，1:打折活动产品
	IsIgSend       int8      `orm:"default(0);index" json:"is_ig_send"`                     // 0:正常， 1:赠送积分
	IsGroup        int8      `orm:"default(0);index" json:"is_group"`                       // 0:非拼购产品 1:拼购产品
	GroupNumber    int64     `orm:"default(100);index" json:"group_number"`                 // 助力人数
	IsIntegral     int8      `orm:"default(0);index" json:"is_integral"`                    // 0:非积分兑换产品 1:积分兑换产品
	LeftTime       int64     `orm:"default(0);index" json:"left_time"`                      // 限时产品剩余时间
	IsLimitTime    int8      `orm:"default(0);index" json:"is_limit_time"`                  // 0:不是限时产品 1:是限时
}

type Select struct {
	Id			int					`json:"id"`
	Name		string				`json:"name"`
}

func (this *Goods) TableName() string {
	return common.TableName("goods")
}

func (this *Goods) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (*Goods) SearchField() []string {
	return []string{"goods_name"}
}

func (this *Goods) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Goods) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Goods) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Goods) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}


func GetLimitTimeGoodsList() ([]*Goods, int, error) {
	var goods_list []*Goods
	if _, err := orm.NewOrm().QueryTable(Goods{}).Filter("IsLimitTime", 1).OrderBy("-SellNums").Limit(6).All(&goods_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_list, types.ReturnSuccess, nil
}


func GetHotGoodsList() ([]*Goods, int, error) {
	var goods_list []*Goods
	if _, err := orm.NewOrm().QueryTable(Goods{}).Filter("IsHot", 1).OrderBy("-SellNums").Limit(16).All(&goods_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_list, types.ReturnSuccess, nil
}


func GetDiscountGoodsList() ([]*Goods, int, error) {
	var goods_list []*Goods
	if _, err := orm.NewOrm().QueryTable(Goods{}).Filter("IsDiscount", 1).OrderBy("-SellNums").Limit(12).All(&goods_list); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return goods_list, types.ReturnSuccess, nil
}


// 0:爆款产品; 1:活动优选；2:积分兑换；3:拼团送
func GetIndexDownGoodsList(page int, pageSize int, query_way int8) ([]*Goods, int64, error) {
	offset := (page - 1) * pageSize
	goods_list := make([]*Goods, 0)
	if query_way == 0 {  // 爆款产品
		query_zero := orm.NewOrm().QueryTable(Goods{}).Filter("IsHot", 1).OrderBy("-SellNums")
		total, _ := query_zero.Count()
		_, err := query_zero.Limit(pageSize, offset).All(&goods_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
		return goods_list, total, nil
	} else if query_way == 1 { // 活动优选
		query_one := orm.NewOrm().QueryTable(Goods{}).Filter("IsDiscount", 1).OrderBy("-SellNums")
		total, _ := query_one.Count()
		_, err := query_one.Limit(pageSize, offset).All(&goods_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
		return goods_list, total, nil
	} else if query_way == 2 { // 积分兑换
		query_two := orm.NewOrm().QueryTable(Goods{}).Filter("IsIntegral", 1)
		total, _ := query_two.Count()
		_, err := query_two.Limit(pageSize, offset).All(&goods_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
		return goods_list, total, nil
	}else if query_way == 3 { // 拼团送
		query_three := orm.NewOrm().QueryTable(Goods{}).Filter("IsGroup", 1)
		total, _ := query_three.Count()
		_, err := query_three.Limit(pageSize, offset).All(&goods_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
		return goods_list, total, nil
	} else {
		return nil, types.InvalidVerifyWay, errors.New("没有这种查询方式")
	}
}


func GetCategoryGoodsList(page, pageSize int, first_level_id, last_level_id int64) ([]*Goods, int64, error) {
	offset := (page - 1) * pageSize
	goods_list := make([]*Goods, 0)
	if first_level_id <= 0 {
		query_dis := orm.NewOrm().QueryTable(Goods{}).Filter("IsDiscount", 1).Filter("IsGroup", 0).Filter("IsIntegral", 0).OrderBy("-SellNums")
		total, _ := query_dis.Count()
		_, err := query_dis.Limit(pageSize, offset).All(&goods_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
		return goods_list, total, nil
	} else {
		query := orm.NewOrm().QueryTable(Goods{}).Filter("GoodsCatId", first_level_id).Filter("IsGroup", 0).Filter("IsIntegral", 0).OrderBy("-SellNums")
		if last_level_id > 0 {
			query.Filter("GoodsLastCatId", last_level_id).OrderBy("-SellNums")
		}
		total, _ := query.Count()
		_, err := query.Limit(pageSize, offset).All(&goods_list)
		if err != nil {
			return nil, 0, errors.New("查询数据库失败")
		}
		return goods_list, total, nil
	}
}


func GetMerchantGoodsNums(metchant_id int64) int64 {
	total, err := orm.NewOrm().QueryTable(Goods{}).Filter("MerchantId", metchant_id).Count()
	if err != nil {
		return 0
	}
	return total
}


// 0:全部；1:活动优选；2:爆款产品
func GetMerchantGoodsList(page, pageSize int, metchant_id int64, query_way int8) ([]*Goods, int64, error) {
	offset := (page - 1) * pageSize
	goods_list := make([]*Goods, 0)
	query := orm.NewOrm().QueryTable(Goods{}).Filter("MerchantId", metchant_id)
	if query_way == 0 {
		query = query
	} else if query_way == 1 {
		query = query.Filter("IsDiscount", 1)
	} else if query_way == 2 {
		query = query.Filter("IsHot", 1).OrderBy("-SellNums")
	} else {
		return nil, types.InvalidVerifyWay, errors.New("没有这种查询方式")
	}
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&goods_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return goods_list, total, nil
}


func GetLtGoodsList(page, pageSize int) ([]*Goods, int64, error) {
	offset := (page - 1) * pageSize
	goods_list := make([]*Goods, 0)
	query := orm.NewOrm().QueryTable(Goods{}).Filter("IsLimitTime", 1)
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&goods_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return goods_list, total, nil
}


func GetOrderDownHotGoodsList(page, pageSize int) ([]*Goods, int64, error) {
	offset := (page - 1) * pageSize
	goods_list := make([]*Goods, 0)
	query := orm.NewOrm().QueryTable(Goods{}).Filter("IsHot", 1)
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&goods_list)
	if err != nil {
		return nil, 0, errors.New("查询数据库失败")
	}
	return goods_list, total, nil
}

func GetGoodsDetail(id int64) (*Goods, int, error) {
	var goods Goods
	if err := orm.NewOrm().QueryTable(Goods{}).Filter("Id", id).RelatedSel().One(&goods); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &goods, types.ReturnSuccess, nil
}

func LimitTimeGoodsList(db orm.Ormer) ([]*Goods, error) {
	var gds_list []*Goods
	if _, err := db.QueryTable(Goods{}).Filter("IsLimitTime", 1).All(&gds_list); err != nil {
		return nil, errors.New("数据库操作错误")
	}
	return gds_list, nil
}


