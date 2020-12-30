package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type CustomerService struct {
	BaseService
}

func (self *CustomerService) GetPaginateData(listRows int, params url.Values) ([]*models.CustomerService, beego_pagination.Pagination) {
	self.SearchField = append(self.SearchField, new(models.CustomerService).SearchField()...)

	var cate []*models.CustomerService
	o := orm.NewOrm().QueryTable(new(models.CustomerService))
	self.WhereField = append(self.WhereField,[]string{"is_removed"}...)
	params.Add("is_removed","0")
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&cate)
	if err != nil {
		return nil, self.Pagination
	} else {
		return cate, self.Pagination
	}
}

func (self *CustomerService) GetPaginateQuestionData(listRows int, params url.Values)([]*models.Questions, beego_pagination.Pagination){
	self.SearchField = append(self.SearchField, new(models.Questions).SearchField()...)

	var cate []*models.Questions
	self.WhereField = append(self.WhereField,[]string{"is_removed"}...)
	params.Add("is_removed","0")
	o := orm.NewOrm().QueryTable(new(models.Questions))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&cate)
	if err != nil {
		return nil, self.Pagination
	} else {
		return cate, self.Pagination
	}
}