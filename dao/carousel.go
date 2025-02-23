package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// ListCarousel 显示所有轮播图
func (dao *CarouselDao) ListCarousel() ([]model.Carousel, error) {
	var carousel []model.Carousel
	err := dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return carousel, err
}
