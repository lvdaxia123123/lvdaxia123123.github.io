package router

import "github.com/gin-gonic/gin"

func AdminRouter(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/fenzhu", func(ctx *gin.Context) {
			ctx.JSON(200, "fenzhu")
		})
		adminRouter.GET("/list", func(ctx *gin.Context) {
			ctx.JSON(200, "list")
		})
	}

}
