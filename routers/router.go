package routers

import (
	"behvosting/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/my_files", &controllers.MyFilesController{})
	beego.Router("/file", &controllers.FileController{})
	beego.Router("/reg", &controllers.RegController{})
	beego.Router("/login", &controllers.LoginController{})
}
