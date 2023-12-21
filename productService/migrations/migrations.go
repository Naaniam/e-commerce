package migrations

import (
	"github.com/e-commerce/product/models"
	"gorm.io/gorm"
)

func Migrators(db *gorm.DB) {
	db.AutoMigrate(&models.Brand{}, &models.RAMSize{})
}
