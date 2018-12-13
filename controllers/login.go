package controllers

import (
	"behvosting/lib"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io"
)

type LoginController struct {
	beego.Controller
}

//Login process
func (c *LoginController) Post() {
	username := c.Ctx.Request.Form.Get("username")
	password := c.Ctx.Request.Form.Get("password")
	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	//now := time.Now().Format("2006-01-02 15:04:05")

	session := c.StartSession()
	userInfo := lib.GetUserInfo(username)
	if userInfo.Pass == newPass {
		session.Set("uid", userInfo.ID)
		session.Set("uname", userInfo.Name)

		c.Ctx.Redirect(302, "/my_files")
	} else {
		c.Ctx.Redirect(302, "/")
	}
}
