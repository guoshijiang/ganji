package cron

import (
	"ganji/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

/*
	购买者 0.08
	邀请者 0.05
 */
const  (
	BuyerRewardPercent = 0.08
	InviteRewardPercent = 0.05

	InviteRewardStatus = 0
)

// 直邀奖励和购物奖励
func IntergralInviteReward() (err error) {
	db := orm.NewOrm()
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
			err = errors.Wrap(err, "rollback db transaction error in IntergralInviteReward")
		} else {
			err = errors.Wrap(db.Commit(), "commit db transaction error in IntergralInviteReward")
		}
	}()
	order_list, err := models.GetRewardGoodsOrderList(db, InviteRewardStatus)
	for _, order := range order_list {
		my_user, err := models.GetUserByUserId(db, order.UserId)
		if err != nil {
			logs.Error(err)
			return err
		}
		up_user, err := models.GetUserByUserId(db, my_user.InviteMeUserId)
		if err != nil {
			logs.Error(err)
			return err
		}
		my_ig, err := models.GetIgByUserId(db, my_user.Id)
		if err != nil {
			logs.Error(err)
			return err
		}
		my_ig.TotalIg = my_ig.TotalIg + (order.PayAmount * BuyerRewardPercent) * 10
		_, err = db.Update(&my_ig)
		if err != nil {
			logs.Error(err)
			return err
		}
		up_user_ig, err := models.GetIgByUserId(db, up_user.Id)
		if err != nil {
			logs.Error(err)
			return err
		}
		up_user_ig.TotalIg = up_user_ig.TotalIg + (order.PayAmount * InviteRewardPercent) * 10
		_, err = db.Update(&my_ig)
		if err != nil {
			logs.Error(err)
			return err
		}
		order.IsReward = 1
		_, err = db.Update(&order)
		if err != nil {
			logs.Error(err)
			return err
		}
	}
	return nil
}
