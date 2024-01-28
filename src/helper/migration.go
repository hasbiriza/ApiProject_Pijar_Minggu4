package helper

import (
	"api_tugas_minggu4/src/config"
	"api_tugas_minggu4/src/models/category_models"
	"api_tugas_minggu4/src/models/customers_models"
	"api_tugas_minggu4/src/models/products_models"
)

func Migration() {
	config.DB.AutoMigrate(&customers_models.Customer{})
	config.DB.AutoMigrate(&category_models.Category{})
	config.DB.AutoMigrate(&products_models.Products{})
}
