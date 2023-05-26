package db

import (
	"log"

	"github.com/steelthedev/go-commerce/connections/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {

		log.Fatalln(err)

	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Shops{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItem{})
	// db.AutoMigrate(&models.Address{})
	// db.AutoMigrate(&models.Payment{})

	return db

}
