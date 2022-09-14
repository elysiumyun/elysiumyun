package model

import "gorm.io/gorm"

type UserMeta struct {
	gorm.Model
}

type UserData struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	UserMeta
	UserData
}
