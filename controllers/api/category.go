package api

import (
	"ganji/models"
	"ganji/types"
	type_cat "ganji/types/category"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}


// @Title CategoryList
// @Description 分类列表接口 CategoryList
// @Success 200 status bool, data interface{}, msg string
// @router /category_list [post]
func (this *CategoryController) CategoryList() {
	first_level_list, code, err := models.GetOneLevelCategoryList()
	if err != nil {
		this.Data["json"] = RetResource(false, code, nil, err.Error())
		this.ServeJSON()
		return
	}
	var category_list []type_cat.CategoryListRet
	for _, first_level := range first_level_list {
		var s_category_list []type_cat.LevelCatListRet
		s_level_list, _, _ := models.GetSecodLevelCategoryList(first_level.Id)
		for _, s_level := range s_level_list {
			s_cat := type_cat.LevelCatListRet{
				SecondCatId: s_level.Id,
				SecondCatName: s_level.Name,
			}
			s_category_list = append(s_category_list, s_cat)
		}
		category := type_cat.CategoryListRet{
			FirstCatId: first_level.Id,
			FirstCatName: first_level.Name,
			SecondCatList: s_category_list,
		}
		category_list = append(category_list, category)
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, category_list, "获取分类列表成功")
	this.ServeJSON()
	return
}

