package model

import (
	"gorm.io/gorm"
)

// 收藏夹
// todo: 数据库的外键
type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserID"` // ToB一般不适用外键，会影响性能 这里用一个是演示怎么使用
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossID"`
	BossID    uint    `gorm:"not null"`
}
