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

// List 返回商品列表
func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15 // 默认每页展示15个商品
	}
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		// 展示特定类别商品
		condition["category_id"] = service.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("数据库表中查找商品错误：", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// Search 搜索商品
func (service *ProductService) Search(ctx context.Context) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	product, count, err := productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(product), uint(count))
}

// Show 显示特定id商品的详细信息
func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pid, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pid))
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
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}
