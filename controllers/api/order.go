package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_comment "ganji/types/comment"
	"github.com/astaxie/beego"
	"strings"
)

type OrderController struct {
	beego.Controller
}


// @Title CreateOrder finished
// @Description 添加评论 CreateOrder
// @Success 200 status bool, data interface{}, msg string
// @router /create_order [post]
func (this *OrderController) CreateOrder() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	requestUser, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}

	var add_comment type_comment.AddCommentCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &add_comment); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := add_comment.AddCommentCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != add_comment.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	cmt := models.GoodsComment{
		GoodsId: add_comment.GoodsId,
		UserId: add_comment.UserId,
		Title: add_comment.Title,
		Star: add_comment.Star,
		Content: add_comment.Content,
		ImgOneId: add_comment.ImgOneId,
		ImgTwoId: add_comment.ImgTwoId,
		ImgThreeId: add_comment.ImgThreeId,
	}
	err, id := cmt.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, err.Error(), "添加评论失败")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, map[string]interface{}{"id": id}, "添加评论成功")
		this.ServeJSON()
		return
	}
}
