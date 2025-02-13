package dao

import (
	"fmt"
	"gin_mall/model"
)

// 迁移
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Address{},
			&model.Admin{},
			&model.Category{},
			&model.Carousel{},
			&model.Cart{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.ProductImg{},
			&model.Product{},
		)
	if err != nil {
		fmt.Println("err:", err)
	}
	return
}
