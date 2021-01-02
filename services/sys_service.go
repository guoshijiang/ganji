package services

import (
	"encoding/json"
	"fmt"
	"ganji/models"
	beego_pagination "ganji/common/utils/beego-pagination"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type SysService struct {
	BaseService
}


func (Self *SysService) GetPaginateDataWalletRecordRaw(listRows int, params url.Values) ([]*models.WalletRecordList, beego_pagination.Pagination){
	var data []*models.WalletRecordList
	var total int64
	om := orm.NewOrm()
	inner := "from  wallet_record as t0 left join user as t1 on t1.id = t0.user_id where t0.id > 0 "
	sql := "select t0.*,t1.user_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.WalletRecord).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)

	if AdminUserVal.MerchantId > 0 {
		where += " and t1.merchant_id = ? "
		param = append(param,AdminUserVal.MerchantId)
	}

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	Self.PaginateRaw(listRows,params)

	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}

	ju,_ := json.Marshal(Self.Pagination)
	fmt.Println("Pagination",string(ju))
	return data,Self.Pagination
}