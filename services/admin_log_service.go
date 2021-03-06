package services

import (
	"encoding/json"
	"ganji/common/utils"
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"net/url"
	"time"
)

type AdminLogService struct {
	BaseService
}

//创建操作日志
func (*AdminLogService) CreateAdminLog(loginUser *models.AdminUser, menu *models.AdminMenu, url string, ctx *context.Context) {
	var adminLog models.AdminLog

	if loginUser == nil {
		adminLog.AdminUserId = 0
	} else {
		adminLog.AdminUserId = loginUser.Id
	}
	adminLog.Name = menu.Name
	adminLog.LogMethod = menu.LogMethod
	adminLog.Url = url
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreateTime = int(time.Now().Unix())
	adminLog.UpdateTime = int(time.Now().Unix())

	o := orm.NewOrm()
	//开启事务
	err := o.Begin()

	adminLogId, err := o.Insert(&adminLog)
	if err != nil {
		o.Rollback()
		beego.Error(err)
		return
	}

	//adminLogData数据表添加数据
	jsonData, _ := json.Marshal(ctx.Request.PostForm)
	cryptData := utils.Encrypt(jsonData, []byte(beego.AppConfig.String("log_aes_key")))
	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogId)
	adminLogData.Data = cryptData
	_, err = o.Insert(&adminLogData)
	if err != nil {
		o.Rollback()
		beego.Error(err)
		return
	}
	o.Commit()
}

//登录日志
func (*AdminLogService) LoginLog(loginUserId int, ctx *context.Context) {
	var adminLog models.AdminLog
	adminLog.AdminUserId = loginUserId
	adminLog.Name = "登录"
	adminLog.Url = "admin/auth/login"
	adminLog.LogMethod = "POST"
	adminLog.LogIp = ctx.Input.IP()
	adminLog.CreateTime = int(time.Now().Unix())
	adminLog.UpdateTime = int(time.Now().Unix())

	o := orm.NewOrm()

	//开启事务
	err := o.Begin()

	adminLogId, err := o.Insert(&adminLog)
	if err != nil {
		o.Rollback()
		beego.Error(err)
		return
	}

	//adminLogData数据表添加数据
	jsonData, _ := json.Marshal(ctx.Request.PostForm)
	cryptData := utils.Encrypt(jsonData, []byte(beego.AppConfig.String("log_aes_key")))

	var adminLogData models.AdminLogData
	adminLogData.AdminLogId = int(adminLogId)
	adminLogData.Data = cryptData
	_, err = o.Insert(&adminLogData)
	if err != nil {
		o.Rollback()
		beego.Error(err)
		return
	}
	o.Commit()
}

//获取admin_log 总数
func (*AdminLogService) GetCount() int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminLog)).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

//获取所有adminuser
func (this *AdminLogService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminLog, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	this.SearchField = append(this.SearchField, new(models.AdminLog).SearchField()...)
	this.WhereField = append(this.WhereField, new(models.AdminLog).WhereField()...)
	this.TimeField = append(this.TimeField, new(models.AdminLog).TimeField()...)

	var adminLog []*models.AdminLog
	o := orm.NewOrm().QueryTable(new(models.AdminLog))
	_, err := this.PaginateAndScopeWhere(o, listRows, params).All(&adminLog)
	if err != nil {
		return nil, this.Pagination
	} else {
		return adminLog, this.Pagination
	}
}
