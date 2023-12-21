package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type OrderResponse struct {
	gorm.Model
	OrderID      uuid.UUID `gorm:"column:order_id"`
	IsDVDInclude bool      `json:"is_dvd_include" gorm:"column:is_dvd_include"`
	DVDPrice     float64   `json:"dvd_price" gorm:"column:dvd_price"`
	OrderStatus  string    `json:"order_status" gorm:"column:order_status"`
	MemberEmail  string    `json:"member_email" gorm:"column:member_email"`
	BrandName    string    `json:"brand_name" gorm:"column:brand_name"`
	BrandPrice   float64   `json:"brand_price" gorm:"column:brand_price"`
	RAMPrice     float64   `json:"ram_price" gorm:"column:ram_price"`
	NetPrice     float64   `json:"net_price" gorm:"column:net_price"`
}

type DeletedAt sql.NullTime

// Scan implements the Scanner interface.
func (n *DeletedAt) Scan(value interface{}) error {
	return (*sql.NullTime)(n).Scan(value)
}

type PlaceOrder struct {
	OrderID      uuid.UUID      `gorm:"type:uuid;primaryKey;column:order_id"`
	MemberMail   string         `json:"member_email" gorm:"column:member_email" validate:"required,email"`
	BrandName    string         `json:"brand_name" gorm:"column:brand_name" validate:"required"`
	RAMSizeINGB  int            `json:"ram_size_in_gb" gorm:"column:ram_size_in_gb" validate:"required"`
	IsDVDInclude bool           `json:"is_dvd_include" gorm:"column:is_dvd_include"`
	CreatedAt    time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
