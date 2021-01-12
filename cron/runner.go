package cron

import (
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	RealTimeExecution  = time.Second * 5
)

func Run() {
	if beego.BConfig.RunMode == "dev" {
		RealTimeExecution = time.Second * 5
	}
	go func() {
		for {
			select {
			case <-time.Tick(RealTimeExecution):
				// 直接邀请奖励发放
				err := IntergralInviteReward()
				if err != nil {
					logrus.Errorf("run miner order invite reward error %v", err)
				} else {
					logrus.Info("run miner order invite reward success.")
				}

				err = LimitTimeGoods()
				if err != nil {
					logrus.Errorf("run miner order invite reward error %v", err)
				} else {
					logrus.Info("run miner order invite reward success.")
				}

				err = ManageReward()
				if err != nil {
					logrus.Errorf("run miner order invite reward error %v", err)
				} else {
					logrus.Info("run miner order invite reward success.")
				}

				err = UserTreeBuild()
				if err != nil {
					logrus.Errorf("run miner order invite reward error %v", err)
				} else {
					logrus.Info("run miner order invite reward success.")
				}
			}
		}
	}()
}

