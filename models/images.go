package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
	"time"
)

type ImageFile struct {
	Id        int64     `json:"id"`
	Url       string    `orm:"unique;size(256);index" json:"url"`
	IsRemoved int8      `orm:"index" json:"is_removed"` //0: 正常，1: 删除
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now_add;type(datetime);index" json:"updated_at"`
}

func (this *ImageFile) TableName() string {
	return common.TableName("user_image")
}

func (this *ImageFile) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *ImageFile) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *ImageFile) GetImageById(id int64) (*ImageFile, int, error) {
	var image ImageFile
	err := image.Query().Filter("Id", id).One(&image)
	if err != nil {
		return nil, 100, err
	}
	return &image, types.ReturnSuccess, nil
}

