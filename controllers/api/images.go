package api

import (
	"ganji/models"
	"ganji/types"
	"github.com/astaxie/beego"
	"os"
	"path"
	"time"
)

type ImageController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
}

// @Title UploadFiles
// @Description 上传图片 UploadFiles
// @Success 200 status bool, data interface{}, msg string
// @router /upload_file [post]
func (this *ImageController) UploadFiles() {
	f, h, err := this.GetFile("file")
	if err != nil {
		this.Data["json"] = RetResource(false, types.GetImagesFileFail, nil, "获取文件失败")
		this.ServeJSON()
		return
	}
	defer f.Close()
	ext := path.Ext(h.Filename)
	var AllowExtMap map[string]bool = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true}
	if _, ok := AllowExtMap[ext]; !ok {
		this.Data["json"] = RetResource(false, types.FileFormatError, nil, "上传的文件格式不符合要求")
		this.ServeJSON()
		return
	}
	var Filebytes = 1 << 24 //文件小于16兆
	if fileSizer, ok := f.(Sizer); ok {
		fileSize := fileSizer.Size()
		if fileSize > int64(Filebytes) {
			this.Data["json"] = RetResource(false, types.FileIsBig, nil, "文件太大了")
			this.ServeJSON()
		} else {
			img_dir := beego.AppConfig.String("uplaod_url")
			time_str := time.Now().Format("2006/01/02/")
			uploadDir := img_dir + time_str
			err = os.MkdirAll(uploadDir, os.ModePerm)
			if err != nil {
				this.Data["json"] = RetResource(false, types.CreateFilePathError, nil, "文件路径创建失败")
				this.ServeJSON()
				return
			}
			fpath := uploadDir + h.Filename
			err = this.SaveToFile("file", fpath)
			if err != nil {
				this.Data["json"] = RetResource(false, types.FileIsBig, nil, "保存文件成功")
				this.ServeJSON()
				return
			}
			img_file := models.ImageFile{
				Url: beego.AppConfig.String("image_path") + time_str + h.Filename,
			}
			err = img_file.Insert()
			if err != nil {
				this.Data["json"] = RetResource(true, types.SaveFileFail, img_file, "保存文件失败")
				this.ServeJSON()
				return
			}
			this.Data["json"] = RetResource(true, types.ReturnSuccess, img_file, "上传文件成功")
			this.ServeJSON()
			return
		}
	} else {
		this.Data["json"] = RetResource(false, types.FileIsBig, nil, "上传文件太大")
		this.ServeJSON()
		return
	}
}

