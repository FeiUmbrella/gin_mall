package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// ListCategory 显示所有类别
func (dao *CategoryDao) ListCategory() ([]*model.Category, error) {
	var carousel []*model.Category
	err := dao.DB.Model(&model.Category{}).Find(&carousel).Error
	return carousel, err
}
