package router

import (
	"modulename/controller/api"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("users", api.UserController{}.List)
		apiRouter.GET("users/:id", api.UserController{}.Delete)
		apiRouter.POST("users", api.UserController{}.Index)
		apiRouter.DELETE("users/:id", api.UserController{}.Add)

	}
}
