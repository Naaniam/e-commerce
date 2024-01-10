package models

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
