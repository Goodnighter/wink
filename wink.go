package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	Telephone string
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	fmt.Print(user)
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(127.0.0.1:3306)/wink?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	fmt.Print("db success")
	return db
}

func main() {
	r := gin.Default()
	db := InitDB()
	print(db)
	r.POST("/api/auth/register", func(g *gin.Context) {
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
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		g.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})
	panic(r.Run(":9090"))
}
