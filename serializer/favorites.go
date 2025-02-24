package serializer

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
)

type FavoriteVO struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	BossId        uint   `json:"boss_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

// 这里需要预加载Product和User表才能用外键引用对应的Product和User的信息。
// Name:item.Product.Name,
func BuildFavorite(item *model.Favorite, product *model.Product, boss *model.User) *FavoriteVO {
	return &FavoriteVO{
		UserId:        item.UserID,
		ProductId:     item.ProductID,
		BossId:        boss.ID,
		CreatedAt:     item.CreatedAt.Unix(),
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []*FavoriteVO) {

	for _, item := range items {
		bossDao := dao.NewUserDao(ctx)
		productDao := dao.NewProductDao(ctx)
		boss, _ := bossDao.GetUserById(item.BossID)
		product, _ := productDao.GetProductById(item.ProductID)

		favorites = append(favorites, BuildFavorite(item, product, boss))
	}
	return
}
