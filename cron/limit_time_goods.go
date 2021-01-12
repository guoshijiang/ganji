package cron

import (
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

func LimitTimeGoods() (err error) {
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

	return nil
}
