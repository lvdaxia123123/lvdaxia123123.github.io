package main

import (
	router "modulename/myrouter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置路由
	r.GET("/gets", func(c *gin.Context) {
		c.String(200, "gets")
	})
	//域名/news?aid=20
	r.GET("/geta", func(c *gin.Context) {
		aid := c.Query("aid")
		c.JSON(200, gin.H{
			"news": aid,
		})
	})
	//动态路由/user/20
	r.GET("/news/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"news": id,
		})
	})
	//返回一个JSON 数据
	r.GET("/json", func(ctx *gin.Context) {
		ctx.Query("aid")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "json",
		})
	})
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务

	//返回结构体

	r.GET("/structJson", func(ctx *gin.Context) {
		var msg struct {
			Username string `json:"username"`
			Msg      string `json:"msg"`
			Age      string `json:"age"`
		}
		msg.Username = "name1"
		msg.Msg = "msg1"
		msg.Age = "18"
		ctx.JSON(200, msg)

	})
	//返回JSOPN
	r.GET("/jsonp", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		ctx.JSONP(http.StatusOK, data)
	})

	//渲染模板
	r.LoadHTMLFiles("index.html")
	r.GET("/index", func(c *gin.Context) {
		// HTML请求
		// 模板的渲染
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "imustctf.top",
		})
	})

	//Get 请求传值GET /user?uid=20&page=1
	r.GET("/user", func(c *gin.Context) {
		uid := c.Query("uid")
		page := c.DefaultQuery("page", "4")
		c.String(200, "uid=%v page=%v", uid, page)
	})
	//Post 请求传值 获取 form 表单数据
	r.POST("/doAddUser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "29")
		ctx.JSON(200, gin.H{
			"usernmae": username, "password": password, "age": age,
		})
	})
	//Get 传值绑定到结构体
	type Userinfo struct {
		Username string `form:"username" json:"user"`
		Password string `form:"password" json:"password"`
	}
	r.GET("/userinfo", func(ctx *gin.Context) {
		var userinfo Userinfo

		if err := ctx.ShouldBind(&userinfo); err == nil {
			ctx.JSON(http.StatusOK, userinfo)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	//Post 传值绑定到结构体
	r.POST("/doLogin", func(ctx *gin.Context) {
		var userinfo Userinfo
		if err := ctx.ShouldBind(&userinfo); err == nil {
			ctx.JSON(http.StatusOK, userinfo)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})
	//路由分离
	// apiRouter := r.Group("/api")
	// {
	// 	apiRouter.GET("fenzhu", func(ctx *gin.Context) {
	// 		ctx.JSON(200, "fenzhu")
	// 	})
	// 	apiRouter.GET("articles", func(ctx *gin.Context) {
	// 		ctx.JSON(200, "articles")
	// 	})
	// }

	router.AdminRouter(r)
	router.ApiRouter(r)
	r.Run(":8001")
}
