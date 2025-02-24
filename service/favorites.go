package service

import (
	"context"
	"fmt"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"strconv"
)

type FavoritesService struct {
	FavoriteId uint `json:"favorite_id" from:"favorite_id"`
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	model.BasePage
}

// List 显示收藏内容(接口)
func (service *FavoritesService) List(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	favorites, err := favoriteDao.ListFavorite(uId)
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))
}

// Create 创建收藏内容(接口)
func (service *FavoritesService) Create(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	exist, err := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}
	// 收藏已存在
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
		}
	}

	// 商品收藏不存在
	// 找到对应user
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}
	// 找到对应boss
	fmt.Println("bossId:", service.BossId)
	fmt.Println("productId:", service.ProductId)
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}
	// 找到对应product
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserID:    uId,
		Product:   *product,
		ProductID: service.ProductId,
		Boss:      *boss,
		BossID:    service.BossId,
	}

	// 创建收藏
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Msg:    e.GetMsg(code),
		Status: code,
	}
}

// Delete 删除收藏内容(接口)
func (service *FavoritesService) Delete(ctx context.Context, uId uint, fId string) serializer.Response {
	favoriteId, _ := strconv.Atoi(fId)
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	err := favoriteDao.DeleteFavorite(uId, uint(favoriteId))
	if err != nil {
		util.LogrusObj.Info("err", err)
		code = e.Error
		return serializer.Response{
			Msg:    e.GetMsg(code),
			Status: code,
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Msg:    e.GetMsg(code),
		Status: code,
	}
}
