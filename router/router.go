package router

import (
	"github.com/gin-gonic/gin"
	"wink/api"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", api.Register)
	return r
}
