package customers_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Customer struct {
	gorm.Model
	Customer_name string
	Email         string
	Password      string
}

func SelectAll_customer() *gorm.DB {
	items := []Customer{}
	return config.DB.Find(&items)
}

func Select_customer(id string) *gorm.DB {
	var item Customer
	return config.DB.First(&item, "id = ?", id)
}

func Post_customer(item *Customer) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_customer(id string, newProduct *Customer) *gorm.DB {
	var item Customer
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes_customer(id string) *gorm.DB {
	var item Customer
	return config.DB.Delete(&item, "id = ?", id)
}
