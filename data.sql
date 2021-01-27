/*
 Navicat Premium Data Transfer

 Source Server         : my8.0
 Source Server Type    : MySQL
 Source Server Version : 80014
 Source Host           : localhost:3356
 Source Schema         : ganji

 Target Server Type    : MySQL
 Target Server Version : 80014
 File Encoding         : 65001

 Date: 27/01/2021 00:19:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `url` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `icon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'fa-list',
  `is_show` tinyint(4) NOT NULL DEFAULT '1',
  `sort_id` int(11) NOT NULL DEFAULT '1000',
  `log_method` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '不记录',
  PRIMARY KEY (`id`),
  KEY `admin_menu_url` (`url`)
) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
BEGIN;
INSERT INTO `admin_menu` VALUES (1, 0, '后台首页', 'admin/index/index', 'fa-home', 1, 99, '不记录');
INSERT INTO `admin_menu` VALUES (2, 0, '系统管理', 'admin/sys', 'fa-desktop', 1, 1099, '不记录');
INSERT INTO `admin_menu` VALUES (3, 2, '用户管理', 'admin/admin_user/index', 'fa-user', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (4, 3, '添加用户界面', 'admin/admin_user/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (5, 3, '修改用户界面', 'admin/admin_user/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (6, 3, '删除用户', 'admin/admin_user/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (7, 2, '角色管理', 'admin/admin_role/index', 'fa-group', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (8, 7, '添加角色界面', 'admin/admin_role/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (9, 7, '修改角色界面', 'admin/admin_role/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (10, 7, '删除角色', 'admin/admin_role/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (11, 7, '角色授权界面', 'admin/admin_role/access', 'fa-key', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (12, 2, '菜单管理', 'admin/admin_menu/index', 'fa-align-justify', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (13, 12, '添加菜单界面', 'admin/admin_menu/add', 'fa-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (14, 12, '修改菜单界面', 'admin/admin_menu/edit', 'fa-edit', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (15, 12, '删除菜单', 'admin/admin_menu/del', 'fa-close', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (16, 2, '操作日志', 'admin/admin_log/index', 'fa-keyboard-o', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (17, 16, '日志详情', 'admin/admin_log/view', 'fa-search-plus', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (18, 2, '个人资料', 'admin/admin_user/profile', 'fa-smile-o', 1, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (19, 0, '订单管理', 'admin/order/mange', 'fa-first-order', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (20, 19, '订单管理', 'admin/order/index', 'fa-first-order', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (55, 3, '修改头像', 'admin/admin_user/update_avatar', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (56, 3, '添加用户', 'admin/admin_user/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (57, 3, '修改用户', 'admin/admin_user/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (58, 3, '用户启用', 'admin/admin_user/enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (59, 3, '用户禁用', 'admin/admin_user/disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (60, 3, '修改昵称', 'admin/admin_user/update_nickname', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (61, 3, '修改密码', 'admin/admin_user/update_password', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (62, 7, '创建角色', 'admin/admin_role/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (63, 7, '修改角色', 'admin/admin_role/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (64, 7, '启用角色', 'admin/admin_role/enable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (65, 7, '禁用角色', 'admin/admin_role/disable', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (66, 7, '角色授权', 'admin/admin_role/access_operate', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (67, 12, '创建菜单', 'admin/admin_menu/create', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (68, 12, '修改菜单', 'admin/admin_menu/update', 'fa-list', 0, 1000, 'POST');
INSERT INTO `admin_menu` VALUES (69, 0, '商户管理', 'admin/merchant', 'fa-address-card-o', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (70, 69, '商户管理', 'admin/merchant/index', 'fa-asterisk', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (71, 70, '添加商户界面', 'admin/merchant/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (72, 70, '修改商户界面', 'admin/merchant/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (73, 70, '添加商户', 'admin/merchant/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (74, 70, '修改商户', 'admin/merchant/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (75, 70, '删除商户', 'admin/merchant/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (76, 0, '商品管理', 'admin/goods', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (77, 76, '商品管理', 'admin/goods/index', 'fa-product-hunt', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (78, 77, '商品添加界面', 'admin/goods/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (79, 77, '商品修改界面', 'admin/goods/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (80, 77, '商品添加', 'admin/goods/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (81, 77, '商品编辑', 'admin/goods/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (82, 77, '商品删除', 'admin/goods/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (83, 76, '商品分类', 'admin/cat-goods/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (84, 83, '商品分类添加界面', 'admin/cat-goods/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (85, 83, '商品分类编辑界面', 'admin/cat-goods/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (86, 83, '商品分类添加', 'admin/cat-goods/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (87, 83, '商品分类编辑', 'admin/cat-goods/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (88, 83, '商品分类删除', 'admin/cat-goods/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (89, 76, '商品评价', 'admin/goods/comment', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (90, 20, '订单编辑界面', 'admin/order/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (91, 0, '用户管理', 'admin/user', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (92, 91, '用户管理', 'admin/user/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (93, 92, '用户添加界面', 'admin/user/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (94, 92, '用户编辑界面', 'admin/user/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (95, 92, '用户添加', 'admin/user/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (96, 92, '用户编辑', 'admin/user/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (97, 92, '用户删除', 'admin/user/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (98, 92, '用户钱包', 'admin/user/wallet', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (99, 92, '用户积分', 'admin/user/integral', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (100, 0, '积分管理', 'admin/integral', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (101, 100, '积分记录', 'admin/integral/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (102, 100, '积分订单', 'admin/integral/trade', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (103, 2, '轮播图管理', 'admin/sys/banner/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (104, 2, '版本管理', 'admin/sys/version/index', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (105, 103, '轮播图添加界面', 'admin/sys/banner/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (106, 103, '轮播图编辑界面', 'admin/sys/banner/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (107, 103, '轮播图添加', 'admin/sys/banner/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (108, 103, '轮播图编辑', 'admin/sys/banner/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (109, 103, '轮播图删除', 'admin/sys/banner/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (110, 104, '版本添加界面', 'admin/sys/version/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (111, 104, '版本编辑界面', 'admin/sys/version/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (112, 104, '版本添加', 'admin/sys/verison/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (113, 104, '版本编辑', 'admin/sys/version/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (114, 19, '订单删除', 'admin/order/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (115, 92, '用户地址', 'admin/user/address', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (116, 92, '用户优惠券', 'admin/user/coupon', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (117, 2, '资金日志', 'admin/sys/wallet/record', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (118, 19, '退货管理', 'admin/order/process', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (119, 118, '退货审核', 'admin/order/process/verify', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (120, 118, '退货单详情', 'admin/order/process/detail', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (121, 2, '常见问题', 'admin/sys/customer/question', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (122, 2, '客户服务', 'admin/sys/customer/service', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (123, 121, '添加问题界面', 'admin/sys/question/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (124, 121, '编辑问题界面', 'admin/sys/question/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (125, 121, '添加问题', 'admin/sys/question/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (126, 121, '更新问题', 'admin/sys/question/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (127, 121, '删除问题', 'admin/sys/question/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (128, 122, '添加服务界面', 'admin/sys/customer/add', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (129, 122, '服务编辑界面', 'admin/sys/customer/edit', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (130, 122, '客户服务添加', 'admin/sys/customer/create', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (131, 122, '客户服务更新', 'admin/sys/customer/update', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (132, 122, '客户服务删除', 'admin/sys/customer/del', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (133, 117, '提现审核', 'admin/sys/record/verify', 'fa-list', 0, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (134, 77, '商品属性', 'admin/goods_type/index', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (135, 134, '商品属性-添加属性-界面', 'admin/goods_type/add', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (136, 134, '商品属性-添加属性-创建', 'admin/goods_type/create', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (137, 134, '商品属性-编辑属性-界面', 'admin/goods_type/edit', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (138, 134, '商品属性-编辑属性-修改', 'admin/goods_type/update', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (139, 134, '商品属性删除', 'admin/goods_type/del', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (140, 0, '结算管理', 'admin/settle', 'fa-list', 1, 1000, '不记录');
INSERT INTO `admin_menu` VALUES (141, 140, '结算配置', 'admin/settle/index', 'fa-list', 1, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (142, 141, '配置管理-添加配置-界面', 'admin/settle/add', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (143, 141, '结算配置-添加配置-创建', 'admin/settle/create', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (144, 141, '结算配置-编辑配置-界面', 'admin/settle/edit', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (145, 141, '结算配置-编辑配置-更新', 'admin/settle/update', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (146, 141, '结算配置-删除', 'admin/settle/del', 'fa-list', 0, 1000, 'GET');
INSERT INTO `admin_menu` VALUES (147, 140, '结算日报', 'admin/settle/daily', 'fa-chain-broken', 1, 1000, 'GET');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
