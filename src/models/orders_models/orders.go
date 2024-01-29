package orders_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	Product_id   string
	Member_id    string
	Delivery_fee int
	Quantity     int
	Payment      string
}

func SelectAll_order() *gorm.DB {
	items := []Orders{}
	return config.DB.Find(&items)
}

func Select_order(id string) *gorm.DB {
	var item Orders
	return config.DB.First(&item, "id = ?", id)
}

func Post_order(item *Orders) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_order(id string, newProduct *Orders) *gorm.DB {
	var item Orders
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes_order(id string) *gorm.DB {
	var item Orders
	return config.DB.Delete(&item, "id = ?", id)
}
