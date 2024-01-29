package helper

import (
	"api_tugas_minggu4/src/config"
	"api_tugas_minggu4/src/models/category_models"
	"api_tugas_minggu4/src/models/members_models"
	"api_tugas_minggu4/src/models/orders_models"
	"api_tugas_minggu4/src/models/products_models"
	"api_tugas_minggu4/src/models/ratings_models"
)

func Migration() {
	config.DB.AutoMigrate(&members_models.Member{})
	config.DB.AutoMigrate(&category_models.Category{})
	config.DB.AutoMigrate(&products_models.Products{})
	config.DB.AutoMigrate(&ratings_models.Ratings{})
	config.DB.AutoMigrate(&orders_models.Orders{})
}
