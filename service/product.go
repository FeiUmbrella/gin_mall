package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	ProductId      uint   `json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	CategoryId     uint   `json:"category_id" form:"category_id"`
	Title          string `json:"title" form:"title"`
	Info           string `json:"info" form:"info"`
	ImgPath        string `json:"img_path" form:"img_path"`
	Price          string `json:"price" form:"price"`
	DiscountPrice  string `json:"discount_price" form:"discount_price"`
	OnSale         bool   `json:"on_sale" form:"on_sale"` // 是否在售/上架
	Num            int    `json:"num" form:"num"`
	model.BasePage        // 分页功能
}

// Create 创建商品
func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	// 创建商品的用户
	boss, _ := userDao.GetUserById(uId)

	// 一个商品有多个展示图，以第一张作为封面图(这里上传到本地)
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("商品封面图片上传错误: ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 创建ProductDao来持久化本地product数据库
	product := model.Product{
		Name:          service.Name,
		CategoryId:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true, // 默认上架
		Num:           service.Num,
		BossId:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(&product)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("数据库创建商品信息错误：", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 并发创建商品图片
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		// todo: 这里不同商品的Name一样的话，生成的path相同，保存在本地会产生覆盖。考虑加上时间戳生成唯一path
		path, err = UploadProductToLocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			util.LogrusObj.Infoln("商品图片上传错误：", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			util.LogrusObj.Infoln("数据库创建商品图片错误：", err)
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(&product),
	}
}
