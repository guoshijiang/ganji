package models


import (
	"ganji/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Version struct {
	Id           int64     `json:"id" form:"id"`
	VersionNum   string    `json:"version_num" form:"version_num"`
	Platforms    int64     `json:"platforms" form:"platforms"`              // 0: 安卓 1: IOS
	Decribe      string    `json:"decribe" form:"decribe"`
	DownloadUrl  string    `json:"download_url" form:"download_url"`
	IsForce      int64     `orm:"index" json:"is_force" form:"is_force"`   // 0: 不强制更新 1: 强制更新
	IsRemove     int64     `orm:"index" json:"is_remove"`                  // 0: 删除 1: 不删除
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime);index"`
	UpdatedAt    time.Time `orm:"auto_now_add;type(datetime);index"`
}

func (this *Version) TableName() string {
	return common.TableName("version")
}

func (this *Version) SearchField() []string {
	return []string{"version_num"}
}

func (this *Version) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Version) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Version) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

//获得某个版本的信息
func (this *Version) FetchOne() *Version{
	err := orm.NewOrm().Read(this)
	if err != nil {
		return nil
	}
	return this
}

//获得版本控制列表
func (this *Version) FetchRows(page,pageSize int,condition *orm.Condition) ([]*Version,int64,error) {
	var versions  []*Version
	offset := (page - 1) * pageSize
	o := orm.NewOrm().QueryTable(this.TableName()).SetCond(condition)
	total,err := o.Count()
	if err != nil {
		return nil,0,err
	}
	_,err = o.OrderBy("-id").Limit(pageSize,offset).All(&versions)
	if err != nil {
		return nil,0,err
	}
	return versions,total,nil
}

func (this *Version) GetVersionInfo() (*Version, error) {
	version := Version{}
	err := orm.NewOrm().QueryTable(this.TableName()).
		Filter("Platforms", this.Platforms).
		OrderBy("-id").Limit(1).
		One(&version)
	if err != nil {
		logs.Info(err)
		return nil, err
	}
	return &version, nil
}
