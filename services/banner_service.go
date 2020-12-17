package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type BannerService struct {
	BaseService
}

func (Self *BannerService) GetPaginateData(listRows int, params url.Values) ([]*models.Banner, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.Banner).SearchField()...)

	var data []*models.Banner
	o := orm.NewOrm().QueryTable(new(models.Banner))
	_, err := Self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, Self.Pagination
	} else {
		return data, Self.Pagination
	}
}

func (*BannerService) Create(form *form_validate.BannerForm) int {
	cate := models.Banner{
		Avator: form.Avator,
		Url: form.Url,
		IsDispay: form.IsDispay,
	}
	id, err := orm.NewOrm().Insert(&cate)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*BannerService) GetBannerById(id int64) *models.Banner {
	o := orm.NewOrm()
	data := models.Banner{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}

func (*BannerService) Update(form *form_validate.BannerForm) int{
	o := orm.NewOrm()
	data := models.Banner{Id: form.Id}
	if o.Read(&data) == nil {
		if len(form.Avator) > 0 {
			data.Avator = form.Avator
		}
		data.Url = form.Url
		data.IsDispay = form.IsDispay
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*BannerService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Banner)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}