package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

// CreateOrder 创建订单记录
func (dao *OrderDao) CreateOrder(order *model.Order) (err error) {
	return dao.Model(&model.Order{}).Create(order).Error
}

// ListOrderByCondition 找到满足筛选条件的订单
func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.Model(&model.Order{}).
		Where(condition).
		Count(&total).Error
	if err != nil {
		return
	}
	
	err = dao.Model(&model.Order{}).
		Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Find(&orders).Error
	return
}

// DeleteOrderByUserId 通过Id删除某一条订单记录
func (dao *OrderDao) DeleteOrderByUserId(uId, cId uint) (err error) {
	err = dao.Model(&model.Order{}).Where("id = ? AND user_id = ?", cId, uId).Delete(&model.Order{}).Error
	return
}

// GetOrderByOId 通过oId获取某一订单
func (dao *OrderDao) GetOrderByOId(uId, oId uint) (order *model.Order, err error) {
	err = dao.Model(&model.Order{}).Where("id=? AND user_id=?", oId, uId).First(&order).Error
	return
}
