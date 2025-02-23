package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type ListProductImg struct{}

func (service *ListProductImg) List(ctx context.Context, id string) serializer.Response {
	code := e.Success
	pid, _ := strconv.Atoi(id)
	productImgDao := dao.NewProductImgDao(ctx)
	productImgs, err := productImgDao.ListProductImg(uint(pid))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProductImgVOs(productImgs), uint(len(productImgs)))
}
