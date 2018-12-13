package controllers

import (
	"behvosting/lib"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
)

type RegController struct {
	BaseController
}

func (c *RegController) Get() {
	c.TplName = "reg.html"
}

//Registration process
func (c *RegController) Post() {
	c.TplName = "reg.html"
	userName := c.Ctx.Request.Form.Get("username")
	password := c.Ctx.Request.Form.Get("password")
	usererr := checkUsername(userName)
	fmt.Println(usererr)
	if usererr == false {
		c.Data["UsernameErr"] = "Username error, Please to again"
		return
	}

	passerr := checkPassword(password)
	if passerr == false {
		c.Data["PasswordErr"] = "Password error, Please to again"
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	//now := time.Now().Format("2006-01-02 15:04:05")

	userInfo := lib.GetUserInfo(userName)

	if userInfo.Name == "" {
		userInfo = lib.CreateNewUser(userName, newPass)

		//Set the session successful login
		//sess, _ := beego.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
		c.Session.Set("uid", userInfo.ID)
		c.Session.Set("uname", userInfo.Name)
		c.Ctx.Redirect(302, "/")
	} else {
		c.Data["UsernameErr"] = "User already exists"
	}

}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}
