package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"` // todo: userid? cartid?
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       uint `json:"num" form:"num"`
}

// Create 创建购物车-api
func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	// 判断有没有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 判断有没有对应boss
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart := &model.Cart{
		UserId:    uId,
		ProductId: product.ID,
		BossId:    product.BossId,
		Num:       service.Num,
	}

	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

// List 展示购物车内容
func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, carts),
	}
}

// Update 更新某一购物车-api
func (service *CartService) Update(ctx context.Context, uId uint, cId string) serializer.Response {
	code := e.Success
	cid, _ := strconv.Atoi(cId)
	cartDao := dao.NewCartDao(ctx)
	// 购物车更新只能更新数量~
	err := cartDao.UpdateCartNumById(service.Num, uint(cid))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Delete 删除某一购物车-api
func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	code := e.Success
	cid, _ := strconv.Atoi(cId)
	cartDao := dao.NewCartDao(ctx)

	err := cartDao.DeleteCartByUserId(uId, uint(cid))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
