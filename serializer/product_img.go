package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

type ProductImgVO struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImgVO(item *model.ProductImg) ProductImgVO {
	return ProductImgVO{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgVOs(items []*model.ProductImg) (productImgs []ProductImgVO) {
	for _, item := range items {
		productImgs = append(productImgs, BuildProductImgVO(item))
	}
	return
}
