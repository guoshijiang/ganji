package comment

import (
	"ganji/types"
	"github.com/pkg/errors"
)

type AddCommentCheck struct {
	GoodsId      int64  `json:"goods_id"`
	UserId       int64  `json:"user_id"`
	Title        string `json:"title"`
	Star         int8   `json:"star"`
	Content      string `json:"content"`
	ImgOneId     int64  `json:"img_one_id"`
	ImgTwoId     int64  `json:"img_two_id"`
	ImgThreeId   int64  `json:"img_three_id"`
}

func (this AddCommentCheck) AddCommentCheckParamValidate() (int, error) {
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("商品ID不能小于0")
	}
	if this.UserId <= 0 {
		return types.ParamLessZero, errors.New("用户ID不能小于0")
	}
	if this.Title == "" {
		return types.ParamEmptyError, errors.New("评论标题为空")
	}
	if this.Star <= 0 {
		return types.ParamEmptyError, errors.New("评论星级不能小于0")
	}
	if this.Content == "" {
		return types.ParamEmptyError, errors.New("评论内容为空")
	}
	if this.ImgOneId < 0 {
		return types.ParamEmptyError, errors.New("评论图片一ID小于0")
	}
	if this.ImgTwoId < 0 {
		return types.ParamEmptyError, errors.New("评论图片二ID小于0")
	}
	if this.ImgThreeId < 0 {
		return types.ParamEmptyError, errors.New("评论图片三ID小于0")
	}
	return types.ReturnSuccess, nil
}

type DelCommentCheck struct {
	CommentId  int64 `json:"comment_id"`
	UserId     int64 `json:"user_id"`
}

func (this DelCommentCheck) DelCommentCheckParamValidate() (int, error) {
	if this.CommentId <= 0 || this.UserId <= 0 {
		return types.ParamLessZero, errors.New("评论ID和用户ID不能小于0")
	}
	return types.ReturnSuccess, nil
}

type CommentListCheck struct {
	types.PageSizeData
	GoodsId  int64 `json:"goods_id"`
}

func (this CommentListCheck) CommentListCheckParamValidate() (int, error) {
	code, err := this.PageSizeDataParamValidate()
	if err != nil {
		return code, err
	}
	if this.GoodsId <= 0 {
		return types.ParamLessZero, errors.New("评论ID和用户ID不能小于0")
	}
	return types.ReturnSuccess, nil
}




