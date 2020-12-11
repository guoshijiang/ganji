package models

import "ganji/common"

type GoodsComment struct {
	BaseModel
	GoodsId      int64     `orm:"size(64);index" json:"goods_id"`                       // 商品ID
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
