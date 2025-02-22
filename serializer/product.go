package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

type ProductVO struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"` // 商品浏览次数
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(product *model.Product) ProductVO {
	return ProductVO{
		Id:            product.ID,
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.View(),
		CreatedAt:     product.CreatedAt.Unix(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossId:        product.BossId,
		BossName:      product.BossName,
		BossAvatar:    conf.Host + conf.HttpPort + conf.AvatarPath + product.BossAvatar,
	}
}
