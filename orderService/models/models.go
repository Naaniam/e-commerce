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

// Member represents the member table
type Member struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id"`
	FirstName string    `json:"first_name" gorm:"column:first_name" validate:"required"`
	Lastname  string    `json:"last_name" gorm:"column:last_name" validate:"required"`
	Email     string    `json:"email" gorm:"unique;column:email" validate:"required"`
	Phone     string    `json:"phone" gorm:"unique;column:phone" validate:"required"`
	Password  string    `json:"password" gorm:"column:password" validate:"required"`
	Addresses []Address `json:"address" gorm:"foreignKey:MemberID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Address represents the addresses table
type Address struct {
	ID           int    `json:"id" gorm:"primaryKey;column:id"`
	MemberID     string `json:"member_id" gorm:"column:member_id"`
	AddressLine1 string `json:"address_line1" gorm:"column:address-line1" validate:"required"`
	AddressLine2 string `json:"address_line2" gorm:"column:address-line2" validate:"required"`
	City         string `json:"city" gorm:"column:city" validate:"required"`
	Zip          string `json:"zip" gorm:"column:zip" validate:"required"`
	State        string `json:"state" gorm:"column:state" validate:"required"`
	Country      string `json:"country" gorm:"column:country" validate:"required"`
}

// Admin represents the admins table
type Admin struct {
	ID       int    `json:"id" gorm:"primaryKey;column:id"`
	Name     string `json:"name" gorm:"column:name" validate:"required"`
	Email    string `json:"email" gorm:"unique;column:email" validate:"required"`
	Password string `json:"password" gorm:"column:password" validate:"required"`
}
