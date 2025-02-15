package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

const (
	PasswordCost        = 12       // 密码加密程度
	Active       string = "active" // 激活用户
)

// 密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes) // 加密后的密码
	return nil
}
