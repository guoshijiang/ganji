package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
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

func (*CustomerService) CreateService(form *form_validate.CustomerServiceForm) int {
	data := models.CustomerService{
		UserName: form.UserName,
		Phone: form.Phone,
		WeiChat: form.WeiChat,
		WcQrcode: form.WcQrcode,
		Type: form.Type,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*CustomerService) GetServiceById(id int64) *models.CustomerService {
	o := orm.NewOrm()
	data := models.CustomerService{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}

func (*CustomerService) UpdateService(form *form_validate.CustomerServiceForm) int{
	o := orm.NewOrm()
	data := models.CustomerService{Id: form.Id}
	if o.Read(&data) == nil {
		if len(form.WcQrcode) > 0 {
			data.WcQrcode = form.WcQrcode
		}
		data.UserName = form.UserName
		data.Phone = form.Phone
		data.WeiChat = form.WeiChat
		data.Type = form.Type
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}


func (*CustomerService) CreateQuestion(form *form_validate.QuestionForm) int {
	data := models.Questions{
		Author: form.Author,
		QsDetail: form.QsDetail,
		QsTitle: form.QsTitle,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*CustomerService) GetQuestionById(id int64) *models.Questions {
	o := orm.NewOrm()
	data := models.Questions{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}

func (*CustomerService) UpdateQuestion(form *form_validate.QuestionForm) int{
	o := orm.NewOrm()
	data := models.Questions{Id: form.Id}
	if o.Read(&data) == nil {
		data.QsTitle = form.QsTitle
		data.QsDetail = form.QsDetail
		data.Author = form.Author
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*CustomerService) DelQuestion(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Questions)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

func (*CustomerService) DelCustomer(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.CustomerService)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}