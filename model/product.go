package model

import (
	"gin_mall/cache"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string // 商家头像
}

// View 商品浏览次数
func (product *Product) View() uint64 {
	// 找到Redis中存储关于某一商品的键值对，返回对应的值
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 增加商品浏览次数
func (product *Product) AddView() {
	// +1 操作
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))

	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
