package form_validate

type VersionForm struct {
	Id           int64     `form:"id"`
	VersionNum   string    `form:"version_num"`
	Platforms    int64     `form:"platforms"`      			// 0: 安卓 1: IOS
	Decribe      string    `form:"decribe"`
	DownloadUrl  string    `form:"download_url"`
	IsForce      int64     `form:"is_force"`   				// 0: 不强制更新 1: 强制更新
	IsRemove      int64    `form:"is_remove"`
	IsCreate 	 int       `form:"_create"`
}