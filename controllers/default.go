package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Head() {
	c.CustomAbort(200, "")
}

func (c *MainController) Options() {
	c.CustomAbort(200, "")
}
