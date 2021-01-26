package cron

import (
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

func LimitTimeGoods() (err error) {
	db := orm.NewOrm()
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
			err = errors.Wrap(err, "rollback db transaction error in LimitTimeGoods")
		} else {
			err = errors.Wrap(db.Commit(), "commit db transaction error in LimitTimeGoods")
		}
	}()
	ltime_gds_list, _ := models.LimitTimeGoodsList(db)
	for _, ltm_gds := range ltime_gds_list{
		if ltm_gds.LeftTime == 0 {
			ltm_gds.Sale = 1
		} else {
			ltm_gds.LeftTime -= 1
		}
		_, err :=db.Update(ltm_gds)
		if err != nil {
			return err
		}
	}
	return nil
}
