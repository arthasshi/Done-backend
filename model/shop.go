package model

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name       string  `json:"name"`
	Adress     string  `json:"address"`
	Status     int     `json:"status"`      //0关闭，1开张
	MoneyScale float64 `json:"money_scale"` //积分比例，10元1积分就是10，10/1的意思
}

func GetShopList(page, size int) ([]*Shop, int64, error) {
	var shopes []*Shop
	var total int64
	result := db.Table("shops").Count(&total).Offset((page - 1) * size).Limit(size).Find(&shopes)
	return shopes, total, result.Error
}
func AddShop(shop *Shop) (uint, error) {
	result := db.Create(shop)
	return shop.ID, result.Error
}
func ChangeShopStatus(id uint, status int) error {

	result := db.Model(&Shop{}).Where("id = ?", id).Update("status", status)
	return result.Error
}
func ChangeShopScale(id uint, scale float64) error {
	if id == 0 {
		result := db.Table("shops").Updates(Shop{MoneyScale: scale})
		return result.Error
	} else {
		result := db.Model(&Shop{}).Where("id = ?", id).Update("money_scale", scale)
		return result.Error
	}

}
