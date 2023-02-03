package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"wink/middleware"
	"wink/model"
	"wink/service"
)

var db = service.InitDB()

func Register(g *gin.Context) {
	name := g.PostForm("name")
	password := g.PostForm("password")
	telephone := g.PostForm("telephone")
	//
	if len(telephone) != 11 {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    402,
			"message": "telephone error",
		})
		return
	}
	if len(password) < 6 {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    402,
			"message": "password error",
		})
		return
	}
	if len(name) == 0 {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    402,
			"message": "name error",
		})
		return
	}

	if IsTelephoneExist(db, telephone) {
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    402,
			"message": "user exist",
		})
		return
	}
	//new user
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	g.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	fmt.Print(user)
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Login(g *gin.Context) {
	//params
	phone := g.PostForm("phone")
	//password := g.PostForm("password")
	//query
	var user model.User
	db.Where("telephone = ?", phone).First(&user)
	if user.ID == 0 {
		g.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "mas": "not exist"})
		return
	}
	token, err := service.ReleaseToken(user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "error",
		})
		log.Printf("token error : %v", err)
		return
	}
	g.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
	})
}

func GetUserInfo(g *gin.Context) {
	user, _ := g.Get("user")
	type UserInfo struct {
		Name string
	}
	middleware.Response(g, http.StatusOK, 200, gin.H{
		"user": model.ToUserInfo(user.(model.User)),
	}, "")
}
