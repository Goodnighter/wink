package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Password  string
	Telephone string
}

type UserInfo struct {
	Name  string
	Phone string
}

func ToUserInfo(user User) UserInfo {
	return UserInfo{
		user.Name,
		user.Telephone,
	}
}
