package services

import (
	"errors"
	"fmt"
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
	"strconv"
	"strings"
	"time"
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
	data := models.MerchantSettleAccount{
		MerchantId: form.MerchantId,
		AcctSeq: form.AcctSeq,
		AcctType: form.AcctType,
		AcctName: form.AcctName,
		RealName: form.RealName,
		Qrcode: form.Qrcode,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (Self *SettleAccountService) CreateSettle(form *form_validate.SettleForm) int {
	data := models.MerchantSettle{
		MerchantId: form.MerchantId,
		SettleAccountId: form.SettleAccountId,
		StartSettleTime: form.StartSettleTime,
		EndSettleTime: form.EndSettleTime,
		SettleAmount: form.SettleAmount,
		Status: form.Status,
		HandUser: AdminUserVal.Username,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*SettleAccountService) UpdateSettle(form *form_validate.SettleForm) int {
	o := orm.NewOrm()
	data := models.MerchantSettle{Id: form.Id}
	if o.Read(&data) == nil {
		data.Status = form.Status
		data.HandUser = AdminUserVal.Username
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
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
		if len(form.Qrcode) > 0 {
			data.Qrcode = form.Qrcode
		}
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
	Self.SearchField = append(Self.SearchField,[]string{"valid_order_num"}...)
	params.Add("_merchant_id",strconv.Itoa(int(AdminUserVal.Id)))
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by t0.created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}

func (Self *SettleAccountService) SettleData(listRows int, params url.Values) ([]*models.MerchantSettleData, beego_pagination.Pagination) {
	var data []*models.MerchantSettleData
	var total int64
	om := orm.NewOrm()
	inner := "from  merchant_settle t0 inner join merchant t1 on t1.id = t0.merchant_id where t0.id > 0 "
	sql := "select t0.* " + inner
	sql1 := "select count(*) total " + inner

	keyword := params.Get("create_time")
	if len(keyword) > 0 {
		keywords := strings.Split(keyword," - ")
		params.Add("start_settle_time;gte",keywords[0])
		params.Add("end_settle_time;lte",keywords[1])
		Self.WhereField = append(Self.WhereField,[]string{"start_settle_time;gte","end_settle_time;lte"}...)
	}

	where,param := Self.ScopeWhereRaw(params)

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
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

//结算数据 开始于结算表的最后一条时间
func (Self *SettleAccountService) SettleDailyData(params url.Values,end *time.Time) ([]*models.DailySettleData, error) {
	o := orm.NewOrm()
	var (
		sql string
		condition []interface{}
		data []*models.DailySettleData
	)
	stm := strings.Split(params.Get("create_time")," - ")
	if AdminUserVal.MerchantId == 0 {
		_,err := orm.NewOrm().Raw("SET sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION'").Exec()
		fmt.Println(err)
		sql = "select sum(valid_order_amount) amounts,merchant_id from merchant_settle_daily where created_at between ? and ? group by merchant_id "
		condition = append(condition,stm[0],stm[1])
	} else {
		//时间判断
		settles,_ := Self.GetSettleNew(AdminUserVal.MerchantId)
		if len(settles) > 0 {
			//判断时间
			searchEnd,_ := time.ParseInLocation("2006-01-02",stm[1],time.Local)
			if settles[0].EndSettleTime.Unix() > searchEnd.Unix() { //已经结算过了
				return data,errors.New("搜索时间错误")
			} else {
				stm[0] = settles[0].EndSettleTime.Format("2006-01-02")
			}
		}
		sql = "select sum(valid_order_amount) amounts from merchant_settle_daily where merchant_id = ? and created_at between ? and ? "
		condition = append(condition,AdminUserVal.MerchantId,stm[0],stm[1])
	}

	if _,err := o.Raw(sql,condition).QueryRows(&data);err != nil {
		return data,err
	}
	return data,nil
}

func (*SettleAccountService) GetList() []*models.MerchantSettleAccount {
	var data []*models.MerchantSettleAccount
	var err error
	if AdminUserVal.MerchantId == 0 {
		_, err = orm.NewOrm().QueryTable(new(models.MerchantSettleAccount)).All(&data)
	} else {
		_, err = orm.NewOrm().QueryTable(new(models.MerchantSettleAccount)).Filter("merchant_id",AdminUserVal.MerchantId).All(&data)
	}
	if err != nil {
		return nil
	}
	return data
}

//获得merchant_settle 商户最新结算的时间
func (*SettleAccountService) GetSettleNew(merchant_id  int) ([]*models.MerchantSettle,error) {
	var data []*models.MerchantSettle
	if _,err := orm.NewOrm().QueryTable("merchant_settle").Filter("merchant_id",merchant_id).OrderBy("-end_settle_time").All(&data);err != nil {
		return data,err
	}
	return data,nil
}
