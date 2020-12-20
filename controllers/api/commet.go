package api

import (
	"encoding/json"
	"ganji/models"
	"ganji/types"
	type_comment "ganji/types/comment"
	"github.com/astaxie/beego"
	"strings"
)

type CommentController struct {
	beego.Controller
}


// @Title AddCommet finished
// @Description 添加评论 AddCommet
// @Success 200 status bool, data interface{}, msg string
// @router /add_comment [post]
func (this *CommentController) AddCommet() {
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


// @Title DelCommet finished
// @Description 删除评论 DelCommet
// @Success 200 status bool, data interface{}, msg string
// @router /del_commet [post]
func (this *CommentController) DelCommet() {
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
	var del_comment type_comment.DelCommentCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &del_comment); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := del_comment.DelCommentCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	if requestUser.Id != del_comment.UserId {
		this.Data["json"] = RetResource(false, types.UserIsNotExist, err, "Token 和用户不匹配，拒绝添加地址")
		this.ServeJSON()
		return
	}
	var gdc models.GoodsComment
	gdc.Id = del_comment.CommentId
	err = gdc.Delete()
	if err != nil {
		this.Data["json"] = RetResource(false, types.AddressIsEmpty, err, "删除评论失败")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, err, "删除评论成功")
	this.ServeJSON()
	return
}


// @Title GetCommentList finished
// @Description 获取地址列表 GetCommentList
// @Success 200 status bool, data interface{}, msg string
// @router /comment_list [post]
func (this *CommentController) GetCommentList() {
	bearerToken := this.Ctx.Input.Header(HttpAuthKey)
	if len(bearerToken) == 0 {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	token := strings.TrimPrefix(bearerToken, "Bearer ")
	u_tk, err := models.GetUserByToken(token)
	if err != nil {
		this.Data["json"] = RetResource(false, types.UserToKenCheckError, nil, "您还没有登陆，请登陆")
		this.ServeJSON()
		return
	}
	var clist type_comment.CommentListCheck
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &clist); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		this.ServeJSON()
		return
	}
	if code, err := clist.CommentListCheckParamValidate(); err != nil {
		this.Data["json"] = RetResource(false, code, err, err.Error())
		this.ServeJSON()
		return
	}
	clst, total, err := models.GetGoodsCommentList(clist.Page, clist.PageSize, clist.GoodsId)
	if err !=  nil {
		this.Data["json"] = RetResource(false, types.SystemDbErr, nil, err.Error())
		this.ServeJSON()
		return
	}
	var cmt_list []type_comment.CommentListRep
	for _, v := range clst {
		var img_mdl_one, img_mdl_two, img_mdl_three   models.ImageFile
		img_mdl_one.Id = v.ImgOneId
		img_mdl_two.Id = v.ImgTwoId
		img_mdl_three.Id = v.ImgThreeId
		var one_url, two_url, three_url string
		ImgOne_img, _, _ := img_mdl_one.GetImageById(v.ImgOneId)
		if ImgOne_img != nil {
			one_url = ImgOne_img.Url
		} else {
			one_url = ""
		}
		ImgTwo_img, _, _ := img_mdl_one.GetImageById(v.ImgTwoId)
		if ImgTwo_img != nil {
			two_url = ImgTwo_img.Url
		} else {
			two_url = ""
		}
		ImgThree_img, _, _ := img_mdl_one.GetImageById(v.ImgThreeId)
		if ImgThree_img != nil {
			three_url = ImgThree_img.Url
		} else {
			three_url = ""
		}
		cl := type_comment.CommentListRep{
			Id: v.Id,
			UserName: u_tk.UserName,
			UserPho: u_tk.Avator,
			GoodsId: v.GoodsId,
			UserId: v.UserId,
			Title: v.Title,
			Star: v.Star,
			Content: v.Content,
			ImgOne: one_url,
			ImgTwo: two_url,
			ImgThree: three_url,
			CreateTime: v.CreatedAt,
		}
		cmt_list = append(cmt_list, cl)
	}
	data := map[string]interface{}{
		"total": total,
		"cmt_lst": cmt_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取评论列表成功")
	this.ServeJSON()
	return
}


