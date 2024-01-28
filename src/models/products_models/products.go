package products_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model
	Product_name string
	Price        int
	Color        string
	Size         string
	Stock        int
	Description  string
	Rating       int
	Condition    string
}

func SelectAll_product() *gorm.DB {
	items := []Products{}
	return config.DB.Find(&items)
}

func Select_product(id string) *gorm.DB {
	var item Products
	return config.DB.First(&item, "id = ?", id)
}

func Post_product(item *Products) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_products(id string, newProduct *Products) *gorm.DB {
	var item Products
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Delete_products(id string) *gorm.DB {
	var item Products
	return config.DB.Delete(&item, "id = ?", id)
}
