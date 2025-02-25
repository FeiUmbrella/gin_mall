package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

// CreateAddress 创建地址记录
func (dao *AddressDao) CreateAddress(address *model.Address) (err error) {
	return dao.Model(&model.Address{}).Create(address).Error
}

// ListAddressesByUserId 找到uId的所有地址记录
func (dao *AddressDao) ListAddressesByUserId(uId uint) (addresses []*model.Address, err error) {
	err = dao.Model(&model.Address{}).Where("user_id = ?", uId).Find(&addresses).Error
	return
}

// GetAddressById 通过Id找到某一条地址记录
func (dao *AddressDao) GetAddressById(aId uint) (address *model.Address, err error) {
	err = dao.Model(&model.Address{}).Where("id = ?", aId).First(&address).Error
	return
}

// UpdateAddressByUserId 通过Id更新某一条地址记录
func (dao *AddressDao) UpdateAddressByUserId(address *model.Address, aId uint) (err error) {
	err = dao.Model(&model.Address{}).Where("id = ? AND user_id = ?", aId, address.UserID).Updates(&address).Error
	return
}

// DeleteAddressByUserId 通过Id删除某一条地址记录
func (dao *AddressDao) DeleteAddressByUserId(uId, aId uint) (err error) {
	err = dao.Model(&model.Address{}).Where("id = ? AND user_id = ?", aId, uId).Delete(&model.Address{}).Error
	return
}
