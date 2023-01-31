package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"wink/model"
	"wink/service"
)

func Register(g *gin.Context) {
	db := service.InitDB()
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
