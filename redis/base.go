package redis

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
)

var (
	RdsConn *redis.Client
	once	sync.Once
)
func init() {
	once.Do(func() {
		address := beego.AppConfig.String("address")
		password := beego.AppConfig.String("password")
		db_index := beego.AppConfig.String("db_index")
		db_num, _ := strconv.Atoi(db_index)
		RdsConn = redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db_num,
		})
		_, err := RdsConn.Ping().Result()
		if err != nil {
			logs.Info("connect redis fail", err)
			panic(err)
		}
	})
}
