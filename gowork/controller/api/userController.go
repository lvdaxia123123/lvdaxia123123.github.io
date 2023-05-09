package api

import "github.com/gin-gonic/gin"

//结构体中的继承
type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	con.success(c)
}

func (con UserController) Add(c *gin.Context) {
	c.String(200, "api UserAdd")
}

func (con UserController) List(c *gin.Context) {
	con.success(c)
}

func (con UserController) Delete(c *gin.Context) {
	c.String(200, "api UserDelete")
}
