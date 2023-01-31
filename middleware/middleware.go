package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wink/model"
	"wink/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		tokenString := g.GetHeader("Authorization")
		//token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			g.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "no auth",
			})
			g.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := service.ParseToken(tokenString)
		if err != nil || token.Vaild {
			g.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "not auth",
			})
			g.Abort()
			return
		}

		//userId
		userId := claims.UserId
		DB := service.InitDB()
		var user model.User
		DB.First(&user, userId)

		//user
		if user.ID == 0 {
			g.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "not auth",
			})
			g.Abort()
			return
		}
	}
}
