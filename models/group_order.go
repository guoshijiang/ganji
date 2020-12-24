package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type GroupOrder struct {
	BaseModel
	Id                int64      `json:"id"`
	GoodsId           int64      `orm:"size(64);index" json:"goods_id"`                         // 商品 ID
	MerchantId        int64      `orm:"size(64);index" json:"merchant_id"`                        // 商户 ID
	AddressId         int64      `orm:"size(64);index" json:"address_id"`                       // 地址 ID
	GoodsTitle        string     `orm:"size(64)" json:"goods_title"`                            // 商品标题
	Logo              string     `orm:"size(150);default(/static/upload/default/user-default-60x60.png)" json:"image"` // 商品Logo
	UserId            int64      `orm:"size(64);index" json:"user_id"`                          // 购买用户
	BuyNums           int64      `orm:"default(0)" json:"buy_nums"`                             // 购买数量
	OrderAmount       float64    `orm:"default(0);digits(22);decimals(8)" json:"order_amount"`  // 订单金额
	GroupNumber       int64      `orm:"default(100);index" json:"group_number"`                 // 助力人数
	HelpNumber        int64      `orm:"default(100);index" json:"help_number"`                  // 已经助力人数
	OrderNumber       string     `orm:"size(64);index" json:"order_number"`                     // 订单号
	ShipNumber        string     `orm:"size(64);index;default('')" json:"ship_number"`          // 运单号
	Logistics	      string     `orm:"size(64);index;default('')" json:"logistics"`            // 物流公司
	OrderStatus       int8       `orm:"index" json:"order_status"`                              // 0: 助力中，1: 助力成功，2:已发货 3：已收货
	FailureReason     string     `json:"failure_reason"`                                        // 失败原因
	PayAt             *time.Time `orm:"type(datetime);null" json:"pay_at"`                      // 助力成功时间
	DealMerchant      string     `orm:"default('')" json:"deal_user"`                           // 处理商家
	DealAt            time.Time  `orm:"null;type(datetime);" json:"deal_at"`                    // 处理时间
	IsValid           int8       `orm:"default(0)" json:"is_valid"`                             // 0:有效 1:无效
	DeadLime          string     `orm:"size(64);index;default('')" json:"dead_lime"`            // 截止助力时间
}

func (this *GroupOrder) TableName() string {
	return common.TableName("group_order")
}

func (this *GroupOrder) Read(fields ...string) error {
	logs.Info(fields)
	return nil
}

func (this *GroupOrder) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GroupOrder) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GroupOrder) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GroupOrder) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func ExistOrderByUser(goods_id, user_id int64) bool {
	return orm.NewOrm().QueryTable(GroupOrder{}).Filter("GoodsId", goods_id).Filter("UserId", user_id).Exist()
}

// 1:我的助力订单 2 好友的助力订单
func GroupOrderList(page, pageSize int, invite_uid, user_id int64, query_way int8) ([]*GroupOrder, int64, error) {
	offset := (page - 1) * pageSize
	group_order_list := make([]*GroupOrder, 0)
	if query_way == 1 {  // 我的助力订单
		query := orm.NewOrm().QueryTable(GroupOrder{}).Filter("UserId", user_id)
		total, _ := query.Count()
		_, err := query.Limit(pageSize, offset).All(&group_order_list)
		if err != nil {
			return nil, types.SystemDbErr, errors.New("查询数据库失败")
		}
		return group_order_list, total, nil
	} else if query_way == 2 {  // 好友的助力订单
		user := User{}
		orm.NewOrm().QueryTable(User{}).Filter("UserId", invite_uid).One(&user)
		query := orm.NewOrm().QueryTable(GroupOrder{}).Filter("UserId", user.Id)
		total, _ := query.Count()
		_, err := query.Limit(pageSize, offset).All(&group_order_list)
		if err != nil {
			return nil, types.SystemDbErr, errors.New("查询数据库失败")
		}
		return group_order_list, total, nil
	} else {
		return nil, types.InvalidVerifyWay, errors.New("无效的查询方式")
	}
}


func GetGroupOrderDetail(id int64) (*GroupOrder, int, error) {
	var order_dtl GroupOrder
	if err := orm.NewOrm().QueryTable(GroupOrder{}).Filter("Id", id).RelatedSel().One(&order_dtl); err != nil {
		return nil, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
	}
	return &order_dtl, types.ReturnSuccess, nil
}


func HelpOrder(order_id, buy_user_id, slef_user_id int64) (bool, int, error){
	exst := orm.NewOrm().QueryTable(GroupHelper{}).Filter("GroupOrderId", order_id).Filter("HelperUserId", slef_user_id).Exist()
	if exst == false {
		g_h := GroupHelper{
			GroupOrderId: order_id,
			BuyUserId: buy_user_id,
			HelperUserId: slef_user_id,
		}
		err := g_h.Insert()
		if err != nil {
			return false, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
		}
		order_dtl, _, _ := GetGroupOrderDetail(order_id)
		order_dtl.HelpNumber = order_dtl.HelpNumber + 1
		err = order_dtl.Update()
		if err != nil {
			return false, types.SystemDbErr, errors.New("数据库查询失败，请联系客服处理")
		}
	} else {
		return false, types.AlreadyHelp, errors.New("您已经助力过了")
	}
	return true, types.ReturnSuccess, nil
}