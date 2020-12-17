package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
)

type GoodsServices struct {
	BaseService
}

func (self *GoodsServices) GetPaginateData(listRows int, params url.Values) ([]*models.Goods, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.Goods).SearchField()...)
	if AdminUserVal.MerchantId > 0 {
		params.Set("_merchant_id", strconv.Itoa(AdminUserVal.MerchantId))
	}
	var goods []*models.Goods
	o := orm.NewOrm().QueryTable(new(models.Goods))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&goods)
	if err != nil {
		return nil, self.Pagination
	} else {
		return goods, self.Pagination
	}
}

func (*GoodsServices) Create(form *form_validate.GoodsForm) int {
	goods := models.Goods{
		GoodsName: form.GoodsName,
		GoodsParams: form.GoodsParams,
		GoodsDetail: form.GoodsDetail,
		Discount: form.Discount,
		Sale: form.Sale,
		Title: form.Title,
		IsHot: form.IsHot,
		IsDisplay: form.IsDisplay,
		Logo: form.Logo,
		GoodsMark: form.GoodsMark,
		IsIgExchange:form.IsIgExchange,
		IsGroup: form.IsGroup,
		IsIntegral: form.IsIntegral,
		IsLimitTime: form.IsLimitTime,
		GoodsPrice: form.GoodsPrice,
		GoodsDisPrice: form.GoodsDisPrice,
		Serveice: form.Serveice,
		CalcWay:form.CalcWay,
		GoodsIntegral:form.GoodsIntegral,
		LeftAmount:form.LeftAmount,
		TotalAmount: form.TotalAmount,
		MerchantId: int64(AdminUserVal.MerchantId),
	}
	id, err := orm.NewOrm().Insert(&goods)

	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*GoodsServices) GetGoodsById(id int64) *models.Goods {
	o := orm.NewOrm()
	var goods models.Goods
	if AdminUserVal.MerchantId > 0 {
		goods = models.Goods{Id: id, MerchantId: int64(AdminUserVal.MerchantId)}
	} else {
		goods = models.Goods{Id:id}
	}
	err := o.Read(&goods)
	if err != nil {
		return nil
	}
	return &goods
}

func (*GoodsServices) IsExistName(goods_name string, id int64) bool {
	if id == 0 {
		if AdminUserVal.MerchantId  > 0 {
			return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Filter("merchant_id",AdminUserVal.MerchantId).Exist()
		}
		return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Exist()
	} else {
		if AdminUserVal.MerchantId > 0 {
			return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Filter("merchant_id", AdminUserVal).Exclude("id", id).Exist()
		}
		return orm.NewOrm().QueryTable(new(models.Goods)).Filter("goods_name", goods_name).Exclude("id", id).Exist()
	}
}

func (*GoodsServices) Update(form *form_validate.GoodsForm) int{
	o := orm.NewOrm()
	goods := models.Goods{Id: form.Id}
	if o.Read(&goods) == nil {
		goods.GoodsName = form.GoodsName
		goods.GoodsParams = form.GoodsParams
		goods.GoodsDetail = form.GoodsDetail
		goods.Logo = form.Logo
		num, err := o.Update(&goods)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*GoodsServices) Del(ids []int) int{
	var (
		count  int64
		err	error
	)
	if AdminUserVal.MerchantId > 0 {
		count, err = orm.NewOrm().QueryTable(new(models.Goods)).Filter("id__in", ids).Delete()
	} else {
		count, err = orm.NewOrm().QueryTable(new(models.Goods)).Filter("id__in", ids).Delete()
	}
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}


func (Self *GoodsServices) GetPaginateCommentData(listRows int, params url.Values) ([]*models.GoodsComment, beego_pagination.Pagination) {
	var data []*models.GoodsComment
	var total int64
	om := orm.NewOrm()
	inner := "from goods_comment as t0 inner join goods as t1 on t1.id = t0.goods_id where t0.goods_id > 0 "
	sql := "select t0.* " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.IntegralRecord).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)
	if AdminUserVal.MerchantId > 0 {
		where += " and t1.merchant_id = ? "
		param = append(param,AdminUserVal.MerchantId)
	}
	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}


