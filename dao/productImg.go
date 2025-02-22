package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

// dao/product 存储某个商品信息只包括商品的第一张图片作为封面
// dao/productImg 存储某个商品的其他图片

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(productImg).Error
}
