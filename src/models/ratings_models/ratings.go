package ratings_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Ratings struct {
	gorm.Model
	Member_id  string
	Product_id string
	Value      int
}

func SelectAll_rating() *gorm.DB {
	items := []Ratings{}
	return config.DB.Find(&items)
}

func Select_rating(id string) *gorm.DB {
	var item Ratings
	return config.DB.First(&item, "id = ?", id)
}

func Post_rating(item *Ratings) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_rating(id string, newProduct *Ratings) *gorm.DB {
	var item Ratings
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Delete_rating(id string) *gorm.DB {
	var item Ratings
	return config.DB.Delete(&item, "id = ?", id)
}
