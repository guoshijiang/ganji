package services

import (
	beego_pagination "ganji/common/utils/beego-pagination"
	"ganji/form_validate"
	"ganji/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type VersionService struct {
	BaseService
}

func (Self *VersionService) GetPaginateData(listRows int, params url.Values) ([]*models.Version, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.Version).SearchField()...)

	var data []*models.Version
	o := orm.NewOrm().QueryTable(new(models.Version))
	_, err := Self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, Self.Pagination
	} else {
		return data, Self.Pagination
	}
}

func (*VersionService) Create(form *form_validate.VersionForm) int {
	cate := models.Version{
		VersionNum: form.VersionNum,
		Platforms: form.Platforms,
		Decribe: form.Decribe,
		DownloadUrl: form.DownloadUrl,
		IsForce: form.IsForce,
		IsRemove: 1,
	}
	id, err := orm.NewOrm().Insert(&cate)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*VersionService) GetVersionById(id int64) *models.Version {
	o := orm.NewOrm()
	data := models.Version{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}

func (*VersionService) IsExistVersionNum(num string, id int64) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.Version)).Filter("version_num", num).Exist()
	} else {
		return orm.NewOrm().QueryTable(new(models.Version)).Filter("version_num", num).Exclude("id", id).Exist()
	}
}

func (*VersionService) Update(form *form_validate.VersionForm) int{
	o := orm.NewOrm()
	data := models.Version{Id: form.Id}
	if o.Read(&data) == nil {
		data.VersionNum = form.VersionNum
		data.Platforms = form.Platforms
		data.Decribe = form.Decribe
		data.DownloadUrl = form.DownloadUrl
		data.IsForce = form.IsForce
		data.IsRemove = form.IsRemove
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*VersionService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Version)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}