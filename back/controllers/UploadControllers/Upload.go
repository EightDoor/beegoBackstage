package UploadControllers

import (
	"beegoBackstage/models/FileModels"
	"beegoBackstage/utils"
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"math/rand"
	"os"
	"path"
	"time"
)

type UploadController struct {
	utils.BaseController
}

// Get 查询单条
// @router /upload/:id [get]
func (c *UploadController) Get() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.RError(utils.R{
			Msg: err.Error(),
		})
	}
	var file FileModels.File
	o := orm.NewOrm()
	fileErr := o.QueryTable(FileModels.File{}).Filter("id", id).One(&file)
	if fileErr != nil {
		c.RError(utils.R{Msg: fileErr.Error()})
	} else {
		c.RSuccess(utils.R{
			Data: file,
		})
	}
}

// FileUpload 上传文件
// @router /upload [post]
func (c *UploadController) FileUpload() {
	var file FileModels.File
	f, h, err := c.GetFile("file")
	if err == nil {
		ext := path.Ext(h.Filename)
		// 创建目录
		uploadDir := "static/upload/" + time.Now().Format("2006_01_02") + "/"
		dirErr := os.MkdirAll(uploadDir, 0777)
		if dirErr != nil {
			logs.Error(dirErr.Error(), "文件创建失败")
		}
		// 文件名称
		rand.Seed(time.Now().UnixNano())
		randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
		hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
		fileName := fmt.Sprintf("%x", hashName) + ext
		logs.Debug(fileName, "文件名称")
		fPath := uploadDir + fileName
		logs.Debug(fPath, "存储路径")
		defer f.Close() // 关闭上传的文件，不然会出现临时文件不能清除的情况
		saveErr := c.SaveToFile("file", fPath)
		if saveErr != nil {
			logs.Error(saveErr.Error(), "saveErr")
			c.RError(utils.R{
				Msg: saveErr.Error(),
			})
		}
		file.FileName = h.Filename
		file.FileSize = h.Size
		configPath, _ := web.AppConfig.String("uploadUrl")
		env, _ := web.AppConfig.String("runmode")
		file.Url = configPath + fileName
		if env == "dev" {
			// 开发环境直接拼接存储地址
			file.Url = configPath + fPath
		}
		file.StoragePath = fPath
		c.RSuccess(utils.R{Msg: "上传成功", Data: file})
	} else {
		c.RError(utils.R{Msg: err.Error(), Data: "文件获取失败"})
	}
}

// FileDel 删除文件
// @router /upload/:id [delete]
func (c *UploadController) FileDel() {
	id, err := c.GetInt(":id")
	if err != nil {
		c.RError(utils.R{
			Msg:  err.Error(),
			Data: "请输入删除的文件id",
		})
	}
	var file FileModels.File
	o := orm.NewOrm()
	o.QueryTable(FileModels.File{}).Filter("id", id).One(&file)
	result, dErr := o.QueryTable(FileModels.File{}).Filter("id", id).Delete()
	// 删除文件
	if dErr != nil {
		c.RError(utils.R{
			Msg:  dErr.Error(),
			Data: result,
		})
	} else {
		// 处理，如果删除成功了
		if result != 0 {
			if file.Id != 0 {
				removeErr := os.Remove(file.StoragePath)
				if removeErr != nil {
					c.RError(utils.R{
						Msg:  removeErr.Error(),
						Data: "os文件删除失败",
					})
				} else {
					c.RSuccess(utils.R{
						Msg: "删除成功",
					})
				}
			} else {
				c.RError(utils.R{
					Msg: "文件不存在",
				})
			}

		} else {
			c.RSuccess(utils.R{
				Data: result,
			})
		}
	}
}
