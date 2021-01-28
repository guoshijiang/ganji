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

				err = UserTreeBuildManageReward()
				if err != nil {
					logrus.Errorf("run miner order invite reward error %v", err)
				} else {
					logrus.Info("run miner order invite reward success.")
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case now := <-time.Tick(time.Minute):
				if now.Hour() == 0 && now.Minute() == 1 {
					yesterday := now.AddDate(0, 0, -1).Format("20060102")
					err := MerchantSettleDaiy(yesterday)
					if err != nil {
						logrus.Errorf("run user patch usdt order income error %v", err)
					} else {
						logrus.Info("run user patch usdt order income success.")
					}
				}
			}
		}
	}()

}

