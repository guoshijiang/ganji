package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type OrderService struct {
	BaseService
}

func (Self *OrderService) GetPaginateData(listRows int, params url.Values) ([]*models.GoodsOrder, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.GoodsOrder).SearchField()...)

	var order []*models.GoodsOrder
	o := orm.NewOrm().QueryTable(new(models.GoodsOrder))
	_, err := Self.PaginateAndScopeWhere(o, listRows, params).All(&order)
	if err != nil {
		return nil, Self.Pagination
	} else {
		return order, Self.Pagination
	}
}


func (Self *OrderService) GetPaginateDataRaw(listRows int, params url.Values) ([]*models.GoodsOrderList, beego_pagination.Pagination) {
	var data []*models.GoodsOrderList
	var total int64
	om := orm.NewOrm()
	inner := "from  goods_order as t0 inner join user as t1 on t1.id = t0.user_id where t0.id > 0 "
	sql := "select t0.*,t1.user_name buy_user " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.GoodsOrder).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)
	if err := om.Raw(sql1+where).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}


func (*OrderService) Del(ids []int) int{
	var (
		count  int64
		err	error
	)
	if AdminUserVal.MerchantId > 0 {
		count, err = orm.NewOrm().QueryTable(new(models.GoodsOrder)).Filter("id__in", ids).Delete()
	} else {
		count, err = orm.NewOrm().QueryTable(new(models.GoodsOrder)).Filter("id__in", ids).Delete()
	}
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}