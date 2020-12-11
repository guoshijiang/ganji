package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
)

type ImageFile struct {
	BaseModel
	Id        int64     `json:"id"`
	Url       string    `orm:"unique;size(256);index" json:"url"`
	ImgType   int8      `json:"img_type"` // 0:用户头像 1:商品评论图片
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

