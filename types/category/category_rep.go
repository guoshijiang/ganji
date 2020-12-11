package category

type LevelCatListRet struct {
	SecondCatId   int64  `json:"second_cat_id"`
	SecondCatName string `json:"second_cat_name"`
}

type CategoryListRet struct {
	FirstCatId int64    `json:"first_cat_id"`
	FirstCatName string `json:"first_cat_name"`
	SecondCatList []LevelCatListRet `json:"second_cat_list"`
}
