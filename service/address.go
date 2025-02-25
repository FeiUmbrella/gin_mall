package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

// Create 创建地址-api
func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address := model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(&address)
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

// List 展示所有地址-api
func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	addresses, err := addressDao.ListAddressesByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildAddresses(addresses), uint(len(addresses)))
}

// Show 展示某一地址
func (service *AddressService) Show(ctx context.Context, aId string) serializer.Response {
	code := e.Success
	aid, _ := strconv.Atoi(aId)
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(uint(aid))
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
		Data:   serializer.BuildAddress(address),
	}
}

// Update 更新某一地址-api
func (service *AddressService) Update(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	aid, _ := strconv.Atoi(aId)
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.UpdateAddressByUserId(address, uint(aid))
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

// Delete 删除某一地址-api
func (service *AddressService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	aid, _ := strconv.Atoi(aId)
	addressDao := dao.NewAddressDao(ctx)

	err := addressDao.DeleteAddressByUserId(uId, uint(aid))
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
