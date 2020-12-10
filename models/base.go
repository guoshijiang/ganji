package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)


func init() {
	dbhost := beego.AppConfig.String("db_host")
	dbport := beego.AppConfig.String("db_port")
	dbuser := beego.AppConfig.String("db_user")
	dbpassword := beego.AppConfig.String("db_pass")
	dbtype := beego.AppConfig.String("db_type")
	dbalias := beego.AppConfig.String("db_alias")
	dbname := beego.AppConfig.String("db_name")
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	if err := orm.RegisterDataBase(dbalias, dbtype, dburl); err != nil {
		panic(errors.Wrap(err, "register data base model"))
	}
	orm.RegisterModel(new(User), new(UserInfo), new(UserWallet),
		new(UserIntegral), new(UserCoupon))
	orm.RunSyncdb(dbalias, false, true)
}


type BaseModel struct {
	IsRemoved      int8       `orm:"default(0);index"`                                       // 0: 正常，1: 删除
	CreatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}


