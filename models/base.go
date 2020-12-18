package models

import (
	"fmt"
	"github.com/astaxie/beego"
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
		new(GoodsComment), new(GoodsCat), new(GoodsImage), new(GoodsOrder), new(GroupOrder), new(ImageFile),
		new(IntegralRecord), new(IntegralTrade), new(UserAddress), new(Version), new(WalletRecord), new(Banner))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//orm.RunSyncdb(mysqlConfig["db_alias"], true, true)
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
	_, err = orm.NewOrm().Raw("INSERT INTO `admin_menu` VALUES (89, 77, '商品评价', 'admin/goods/comment', 'fa-list', 0, 1000, '不记录');").Exec()
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

type UserIntegralList struct {
	UserIntegral
	UserName				string					`json:"user_name"`
}

type IntegralRecordeList struct {
	IntegralRecord
	UserName				string					`json:"user_name"`
	SourceName				string					`json:"source_name"`
}