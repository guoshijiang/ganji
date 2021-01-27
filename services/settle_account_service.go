package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type SettleAccountService struct {
	BaseService
}

func (Self *SettleAccountService) GetPaginateData(listRows int, params url.Values) ([]*models.MerchantSettleAccountData, beego_pagination.Pagination) {
	var data []*models.MerchantSettleAccountData
	var total int64
	om := orm.NewOrm()
	inner := "from  merchant_settle_account as t0 inner join merchant as t1 on t1.id = t0.merchant_id where t0.id > 0 "
	sql := "select t0.*,t1.merchant_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	//Self.SearchField = append(Self.SearchField, new(models.GoodsOrder).SearchField()...)
	Self.WhereField = append(Self.WhereField,[]string{"t1.merchant_name"}...)
	where,param := Self.ScopeWhereRaw(params)

	if err := om.Raw(sql1+where).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	Self.PaginateRaw(listRows,params)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by t0.created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}

func (*SettleAccountService) Create(form *form_validate.SettleAccountForm) int {
	cate := models.MerchantSettleAccount{
		MerchantId: form.MerchantId,
		AcctSeq: form.AcctSeq,
		AcctType: form.AcctType,
		AcctName: form.AcctName,
		RealName: form.RealName,
		Qrcode: form.Qrcode,
	}
	id, err := orm.NewOrm().Insert(&cate)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*SettleAccountService) GetById(id int64) *models.MerchantSettleAccount {
	o := orm.NewOrm()
	data := models.MerchantSettleAccount{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*SettleAccountService) Update(form *form_validate.SettleAccountForm) int{
	o := orm.NewOrm()
	data := models.MerchantSettleAccount{Id: form.Id}
	if o.Read(&data) == nil {
		data.AcctName = form.AcctName
		data.AcctType = form.AcctType
		data.AcctSeq = form.AcctSeq
		data.MerchantId = form.MerchantId
		data.Qrcode = form.Qrcode
		data.RealName = form.RealName
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*SettleAccountService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.MerchantSettleAccount)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}


func (Self *SettleAccountService) DailyData(listRows int, params url.Values) ([]*models.MerchantSettleDailyData, beego_pagination.Pagination) {
	var data []*models.MerchantSettleDailyData
	var total int64
	om := orm.NewOrm()
	inner := "from merchant_settle_daily t0 inner join merchant t1 on t1.id = t0.merchant_id "
	sql := "select t0.*,t1.merchant_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.WhereField = append(Self.WhereField,[]string{"t0.valid_order_num"}...)
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by t0.created desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}