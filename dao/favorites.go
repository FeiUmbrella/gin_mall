package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) ListFavorite(uId uint) (favorites []*model.Favorite, err error) {
	err = dao.Model(&model.Favorite{}).Where("user_id = ?", uId).Find(&favorites).Error
	return
}

func (dao *FavoriteDao) FavoriteExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ? AND product_id = ?", uId, pId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count >= 1, nil
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Create(favorite).Error
	return
}

func (dao *FavoriteDao) DeleteFavorite(uId, fId uint) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("id = ? AND user_id = ?", fId, uId).Delete(&model.Favorite{}).Error
	return
}
