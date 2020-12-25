package services

import (
	"ganji/common"
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"log"
	"net/url"
)

type UserService struct {
	BaseService
}

func (self *UserService) GetPaginateData(listRows int, params url.Values) ([]*models.User, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.User).SearchField()...)

	var data []*models.User
	o := orm.NewOrm().QueryTable(new(models.User))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}


func (Self *UserService) GetPaginateDataWalletRaw(listRows int, params url.Values) ([]*models.UserWalletList, beego_pagination.Pagination) {
	var data []*models.UserWalletList
	var total int64
	om := orm.NewOrm()
	inner := "from  user_wallet as t0 inner join user as t1 on t1.id = t0.user_id where t0.id > 0 "
	sql := "select t0.*,t1.user_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.UserWallet).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)
	//用户条件过滤
	userId := params.Get("user_id")
	where += " and user_id = ?"
	param = append(param,userId)

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

func (Self *UserService) GetPaginateDataIntegralRaw(listRows int, params url.Values) ([]*models.UserAddress, beego_pagination.Pagination) {
	var data []*models.UserAddress
	var total int64
	om := orm.NewOrm()
	inner := "from  user_integral as t0 inner join user as t1 on t1.id = t0.user_id where t0.id > 0 "
	sql := "select t0.*,t1.user_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.UserWallet).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)
	//用户条件过滤
	userId := params.Get("user_id")
	where += " and user_id = ?"
	param = append(param,userId)

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

func (*UserService) Create(form *form_validate.UserForm) error {
	user := &models.User{
		Phone: form.Phone,
		UserName: form.UserName,
		Email: form.Email,
		Password: form.Password,
		Avator: form.Avator,
	}
	if err := user.RegisterUserByAdmin();err != nil {
		log.Println("user--create--error",err)
		return err
	}
	return nil
}

func (*UserService) GetUserById(id int64) *models.User {
	o := orm.NewOrm()
	user := models.User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return nil
	}
	return &user
}

func (*UserService) IsExistName(user_name string, id int64) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.User)).Filter("user_name", user_name).Exist()
	} else {
		return orm.NewOrm().QueryTable(new(models.User)).Filter("user_name", user_name).Exclude("id", id).Exist()
	}
}

func (*UserService) Update(form *form_validate.UserForm) int{
	o := orm.NewOrm()
	u := models.User{Id: int64(form.Id)}
	if o.Read(&u) == nil {
		u.Phone = form.Phone
		u.UserName = form.UserName
		u.Email = form.Email
		if len(form.Password) > 0 {
			u.Password =  common.ShaOne(form.Password)
		}
		u.Avator = form.Avator
		if num,err := o.Update(&u);err != nil {
			return 0
		} else {
			return int(num)
		}
	}
	return 0
}

func (*UserService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.User)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

func (self *UserService) GetPaginateAddressData(listRows int, params url.Values) ([]*models.UserAddress, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.UserAddress).SearchField()...)

	var data []*models.UserAddress
	o := orm.NewOrm().QueryTable(new(models.UserAddress))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}