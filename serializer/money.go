package serializer

import (
	"gin_mall/model"
	"gin_mall/pkg/util"
)

type MoneyVO struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) *MoneyVO {
	util.Encrypt.SetKey(key)
	return &MoneyVO{
		UserId:   user.ID,
		UserName: user.UserName,
		// 数据库中money是加密状态，需要解密显示
		UserMoney: util.Encrypt.AesDecoding(user.Money),
	}
}
