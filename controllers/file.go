package controllers

type FileController struct {
	BaseController
}

func (c *FileController) Get() {
	fileName := c.GetString("name")
	fullPath := "./uploads/" + c.UserName + "/" + fileName + ".png"
	c.Ctx.Output.Download(fullPath)
}
