package serializer

import "gin_mall/model"

type AddressVO struct {
	Id        uint   `json:"id"`
	CreatedAt int64  `json:"created_at"`
	UserId    uint   `json:"user_id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

func BuildAddress(item *model.Address) *AddressVO {
	return &AddressVO{
		Id:        item.ID,
		CreatedAt: item.CreatedAt.Unix(),
		UserId:    item.UserID,
		Name:      item.Name,
		Address:   item.Address,
		Phone:     item.Phone,
	}
}

func BuildAddresses(items []*model.Address) []*AddressVO {
	var addresses []*AddressVO
	for _, item := range items {
		addresses = append(addresses, BuildAddress(item))
	}
	return addresses
}
