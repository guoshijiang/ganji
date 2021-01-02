package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type IntegralService struct {
	BaseService
}

func (Self *IntegralService) GetPaginateDataRecordRaw(listRows int, params url.Values) ([]*models.UserWalletList, beego_pagination.Pagination) {
	var data []*models.UserWalletList
	var total int64
	om := orm.NewOrm()
	inner := "from integral_record as t0 inner join user as t1 on t1.id = t0.user_id inner join user as t2 on t2.id = t0.source_user_id where t0.id > 0 "
	sql := "select t0.*,t1.user_name,t2.user_name source_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.IntegralRecord).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	Self.PaginateRaw(listRows,params)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}


func (Self *IntegralService) GetPaginateDataTradeRaw(listRows int, params url.Values) ([]*models.IntegralTradeList, beego_pagination.Pagination) {
	var data []*models.IntegralTradeList
	var total int64
	om := orm.NewOrm()
	inner := "from integral_trade as t0 inner join user as t1 on t1.id = t0.user_id where t0.id > 0 "
	sql := "select t0.*,t1.user_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.IntegralRecord).SearchField()...)
	params.Add("t0.is_removed","0")
	Self.WhereField = append(Self.WhereField,"t0.is_removed")
	where,param := Self.ScopeWhereRaw(params)

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	Self.PaginateRaw(listRows,params)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}