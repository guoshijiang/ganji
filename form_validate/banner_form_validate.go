package form_validate


type BannerForm struct {
	Id           int64     `form:"id"`
	Avator       string    `form:"avator"`
	Url          string    `form:"url"`
	IsDispay     int8      `form:"is_dispay"`
	IsCreate	 int	   `form:"_create"`
}