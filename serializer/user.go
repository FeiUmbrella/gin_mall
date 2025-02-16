package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

// VO-Value Object:用于表示层展示数据
// 前端所展示的数据类型
type UserVO struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"` // 用户的类型
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 转换为前端所需要的数据类型
func BuildUser(user *model.User) *UserVO {
	return &UserVO{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
