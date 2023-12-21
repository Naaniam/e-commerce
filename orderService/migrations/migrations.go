package migrations

import (
	"github.com/e-commerce/order/models"
	"gorm.io/gorm"
)

func Migrators(db *gorm.DB) {
	db.AutoMigrate(&models.PlaceOrder{}, &models.OrderResponse{}, models.OrderResponse{})
}
