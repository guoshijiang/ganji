package main

import (
	_ "ganji/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.BConfig.WebConfig.Session.SessionOn = true                // 开启Session模块
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400    // 设置Session有效期,单位秒
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 86400
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

