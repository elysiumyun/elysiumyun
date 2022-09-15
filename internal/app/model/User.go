package model

import "gorm.io/gorm"

type UserMeta struct {
	gorm.Model
}

type UserData struct {
	UserName string `json:"username" xml:"username"`
	Password string `json:"password" xml:"password"`
}

type User struct {
	UserMeta
	UserData
}
