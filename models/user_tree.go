package models

import (
	"ganji/common"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type CrfrUserTree struct {
	Id                    int64   `json:"id"`
	UserId                int64   `orm:"index;" json:"user_id"`
	FatherUserId          int64   `orm:"default(0);index" json:"father_user_id"`    // 邀请人UserID 注册用户时更新
	IsValid               bool    `orm:"default(false)" json:"is_valid"`            // 是否为有效用户，曾经购买过订单即有效用户
	SelfBuyPrice          float64 `orm:"default(0)" json:"self_buy_price"`          // 自购数量
	DescendantBuyPrice    float64 `orm:"default(0)" json:"descendant_buy_price"`    // 网体数量
	UserLevel             int8    `orm:"default(0)" json:"user_level"`
}

func (cut *CrfrUserTree) TableName() string {
	return common.TableName("user_tree")
}

func (cut *CrfrUserTree) TableUnique() [][]string {
	return [][]string{
		{"UserId", "FatherUserId"},
	}
}

func (cut *CrfrUserTree) Insert() error {
	if _, err := orm.NewOrm().Insert(cut); err != nil {
		return err
	}
	return nil
}

func GetUserTreeByid(db orm.Ormer, user_id int64) (*CrfrUserTree, error) {
	crfr_tree := CrfrUserTree{}
	if err := db.QueryTable(CrfrUserTree{}).Filter("UserId", user_id).RelatedSel().One(&crfr_tree); err != nil {
		return nil,  errors.New("数据库查询失败，请联系客服处理")
	}
	return &crfr_tree, nil
}