package serializer

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
)

type CartVO struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	MaxNum        int    `json:"max_num"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"` // 是否支付
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) *CartVO {
	return &CartVO{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		CreatedAt:     cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		Check:         cart.Check,
		Name:          product.Name,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []*CartVO) {

	for _, item := range items {
		bossDao := dao.NewUserDao(ctx)
		productDao := dao.NewProductDao(ctx)
		boss, _ := bossDao.GetUserById(item.BossId)
		product, _ := productDao.GetProductById(item.ProductId)

		carts = append(carts, BuildCart(item, product, boss))
	}
	return
}
