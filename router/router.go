package router

import (
	"github.com/gin-gonic/gin"
	"wink/api"
	"wink/middleware"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), api.GetUserInfo)
	return r
}
