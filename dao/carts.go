package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

// CreateCart 创建购物车记录
func (dao *CartDao) CreateCart(cart *model.Cart) (err error) {
	return dao.Model(&model.Cart{}).Create(cart).Error
}

// ListCartByUserId 找到uId的所有购物车内容
func (dao *CartDao) ListCartByUserId(uId uint) (carts []*model.Cart, err error) {
	err = dao.Model(&model.Cart{}).Where("user_id = ?", uId).Find(&carts).Error
	return
}

// DeleteCartByUserId 通过Id删除某一条购物车记录
func (dao *CartDao) DeleteCartByUserId(uId, cId uint) (err error) {
	err = dao.Model(&model.Cart{}).Where("id = ? AND user_id = ?", cId, uId).Delete(&model.Cart{}).Error
	return
}

// UpdateCartNumById 通过Id更新购物车中某一商品数量
func (dao *CartDao) UpdateCartNumById(num uint, cId uint) error {
	return dao.Model(&model.Cart{}).Where("id=?", cId).Update("num", num).Error
}
