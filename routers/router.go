package routers

import (
	controllers "ganji/controllers/admin"
	"ganji/controllers/api"
	"ganji/middleware"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dchest/captcha"
	"net/http"
)

func init() {
	//授权登录中间件
	middleware.AuthMiddle()
	beego.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	//admin模块路由
	admin := beego.NewNamespace("/admin",
		//操作日志
		beego.NSRouter("/admin_log/index", &controllers.AdminLogController{}, "get:Index"),
		//登录页
		beego.NSRouter("/auth/login", &controllers.AuthController{}, "get:Login"),
		//退出登录
		beego.NSRouter("/auth/logout", &controllers.AuthController{}, "get:Logout"),
		//二维码图片输出
		beego.NSHandler("/auth/captcha/*.png", captcha.Server(240, 80)),
		//登录认证
		beego.NSRouter("/auth/check_login", &controllers.AuthController{}, "post:CheckLogin"),
		//刷新验证码
		beego.NSRouter("/auth/refresh_captcha", &controllers.AuthController{}, "post:RefreshCaptcha"),

		//首页
		beego.NSRouter("/index/index", &controllers.IndexController{}, "get:Index"),

		beego.NSRouter("/admin_user/index", &controllers.AdminUserController{}, "get:Index"),

		//菜单管理
		beego.NSRouter("/admin_menu/index", &controllers.AdminMenuController{}, "get:Index"),
		//菜单管理-添加菜单-界面
		beego.NSRouter("/admin_menu/add", &controllers.AdminMenuController{}, "get:Add"),
		//菜单管理-添加菜单-创建
		beego.NSRouter("/admin_menu/create", &controllers.AdminMenuController{}, "post:Create"),
		//菜单管理-修改菜单-界面
		beego.NSRouter("/admin_menu/edit", &controllers.AdminMenuController{}, "get:Edit"),
		//菜单管理-更新菜单
		beego.NSRouter("/admin_menu/update", &controllers.AdminMenuController{}, "post:Update"),
		//菜单管理-删除菜单
		beego.NSRouter("/admin_menu/del", &controllers.AdminMenuController{}, "post:Del"),

		//系统管理-个人资料
		beego.NSRouter("/admin_user/profile", &controllers.AdminUserController{}, "get:Profile"),
		//系统管理-个人资料-修改昵称
		beego.NSRouter("/admin_user/update_nickname", &controllers.AdminUserController{}, "post:UpdateNickName"),
		//系统管理-个人资料-修改密码
		beego.NSRouter("/admin_user/update_password", &controllers.AdminUserController{}, "post:UpdatePassword"),
		//系统管理-个人资料-修改头像
		beego.NSRouter("/admin_user/update_avatar", &controllers.AdminUserController{}, "post:UpdateAvatar"),
		//系统管理-用户管理-添加界面
		beego.NSRouter("/admin_user/add", &controllers.AdminUserController{}, "get:Add"),
		//系统管理-用户管理-添加
		beego.NSRouter("/admin_user/create", &controllers.AdminUserController{}, "post:Create"),
		//系统管理-用户管理-修改界面
		beego.NSRouter("/admin_user/edit", &controllers.AdminUserController{}, "get:Edit"),
		//系统管理-用户管理-修改
		beego.NSRouter("/admin_user/update", &controllers.AdminUserController{}, "post:Update"),
		//系统管理-用户管理-启用
		beego.NSRouter("/admin_user/enable", &controllers.AdminUserController{}, "post:Enable"),
		//系统管理-用户管理-禁用
		beego.NSRouter("/admin_user/disable", &controllers.AdminUserController{}, "post:Disable"),
		//系统管理-用户管理-删除
		beego.NSRouter("/admin_user/del", &controllers.AdminUserController{}, "post:Del"),


		//系统管理-轮播图管理
		beego.NSRouter("/sys/banner/index", &controllers.SysController{}, "get:BannerIndex"),
		beego.NSRouter("/sys/banner/add", &controllers.SysController{}, "get:BannerAdd"),
		beego.NSRouter("/sys/banner/create", &controllers.SysController{}, "post:BannerCreate"),
		beego.NSRouter("/sys/banner/edit", &controllers.SysController{}, "get:BannerEdit"),
		beego.NSRouter("/sys/banner/update", &controllers.SysController{}, "post:BannerUpdate"),
		beego.NSRouter("/sys/banner/del", &controllers.SysController{}, "post:BannerDel"),

		//系统管理-版本管理
		beego.NSRouter("/sys/version/index", &controllers.SysController{}, "get:VerIndex"),
		beego.NSRouter("/sys/version/add", &controllers.SysController{}, "get:VerAdd"),
		beego.NSRouter("/sys/version/create", &controllers.SysController{}, "post:VerCreate"),
		beego.NSRouter("/sys/version/edit", &controllers.SysController{}, "get:VerEdit"),
		beego.NSRouter("/sys/version/update", &controllers.SysController{}, "post:VerUpdate"),
		beego.NSRouter("/sys/version/del", &controllers.SysController{}, "post:VerDel"),

		//系统管理-角色管理
		beego.NSRouter("/admin_role/index", &controllers.AdminRoleController{}, "get:Index"),
		//系统管理-角色管理-添加界面
		beego.NSRouter("/admin_role/add", &controllers.AdminRoleController{}, "get:Add"),
		//系统管理-角色管理-添加
		beego.NSRouter("/admin_role/create", &controllers.AdminRoleController{}, "post:Create"),
		//菜单管理-角色管理-修改界面
		beego.NSRouter("/admin_role/edit", &controllers.AdminRoleController{}, "get:Edit"),
		//菜单管理-角色管理-修改
		beego.NSRouter("/admin_role/update", &controllers.AdminRoleController{}, "post:Update"),
		//菜单管理-角色管理-删除
		beego.NSRouter("/admin_role/del", &controllers.AdminRoleController{}, "post:Del"),
		//菜单管理-角色管理-启用角色
		beego.NSRouter("/admin_role/enable", &controllers.AdminRoleController{}, "post:Enable"),
		//菜单管理-角色管理-禁用角色
		beego.NSRouter("/admin_role/disable", &controllers.AdminRoleController{}, "post:Disable"),
		//菜单管理-角色管理-角色授权界面
		beego.NSRouter("/admin_role/access", &controllers.AdminRoleController{}, "get:Access"),
		//菜单管理-角色管理-角色授权
		beego.NSRouter("/admin_role/access_operate", &controllers.AdminRoleController{}, "post:AccessOperate"),

		//商户管理-商户管理
		beego.NSRouter("/merchant/index", &controllers.MerchantController{}, "get:Index"),
		//商户管理-添加界面
		beego.NSRouter("/merchant/add", &controllers.MerchantController{}, "get:Add"),
		//商户管理-添加
		beego.NSRouter("/merchant/create", &controllers.MerchantController{}, "post:Create"),
		//商户管理-修改界面
		beego.NSRouter("/merchant/edit", &controllers.MerchantController{}, "get:Edit"),
		//商户管理-修改
		beego.NSRouter("/merchant/update", &controllers.MerchantController{}, "post:Update"),
		//商户管理-删除
		beego.NSRouter("/merchant/del", &controllers.MerchantController{}, "post:Del"),


		//商品管理-商品管理
		beego.NSRouter("/goods/index", &controllers.GoodsController{}, "get:Index"),
		//商品管理-添加界面
		beego.NSRouter("/goods/add", &controllers.GoodsController{}, "get:Add"),
		//商品管理-添加
		beego.NSRouter("/goods/create", &controllers.GoodsController{}, "post:Create"),
		//商品管理-修改界面
		beego.NSRouter("/goods/edit", &controllers.GoodsController{}, "get:Edit"),
		//商品管理-修改
		beego.NSRouter("/goods/update", &controllers.GoodsController{}, "post:Update"),
		//商品管理-删除
		beego.NSRouter("/goods/del", &controllers.GoodsController{}, "post:Del"),
		//商品管理-商品评价
		beego.NSRouter("/goods/comment", &controllers.GoodsController{}, "get:Comment"),


		//商品分类管理-商品分类管理
		beego.NSRouter("/cat-goods/index", &controllers.GoodsCateController{}, "get:Index"),
		//商品分类管理-添加界面
		beego.NSRouter("/cat-goods/add", &controllers.GoodsCateController{}, "get:Add"),
		//商品分类管理-添加
		beego.NSRouter("/cat-goods/create", &controllers.GoodsCateController{}, "post:Create"),
		//商品分类管理-修改界面
		beego.NSRouter("/cat-goods/edit", &controllers.GoodsCateController{}, "get:Edit"),
		//商品分类管理-修改
		beego.NSRouter("/cat-goods/update", &controllers.GoodsCateController{}, "post:Update"),
		//商品分类管理-删除
		beego.NSRouter("/cat-goods/del", &controllers.GoodsCateController{}, "post:Del"),

		//商品分类管理-商品属性管理
		beego.NSRouter("/goods_type/index", &controllers.GoodsTypeController{}, "get:Index"),
		//商品属性管理-添加界面
		beego.NSRouter("/goods_type/add", &controllers.GoodsTypeController{}, "get:Add"),
		//商品属性管理-添加
		beego.NSRouter("/goods_type/create", &controllers.GoodsTypeController{}, "post:Create"),
		//商品属性管理-修改界面
		beego.NSRouter("/goods_type/edit", &controllers.GoodsTypeController{}, "get:Edit"),
		//商品属性管理-修改
		beego.NSRouter("/goods_type/update", &controllers.GoodsTypeController{}, "post:Update"),
		//商品属性管理-删除
		beego.NSRouter("/goods_type/del", &controllers.GoodsTypeController{}, "post:Del"),

		//Ueditor
		beego.NSRouter("/editor/server", &controllers.EditorController{}, "get,post:Server"),

		//订单管理
		beego.NSRouter("/order/index",&controllers.OrderController{},"get:Index"),
		beego.NSRouter("/order/edit",&controllers.OrderController{},"get:Edit"),
		beego.NSRouter("/order/update",&controllers.OrderController{},"post:Update"),
		beego.NSRouter("/order/del",&controllers.OrderController{},"post:Del"),
		beego.NSRouter("/order/process",&controllers.OrderController{},"get:Process"),
		beego.NSRouter("/order/process/verify",&controllers.OrderController{},"post:Verify"),
		beego.NSRouter("/order/process/detail",&controllers.OrderController{},"get:Detail"),

		//用户管理-用户管理
		beego.NSRouter("/user/index", &controllers.UserController{}, "get:Index"),
		//用户管理-添加界面
		beego.NSRouter("/user/add", &controllers.UserController{}, "get:Add"),
		//用户管理-添加
		beego.NSRouter("/user/create", &controllers.UserController{}, "post:Create"),
		//用户管理-修改界面
		beego.NSRouter("/user/edit", &controllers.UserController{}, "get:Edit"),
		//用户管理-修改
		beego.NSRouter("/user/update", &controllers.UserController{}, "post:Update"),
		//用户管理-删除
		beego.NSRouter("/user/del", &controllers.UserController{}, "post:Del"),
		//用户钱包
		beego.NSRouter("/user/wallet", &controllers.UserController{}, "get:Wallet"),
		//用户积分
		beego.NSRouter("/user/integral", &controllers.UserController{}, "get:Integral"),
		//用户地址
		beego.NSRouter("/user/address", &controllers.UserController{}, "get:Address"),
		beego.NSRouter("/user/coupon", &controllers.UserController{}, "get:Coupon"),
		//积分管理-积分记录
		beego.NSRouter("/integral/index", &controllers.IntegralController{}, "get:Index"),
		beego.NSRouter("/integral/trade", &controllers.IntegralController{}, "get:Trade"),

		//充值记录
		beego.NSRouter("/sys/wallet/record", &controllers.SysController{}, "get:WalletRecord"),
		beego.NSRouter("/sys/customer/service", &controllers.SysController{}, "get:CustomerService"),
		beego.NSRouter("/sys/customer/question", &controllers.SysController{}, "get:CustomerQuestion"),
		beego.NSRouter("/sys/question/add", &controllers.SysController{}, "get:CustomerQuestionAdd"),
		beego.NSRouter("/sys/question/update", &controllers.SysController{}, "post:CustomerQuestionUpdate"),
		beego.NSRouter("/sys/question/create", &controllers.SysController{}, "post:CustomerQuestionCreate"),
		beego.NSRouter("/sys/question/del", &controllers.SysController{}, "post:CustomerQuestionDel"),
		beego.NSRouter("/sys/question/edit", &controllers.SysController{}, "get:CustomerQuestionEdit"),
		beego.NSRouter("/sys/customer/add", &controllers.SysController{}, "get:CustomerServiceAdd"),
		beego.NSRouter("/sys/customer/edit", &controllers.SysController{}, "get:CustomerServiceEdit"),
		beego.NSRouter("/sys/customer/update", &controllers.SysController{}, "post:CustomerServiceUpdate"),
		beego.NSRouter("/sys/customer/create", &controllers.SysController{}, "post:CustomerServiceCreate"),
		beego.NSRouter("/sys/customer/del", &controllers.SysController{}, "post:CustomerServiceDel"),
		beego.NSRouter("/sys/record/verify", &controllers.SysController{}, "post:WalletRecordVerify"),

		//结算账户管理-账户管理
		beego.NSRouter("/settle/index", &controllers.SettleAccountController{}, "get:Index"),
		//结算账户管理-添加界面
		beego.NSRouter("/settle/add", &controllers.SettleAccountController{}, "get:Add"),
		//结算账户管理-添加
		beego.NSRouter("/settle/create", &controllers.SettleAccountController{}, "post:Create"),
		//结算账户管理-修改界面
		beego.NSRouter("/settle/edit", &controllers.SettleAccountController{}, "get:Edit"),
		//结算账户管理-修改
		beego.NSRouter("/settle/update", &controllers.SettleAccountController{}, "post:Update"),
		//结算账户管理-删除
		beego.NSRouter("/settle/del", &controllers.SettleAccountController{}, "post:Delete"),

		//日结算列表
		beego.NSRouter("/settle/daily", &controllers.SettleController{}, "get:Daily"),
		//商家结算
		beego.NSRouter("/settle/bill", &controllers.SettleController{}, "get:BillSettle"),
		//商家结算
		beego.NSRouter("/settle/search_settle", &controllers.SettleController{}, "get:SearchSettle"),
		//商家结算-确认界面
		beego.NSRouter("/settle/configure", &controllers.SettleController{}, "get:Configure"),
		//商家结算-创建
		beego.NSRouter("/bill/create", &controllers.SettleController{}, "post:BillSettleCreate"),
		//商家结算-平台确认结算
		beego.NSRouter("/bill/confirm", &controllers.SettleController{}, "post:BillSettleUpdate"),

	)
	beego.AddNamespace(admin)
	// API 部分
	api_path := beego.NewNamespace("/v1",
		beego.NSNamespace("/image",
			beego.NSInclude(
				&api.ImageController{},
			),
		),

		beego.NSNamespace("/address",
			beego.NSInclude(
				&api.UserAddressController{},
			),
		),

		beego.NSNamespace("/category",
			beego.NSInclude(
				&api.CategoryController{},
			),
		),

		beego.NSNamespace("/comment",
			beego.NSInclude(
				&api.CommentController{},
			),
		),

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&api.GoodsController{},
			),
		),

		beego.NSNamespace("/index",
			beego.NSInclude(
				&api.IndexController{},
			),
		),

		beego.NSNamespace("/merchant",
			beego.NSInclude(
				&api.MerchantController{},
			),
		),

		beego.NSNamespace("/order",
			beego.NSInclude(
				&api.OrderController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&api.UserController{},
			),
		),

		beego.NSNamespace("/user_info",
			beego.NSInclude(
				&api.UserInfoController{},
			),
		),

		beego.NSNamespace("/wallet_integral",
			beego.NSInclude(
				&api.WalletIntegralController{},
			),
		),

		beego.NSNamespace("/goods_car",
			beego.NSInclude(
				&api.GoodsCarController{},
			),
		),

		beego.NSNamespace("/group_order",
			beego.NSInclude(
				&api.GroupOrderController{},
			),
		),

		beego.NSNamespace("/market",
			beego.NSInclude(
				&api.MarketController{},
			),
		),

		beego.NSNamespace("/user_account",
			beego.NSInclude(
				&api.UserAccountController{},
			),
		),

		beego.NSNamespace("/w_or_d",
			beego.NSInclude(
				&api.DepositWithdrawController{},
			),
		),

		beego.NSNamespace("/pay",
			beego.NSInclude(
				&api.PayController{},
			),
		),

		beego.NSNamespace("/notify",
			beego.NSInclude(
				&api.NotifyController{},
			),
		),

		beego.NSNamespace("/version",
			beego.NSInclude(
				&api.VersionController{},
			),
		),
	)
	beego.AddNamespace(api_path)
}
