package category_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Category_name string
}

func SelectAll_category() *gorm.DB {
	items := []Category{}
	return config.DB.Find(&items)
}

func Select_category(id string) *gorm.DB {
	var item Category
	return config.DB.First(&item, "id = ?", id)
}

func Post_category(item *Category) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_category(id string, newProduct *Category) *gorm.DB {
	var item Category
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes_category(id string) *gorm.DB {
	var item Category
	return config.DB.Delete(&item, "id = ?", id)
}
