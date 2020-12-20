package models

import (
	"ganji/common"
	"ganji/types"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

type GoodsComment struct {
	BaseModel
	Id           int64     `json:"id"`
	GoodsId      int64     `orm:"size(64)" json:"goods_id"`                          // 商品ID
	UserId       int64     `orm:"default(1);" json:"user_id"`                           // 评论用户
	Title        string    `orm:"size(512);index" json:"title" form:"title"`            // 评论标题
	Star         int8      `orm:"default(5);index" json:"star"`                         // 评论级别 1-10 没增加一个数字代表半星
 	Content      string    `orm:"type(text)" json:"content"`                            // 评论内容
 	ImgOneId     int64     `orm:"size(64);index" json:"img_one_id"`                     // 评论图片 1
	ImgTwoId     int64     `orm:"size(64);index" json:"img_two_id"`                     // 评论图片 2
	ImgThreeId   int64     `orm:"size(64);index" json:"img_three_id"`                   // 评论图片 3
}

func (this *GoodsComment) TableName() string {
	return common.TableName("goods_comment")
}

func (this *GoodsComment) SearchField() []string {
	return []string{"title"}
}

func (this *GoodsComment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GoodsComment) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *GoodsComment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *GoodsComment) Insert() (error, int64) {
	id, err := orm.NewOrm().Insert(this)
	if err != nil {
		return err, 0
	}
	return nil, id
}

func GetGoodsCommentList(page, pageSize int, goods_id int64) ([]*GoodsComment, int64, error) {
	offset := (page - 1) * pageSize
	gct_list := make([]*GoodsComment, 0)
	query := orm.NewOrm().QueryTable(GoodsComment{}).Filter("GoodsId", goods_id)
	total, _ := query.Count()
	_, err := query.Limit(pageSize, offset).All(&gct_list)
	if err != nil {
		return nil, types.SystemDbErr, errors.New("查询数据库失败")
	}
	return gct_list, total, nil
}


