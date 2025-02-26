package service

import (
	"context"
	"fmt"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"gorm.io/gorm"
	"strconv"
)

type OrderPay struct {
	OrderId   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductId uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	BossName  string  `json:"boss_name" form:"boss_name"`
	Num       int     `json:"num" form:"num"`
	Key       string  `json:"key" form:"key"` // 支付金额？
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	err := dao.NewOrderDao(ctx).Transaction(func(tx *gorm.DB) error { // 事务  当闭包函数返回不为nil则自动回滚

		//todo:这里商家和买家加密的key是同一个，正式地应该分为两个key1，key2分别进行加密
		util.Encrypt.SetKey(service.Key)
		orderDao := dao.NewOrderDaoByDB(tx)
		order, err := orderDao.GetOrderByOId(uId, service.OrderId) //找到对应订单
		if err != nil {
			code = e.Error
			return err
		}
		money := order.Money
		num := order.Num
		money = money * float64(num) // 计算订单金额

		// 买家
		userDao := dao.NewUserDaoByDB(tx)
		user, err := userDao.GetUserById(uId)
		if err != nil {
			code = e.Error
			return err
		}
		// 对买家的余额解密，并扣除订单金额，再加密保存
		moneyStr := util.Encrypt.AesDecoding(user.Money)
		moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
		if moneyFloat < money { // 余额不足
			code = e.Error
			return err
		}
		// 加密保存
		finMoney := fmt.Sprintf("%f", moneyFloat-money)
		user.Money = util.Encrypt.AesEncoding(finMoney)
		err = userDao.UpdateUserById(uId, user)
		if err != nil {
			code = e.Error
			return err
		}

		// 卖家
		var boss *model.User
		boss, err = userDao.GetUserById(service.BossId)
		// 对卖家的余额解密，并打入订单金额，再加密保存
		moneyStr = util.Encrypt.AesDecoding(boss.Money)
		moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
		// 加密保存
		finMoney = fmt.Sprintf("%f", moneyFloat+money)
		boss.Money = util.Encrypt.AesEncoding(finMoney)
		err = userDao.UpdateUserById(boss.ID, boss)
		if err != nil {
			code = e.Error
			return err
		}

		// 对应商品数目-num
		var product *model.Product
		productDao := dao.NewProductDaoByDB(tx)
		product, err = productDao.GetProductById(service.ProductId)
		if err != nil {
			code = e.Error
			return err
		}
		product.Num -= num
		// 更新对应商品
		err = productDao.UpdateProductById(service.ProductId, product)
		if err != nil {
			tx.Rollback()
			code = e.Error
			return err
		}

		// 将订单Type设置为2（已支付）
		order.Type = 2
		err = orderDao.UpdateOrderById(service.OrderId, order)
		if err != nil {
			code = e.Error
			return err
		}

		// 买家买完商品后，数据库中买家对应的product的数目要改变
		userProduct := &model.Product{
			Name:          product.Name,
			CategoryId:    product.CategoryId,
			Title:         product.Title,
			Info:          product.Info,
			ImgPath:       product.ImgPath,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			OnSale:        false,
			Num:           num,
			BossId:        uId,
			BossName:      user.UserName,
			BossAvatar:    user.Avatar,
		}
		err = productDao.CreateProduct(userProduct)
		if err != nil {
			tx.Rollback()
			code = e.Error
			return err
		}

		return nil
	})

	if err != nil {
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
