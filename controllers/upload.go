package controllers

import (
	"behvosting/lib"
	"github.com/beego/bee/logger"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Post() {
	userName := c.UserName
	fileHeaders, err := c.GetFiles("file-name")
	if err != nil {
		beeLogger.Log.Error("Error get files")
		return
	}
	for _, fileHeader := range fileHeaders {
		lib.SaveFile(userName, fileHeader)
	}
	c.Redirect("../my_files", 303)
}
