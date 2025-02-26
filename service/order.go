package service

import (
	"context"
	"fmt"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  string  `json:"order_num" form:"order_num"` // 订单编号
	Type      int     `json:"type" form:"type"`           // 是否支付、是否评价等
	model.BasePage
}

// Create 创建订单-api
func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success

	orderDao := dao.NewOrderDao(ctx)
	order := &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, // 默认未支付
	}
	// 检验地址是否存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressId = address.ID
	// 利用时间戳种子，生成9位编码，再拼接唯一标识的商品ID和用户ID
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	oderNum, err := strconv.ParseUint(number, 10, 64)
	order.OrderNum = oderNum

	err = orderDao.CreateOrder(order)
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

// List 展示所有订单
func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	// 分页相关
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uId
	orders, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(total))
}

// Show 显示某一订单详细内容-api
func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	code := e.Success
	oid, _ := strconv.Atoi(oId)

	// 拿到对应订单
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderByOId(uId, uint(oid))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 拿到订单对应的地址信息
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(order.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 拿到订单对应的商品信息
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
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
		Data:   serializer.BuildOrder(order, product, address),
	}
}

// Delete 删除某一订单-api
func (service *OrderService) Delete(ctx context.Context, uId uint, oId string) serializer.Response {
	code := e.Success
	oid, _ := strconv.Atoi(oId)
	orderDao := dao.NewOrderDao(ctx)

	err := orderDao.DeleteOrderByUserId(uId, uint(oid))
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
