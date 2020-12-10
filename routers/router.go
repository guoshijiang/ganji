package routers

import (
	"ganji/controllers/admin"
	"ganji/controllers/api"
	"github.com/astaxie/beego"
)

func init() {
	// 后台部分
    beego.Router("/", &admin.MainController{})


	// API 部分
	api_path := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&api.UserController{},
			),
		),
	)
	beego.AddNamespace(api_path)
}
