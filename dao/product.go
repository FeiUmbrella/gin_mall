package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{
		NewDBClient(ctx),
	}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// CreateProduct 在数据库中的Product表中创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(product).Error
}

// CountProductByCondition 查找数据库中符合condition的所有products
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	// where 可以传入map映射作为查找条件
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, basePage model.BasePage) (products []*model.Product, err error) {
	// Model(&model.product{})显示指定查询的数据库表，使用于更为复杂的操作。如关联查询Joins()
	// Find()会根据传入的products的类型来隐式判断要查询表 model.product
	err = dao.DB.Where(condition).Offset((basePage.PageNum - 1) * basePage.PageSize).Limit(basePage.PageSize).Find(&products).Error
	return
}
