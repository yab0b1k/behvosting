package controllers

import (
	"behvosting/lib"
	"github.com/beego/bee/logger"
	"strconv"
)

type MyFilesController struct {
	BaseController
}

type MyFile struct {
	Url  string
	Time string
}

func (c *MyFilesController) Get() {
	//ip := c.Ctx.Input.Host()
	images := lib.GetUserImages(c.UserName)
	listFiles := make([]MyFile, len(images))
	for i, image := range images {
		beeLogger.Log.Error("Found image: " + image.ImageID)
		myFile := MyFile{}
		myFile.Url = c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/file?name=" + image.ImageID
		myFile.Time = image.CreateTime.Format("15:04:05 02.01.2006")
		listFiles[i] = myFile
	}

	/*userDir := "./uploads/" + ip + "/"
	files, err := ioutil.ReadDir(userDir)
	if err != nil {
		beeLogger.Log.Error("Error open dir " + userDir + " " + err.Error())
	} else {
		sort.Slice(files, func(i, j int) bool {
			return files[j].ModTime().Before(files[i].ModTime())
		})
		for i, file := range files {
			myFile := MyFile{}
			myFile.Url = c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/file?name=" + file.Name()
			myFile.Time = file.ModTime().Format("15:04:05 02.01.2006")
			listFiles[i] = myFile
		}
	}*/
	c.Data["Files"] = listFiles
	c.TplName = "my_files.html"
}
