package cron

import (
	"ganji/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)


const (
	V1LevelSize   = 500
	V2LevelSize   = 30000
	V3LevelSize   = 100000
	V4LevelSize   = 350000
	V5LevelSize   = 2000000

	V1Level       = 1
	V2Level       = 2
	V3Level       = 3
	V4Level       = 4
	V5Level       = 5

	V1LevelReward = 0.4
	V2LevelReward = 0.5
	V3LevelReward = 0.6
	V4LevelReward = 0.7
	V5LevelReward = 0.8

)

func UserTreeBuildManageReward() (err error) {
	db := orm.NewOrm()
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
			err = errors.Wrap(err, "rollback db transaction error in ManageReward")
		} else {
			err = errors.Wrap(db.Commit(), "commit db transaction error in ManageReward")
		}
	}()
	unbuild_order_list, err := models.GetUnbuildTreeGoodsOrders(db)
	if err != nil {
		logs.Info(err)
		return err
	}
	for _, order := range unbuild_order_list {
		goods := models.Goods{}
		err := db.QueryTable(models.Goods{}).Filter("Id", order.GoodsId).RelatedSel().One(&goods)
		if err != nil {
			logs.Error(err)
			return err
		}
		buy_user := models.User{}
		err = db.QueryTable(models.User{}).Filter("id", order.UserId).RelatedSel().One(&buy_user)
		if err != nil {
			logs.Error(err)
			return err
		}
		var goods_price float64
		if goods.IsDiscount == 1 {
			goods_price = goods.GoodsDisPrice
		} else {
			goods_price = goods.GoodsPrice
		}
		pay_total_amount := float64(order.BuyNums) * goods_price
		err = UserTreeCalc(&buy_user, &buy_user, pay_total_amount, db)
		if err != nil {
			logs.Info(err)
			return err
		}
		err = SendLsdtReward(&buy_user, &buy_user, order.SendIntegral, order.OrderNumber, true, db)
		if err != nil {
			err = db.Rollback()
			logs.Info(err)
			return err
		}
		order.IsBuild = 1
		_, err = db.Update(order)
		if err != nil {
			logs.Info(err)
			return err
		}
	}
	return nil
}

func UserTreeCalc(my *models.User, bottom_user *models.User, order_amount float64, db orm.Ormer) error {
	crfr_my_tree, _ := models.GetUserTreeByid(db, my.Id)
	var my_self_buy_price, my_descendant_buy_price float64
	if crfr_my_tree.UserId == bottom_user.Id {
		if crfr_my_tree.IsValid == false {
			crfr_my_tree.IsValid = true
		}
		my_self_buy_price = crfr_my_tree.SelfBuyPrice + order_amount
		crfr_my_tree.SelfBuyPrice = my_self_buy_price
	}
	my_descendant_buy_price = crfr_my_tree.DescendantBuyPrice + order_amount
	crfr_my_tree.DescendantBuyPrice = my_descendant_buy_price
	_, err := db.Update(crfr_my_tree)
	if err != nil {
		logs.Error(err)
	}
	err = SetUserLevel(my_self_buy_price, my_descendant_buy_price, my, crfr_my_tree, db)
	if err != nil {
		logs.Error(err)
	}
	crfr_my_father, _ := my.GetFatherNode(db)
	if crfr_my_father != nil {
		err = UserTreeCalc(crfr_my_father, bottom_user, order_amount, db)
		if err != nil {
			logs.Error(err)
		}
	}
	return nil
}

func SetUserLevel(self_buy_price float64, descendant_buy_price float64, user *models.User, crfr_tree *models.CrfrUserTree, db orm.Ormer) error {
	var user_level int8
	if (self_buy_price >= V1LevelSize && self_buy_price < V2LevelSize) ||
		(descendant_buy_price >= V1LevelSize && descendant_buy_price < V2LevelSize) {
		if user.MemberLevel >= V1Level {
			user_level = user.MemberLevel
		} else {
			user_level = V1Level
		}
	} else if (self_buy_price >= V2Level && self_buy_price < V3LevelSize) ||
		(descendant_buy_price >= V2LevelSize && descendant_buy_price <= V3LevelSize) {
		if user.MemberLevel >= V2Level {
			user_level = user.MemberLevel
		} else {
			user_level = V2Level
		}
	} else if (self_buy_price >= V3LevelSize && self_buy_price < V4LevelSize) ||
		(descendant_buy_price >= V3LevelSize && descendant_buy_price < V4LevelSize) {
		if user.MemberLevel >= V3Level {
			user_level = user.MemberLevel
		} else {
			user_level = V3Level
		}
	} else if (self_buy_price >= V4LevelSize && self_buy_price < V5LevelSize) ||
		(descendant_buy_price >= V2LevelSize && descendant_buy_price < V5LevelSize) {
		if user.MemberLevel >= V4Level {
			user_level = user.MemberLevel
		} else {
			user_level = V4Level
		}
	} else if self_buy_price >= V5LevelSize ||
		descendant_buy_price >= V5LevelSize {
		if user.MemberLevel >= V5Level {
			user_level = user.MemberLevel
		} else {
			user_level = V5Level
		}
	} else {
		return errors.New("invalid user level set")
	}
	user.MemberLevel = user_level
	crfr_tree.UserLevel = user_level
	_, err := db.Update(user)
	if err != nil {
		err = db.Rollback()
		logs.Error(err)
		return err
	}
	_, err = db.Update(crfr_tree)
	if err != nil {
		err = db.Rollback()
		logs.Error(err)
		return err
	}
	return nil
}

func SendLsdtReward(my *models.User, bottom_user *models.User, order_pay_amount float64, order_number string, is_me bool, db orm.Ormer) error {
	father, _ := my.GetFatherNode(db)
	if father != nil && father.MemberLevel <= V5Level {
		crfr_father_tree, _ := models.GetUserTreeByid(db, father.Id)
		if crfr_father_tree.IsValid == true {
			var reward_per float64
			if is_me == true {
				reward_per = GeRewardtPercent(father.MemberLevel)
			} else {
				reward_per = GeRewardtPercent(father.MemberLevel) - GeRewardtPercent(my.MemberLevel)
			}
			if reward_per > 0 {
				reward_amount := order_pay_amount * reward_per
				reward_lsdt := models.IntegralRecord{
					UserId:         father.Id,
					IntegralName:   "积分",
					IntegralType:   3,
					IntegralSource: "积分管理奖励",
					SourceUserId:   bottom_user.Id,
					OrderNumber:    order_number,
					IntegralAmount: reward_amount,
				}
				err := reward_lsdt.InsertDb(db)
				if err != nil {
					logs.Error(err)
				}
				user_integral := models.UserIntegral{}
				err = db.QueryTable(models.UserIntegral{}).Filter("UserId", father.Id).One(&user_integral)
				if err != nil {
					logs.Error(err)
				}
				user_integral.TotalIg += user_integral.TotalIg + reward_amount
				err = user_integral.UpdateDb(db)
				if err != nil {
					logs.Error(err)
				}
				err = SendLsdtReward(father, bottom_user, order_pay_amount, order_number, false, db)
				if err != nil {
					logs.Error(err)
				}
			} else {
				err := SendLsdtReward(father, bottom_user, order_pay_amount, order_number, false, db)
				if err != nil {
					logs.Error(err)
				}
			}
		} else {
			err := SendLsdtReward(father, bottom_user, order_pay_amount, order_number, false, db)
			if err != nil {
				logs.Error(err)
			}
		}
	}
	return nil
}

func GeRewardtPercent(level int8) float64 {
	if level == V1Level {
		return V1LevelReward
	} else if level == V2Level {
		return V2LevelReward
	} else if level == V3Level {
		return V3LevelReward
	} else if level == V4Level {
		return V4LevelReward
	} else if level == V5Level {
		return V5LevelReward
	} else {
		logs.Error("invalid user level")
	}
	return 0
}
