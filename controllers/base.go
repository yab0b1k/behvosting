package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/beego/bee/logger"
)

type BaseController struct {
	beego.Controller
	Session  session.Store
	UserName string
}

func (c *BaseController) Prepare() {
	if c.Ctx.Request.RequestURI == "/reg" {
		return
	}
	c.Session = c.StartSession()
	userName, ok := c.Session.Get("uname").(string)
	if !ok {
		beeLogger.Log.Warn("Session username is empty. Need login.")
		c.Ctx.Redirect(302, "/")
	}
	c.UserName = userName
	c.Data["Username"] = userName
}
