package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string // 是否被封禁
	Avatar         string // 头像
	Money          string // 用string存储金额的密文
}
