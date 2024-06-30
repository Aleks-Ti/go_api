package db

import (
	"api_fiber/src/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	fmt.Println("Crreating structure")
	godotenv.Load()
	db_host := os.Getenv("PG_HOST")
	db_username := os.Getenv("PG_USERNAME")
	db_password := os.Getenv("PG_PASSWORD")
	db_name := os.Getenv("PG_NAME_DB")

	// connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow", db_host, db_username, db_password, db_name)
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("db connection error")
	}

	DB = db
	fmt.Println("db connected successfully")

	AutoMigrate(db)
}

func AutoMigrate(dsn *gorm.DB) {
	dsn.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)
}
