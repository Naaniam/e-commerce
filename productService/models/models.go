package models

import "gorm.io/gorm"

// Brand represents the brand table
type Brand struct {
	gorm.Model
	BrandName  string    `json:"brand_name" gorm:"unique;column:brand_name" validate:"required"`
	BrandPrice float64   `json:"brand_price" gorm:"column:brand_price" validate:"required"`
	RAMSize    []RAMSize `json:"ram_size" gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// RAMSize represents the ram_size table
type RAMSize struct {
	gorm.Model
	BrandID  uint    `json:"brand_id" gorm:"column:brand_id"`
	SizeInGB int     `json:"size_in_gb" gorm:"column:size_in_gb" validate:"required"`
	RamPrice float64 `json:"ram_price" gorm:"column:ram_price" validate:"required"`
}
