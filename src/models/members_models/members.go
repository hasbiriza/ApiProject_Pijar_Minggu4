package members_models

import (
	"api_tugas_minggu4/src/config"

	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Member_name string
	Email       string
	Password    string
	Role        string
	Address     string
	Phone       int
}

func SelectAll_member() *gorm.DB {
	items := []Member{}
	return config.DB.Find(&items)
}

func Select_member(id string) *gorm.DB {
	var item Member
	return config.DB.First(&item, "id = ?", id)
}

func Post_member(item *Member) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates_member(id string, newProduct *Member) *gorm.DB {
	var item Member
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes_members(id string) *gorm.DB {
	var item Member
	return config.DB.Delete(&item, "id = ?", id)
}
