package migrations

import (
	"github.com/e-commerce/user/models"
	"gorm.io/gorm"
)

func Migrators(db *gorm.DB) {
	db.AutoMigrate(&models.Admin{}, models.Member{}, models.Address{})
}
