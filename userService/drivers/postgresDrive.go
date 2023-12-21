package drivers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connectdatabase fuction is used to connect to the postgresSQL
func Connectdatabase() *gorm.DB {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error Loading.env", err)
	}

	dsn := fmt.Sprintf("host = %s port=%s user=%s password=%s database=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DATABASE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return db
}
