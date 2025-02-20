package serializer

import "gin_mall/model"

type Carousel struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"image_path"`
	ProductId uint   `json:"product"`
	CreatedAt int64  `json:"created_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildCarousels(item []model.Carousel) (carousels []Carousel) {
	for _, item := range item {
		carousels = append(carousels, BuildCarousel(&item))
	}
	return
}
