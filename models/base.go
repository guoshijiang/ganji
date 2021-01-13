package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

func init() {
	mysqlConfig, _ := beego.AppConfig.GetSection("mysql")
	dburl := mysqlConfig["db_user"] + ":" + mysqlConfig["db_pass"] + "@tcp(" + mysqlConfig["db_host"] + ":" + mysqlConfig["db_port"] + ")/" + mysqlConfig["db_name"] + "?charset=utf8&loc=Asia%2FShanghai"
	if err := orm.RegisterDataBase(mysqlConfig["db_alias"], mysqlConfig["db_type"], dburl); err != nil {
		panic(errors.Wrap(err, "register data base model"))
	}
	orm.RegisterModel(new(User), new(UserInfo), new(UserWallet), new(UserIntegral), new(UserCoupon),
		new(AdminUser), new(AdminMenu), new(AdminRole), new(Goods), new(GoodsCar), new(Merchant),
		new(GoodsComment), new(GoodsCat), new(GoodsImage), new(GoodsOrder), new(OrderProcess), new(GroupOrder),
		new(GroupHelper), new(ImageFile),  new(IntegralRecord), new(IntegralTrade), new(UserAddress),
		new(Version), new(WalletRecord), new(Banner), new(CustomerService), new(Questions), new(UserAccount),
		new(AssetDebt), new(MerchantSettle), new(MerchantWallet), new(MerchantWithdraw))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	err := orm.RunSyncdb(mysqlConfig["db_alias"], false, true)
	if err != nil {
		logs.Error(err.Error())
	}
	////admin asd..123 aaa/bbb 123456
	//insertAdmin()
	//insertRole()
	//insertMenu()
}

type BaseModel struct {
	IsRemoved      int8       `orm:"default(0);index"`                                       // 0: 正常，1: 删除
	CreatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt      time.Time  `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}

func insertRole() {
	var err error
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_role` VALUES (1, '管理员', '后台管理员角色', '1,19,20,21,22,23,24,25,53,54,76', 1);").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_role` VALUES (2, '商户A', '商户A', '1,2,18,19,20', 1);").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_role` VALUES (6, '商户B', '商户B', '1,2,18,19,20', 1);").Exec()
	fmt.Println("err---", err)
}

func insertAdmin(){
	var err error
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_user` VALUES (1, 'admin', 'JDJhJDEwJFdRaU5qRlpLUmZ1dG8uUXdpaXNaaS40SkIwdXNhQmRZOTZsMmc5by53SldMUi9qTjVLc1dp', '超级管理员', '/static/uploads/attachment/aecb9fb7-871b-43fc-9414-a4265d0cb72d.png', '1', 1, 0,0);").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_user` VALUES (2, 'aaa', 'JDJhJDEwJEhHaWZ0LkdzaTRtYzRRMWNvNncxTC5HL0NEZnk5bkpJdmw1bzdiRDE2OEVSOXROamk2MWxX', 'aaa', '/static/admin/images/avatar.png', '2', 1, 0,0);").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_user` VALUES (3, 'bbb', 'JDJhJDEwJFpKVElLZVpBLjV5YXRObC5FUDdMVy5sQ1F4ekx0VjVzd3laQ0p1L05ERU1kZDlvNTFJcnhh', 'bbb', '/static/admin/images/avatar.png', '6', 1, 0,0);").Exec()
	fmt.Println("err---", err)
}

func insertMenu(){
	var err error
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (1, 0, '后台首页', 'admin/index/index', 'fa-home', 1, 99, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (2, 0, '系统管理', 'admin/sys', 'fa-desktop', 1, 1099, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (3, 2, '用户管理', 'admin/admin_user/index', 'fa-user', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (4, 3, '添加用户界面', 'admin/admin_user/add', 'fa-plus', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (5, 3, '修改用户界面', 'admin/admin_user/edit', 'fa-edit', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (6, 3, '删除用户', 'admin/admin_user/del', 'fa-close', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (7, 2, '角色管理', 'admin/admin_role/index', 'fa-group', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (8, 7, '添加角色界面', 'admin/admin_role/add', 'fa-plus', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (9, 7, '修改角色界面', 'admin/admin_role/edit', 'fa-edit', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (10, 7, '删除角色', 'admin/admin_role/del', 'fa-close', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (11, 7, '角色授权界面', 'admin/admin_role/access', 'fa-key', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (12, 2, '菜单管理', 'admin/admin_menu/index', 'fa-align-justify', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (13, 12, '添加菜单界面', 'admin/admin_menu/add', 'fa-plus', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (14, 12, '修改菜单界面', 'admin/admin_menu/edit', 'fa-edit', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (15, 12, '删除菜单', 'admin/admin_menu/del', 'fa-close', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (16, 2, '操作日志', 'admin/admin_log/index', 'fa-keyboard-o', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (17, 16, '日志详情', 'admin/admin_log/view', 'fa-search-plus', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (18, 2, '个人资料', 'admin/admin_user/profile', 'fa-smile-o', 1, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (19, 0, '订单管理', 'admin/order/mange', 'fa-first-order', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (20, 19, '订单管理', 'admin/order/index', 'fa-first-order', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (55, 3, '修改头像', 'admin/admin_user/update_avatar', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (56, 3, '添加用户', 'admin/admin_user/create', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (57, 3, '修改用户', 'admin/admin_user/update', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (58, 3, '用户启用', 'admin/admin_user/enable', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (59, 3, '用户禁用', 'admin/admin_user/disable', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (60, 3, '修改昵称', 'admin/admin_user/update_nickname', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (61, 3, '修改密码', 'admin/admin_user/update_password', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (62, 7, '创建角色', 'admin/admin_role/create', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (63, 7, '修改角色', 'admin/admin_role/update', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (64, 7, '启用角色', 'admin/admin_role/enable', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (65, 7, '禁用角色', 'admin/admin_role/disable', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (66, 7, '角色授权', 'admin/admin_role/access_operate', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (67, 12, '创建菜单', 'admin/admin_menu/create', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (68, 12, '修改菜单', 'admin/admin_menu/update', 'fa-list', 0, 1000, 'POST');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (69, 0, '商户管理', 'admin/merchant', 'fa-address-card-o', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (70, 69, '商户管理', 'admin/merchant/index', 'fa-asterisk', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (71, 70, '添加商户界面', 'admin/merchant/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (72, 70, '修改商户界面', 'admin/merchant/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (73, 70, '添加商户', 'admin/merchant/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (74, 70, '修改商户', 'admin/merchant/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (75, 70, '删除商户', 'admin/merchant/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (76, 0, '商品管理', 'admin/goods', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (77, 76, '商品管理', 'admin/goods/index', 'fa-product-hunt', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (78, 77, '商品添加界面', 'admin/goods/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (79, 77, '商品修改界面', 'admin/goods/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (80, 77, '商品添加', 'admin/goods/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (81, 77, '商品编辑', 'admin/goods/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (82, 77, '商品删除', 'admin/goods/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (83, 76, '商品分类', 'admin/cat-goods/index', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (84, 83, '商品分类添加界面', 'admin/cat-goods/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (85, 83, '商品分类编辑界面', 'admin/cat-goods/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (86, 83, '商品分类添加', 'admin/cat-goods/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (87, 83, '商品分类编辑', 'admin/cat-goods/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (88, 83, '商品分类删除', 'admin/cat-goods/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (89, 76, '商品评价', 'admin/goods/comment', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (90, 20, '订单编辑界面', 'admin/order/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (91, 0, '用户管理', 'admin/user', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (92, 91, '用户管理', 'admin/user/index', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (93, 92, '用户添加界面', 'admin/user/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (94, 92, '用户编辑界面', 'admin/user/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (95, 92, '用户添加', 'admin/user/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (96, 92, '用户编辑', 'admin/user/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (97, 92, '用户删除', 'admin/user/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (98, 92, '用户钱包', 'admin/user/wallet', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (99, 92, '用户积分', 'admin/user/integral', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (100, 0, '积分管理', 'admin/integral', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (101, 100, '积分记录', 'admin/integral/index', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (102, 100, '积分订单', 'admin/integral/trade', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (103, 2, '轮播图管理', 'admin/sys/banner/index', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (104, 2, '版本管理', 'admin/sys/version/index', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (105, 103, '轮播图添加界面', 'admin/sys/banner/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (106, 103, '轮播图编辑界面', 'admin/sys/banner/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (107, 103, '轮播图添加', 'admin/sys/banner/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (108, 103, '轮播图编辑', 'admin/sys/banner/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (109, 103, '轮播图删除', 'admin/sys/banner/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (110, 104, '版本添加界面', 'admin/sys/version/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (111, 104, '版本编辑界面', 'admin/sys/version/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (112, 104, '版本添加', 'admin/sys/verison/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (113, 104, '版本编辑', 'admin/sys/version/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (114, 19, '订单删除', 'admin/order/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (115, 92, '用户地址', 'admin/user/address', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (116, 92, '用户优惠券', 'admin/user/coupon', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (117, 2, '资金日志', 'admin/sys/wallet/record', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (118, 19, '退货管理', 'admin/order/process', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (119, 118, '退货审核', 'admin/order/process/verify', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (120, 118, '退货单详情', 'admin/order/process/detail', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (121, 2, '常见问题', 'admin/sys/customer/question', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (122, 2, '客户服务', 'admin/sys/customer/service', 'fa-list', 1, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (123, 121, '添加问题界面', 'admin/sys/question/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (124, 121, '编辑问题界面', 'admin/sys/question/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (126, 121, '更新问题', 'admin/sys/question/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (127, 121, '删除问题', 'admin/sys/question/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (128, 122, '添加服务界面', 'admin/sys/customer/add', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (129, 122, '服务编辑界面', 'admin/sys/customer/edit', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (130, 122, '客户服务添加', 'admin/sys/customer/create', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (131, 122, '客户服务更新', 'admin/sys/customer/update', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (132, 122, '客户服务删除', 'admin/sys/customer/del', 'fa-list', 0, 1000, '不记录');").Exec()
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (133, 117, '提现审核', 'admin/sys/record/verify', 'fa-list', 0, 1000, '不记录');").Exec()
	fmt.Println("err---", err)
}


type GoodsOrderList struct {
	GoodsOrder
	BuyName					string					`json:"buy_name"`
}

type UserWalletList struct {
	UserWallet
	UserName				string					`json:"user_name"`
}

type UserCouponList struct {
	UserCoupon
	UserName				string					`json:"user_name"`
}

type UserIntegralList struct {
	UserIntegral
	UserName				string					`json:"user_name"`
}

type IntegralRecordeList struct {
	IntegralRecord
	UserName				string					`json:"user_name"`
	SourceName				string					`json:"source_name"`
}

type IntegralTradeList struct {
	IntegralTrade
	UserName				string					`json:"user_name"`
}

type WalletRecordList struct {
	WalletRecord
	UserName				string					`json:"user_name"`
}

type OrderProcessList struct {
	OrderProcess
	UserName				string					`json:"user_name"`
	OrderNumber				string					`json:"order_number"`
	GoodsName				string					`json:"goods_name"`
	GoodsTitle				string					`json:"goods_title"`
	PayAmount				float64					`json:"pay_amount"`
}