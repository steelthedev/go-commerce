package models

import (
	"github.com/lib/pq"
)

type User struct {
	ID           uint   `gorm:"primary_key;index:idx_name,unique"`
	FirstName    string `json:"first_name" validate:"required,min=2,max=40"`
	LastName     string `json:"last_name" validate:"required,min=2,max=40"`
	Password     string `json:"password" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone" validate:"required"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type Product struct {
	ID          uint           `gorm:"primary_key;index:idx_name,unique"`
	Title       string         `json:"title" validate:"required, min=2 max=40"`
	Description string         `json:"description" validate:"required"`
	Category    []Categories   `json:"product_category" gorm:"many2many:product_categories;"`
	Price       float64        `json:"price"`
	MainImage   string         `json:"main_image"`
	SubImages   pq.StringArray `json:"sub_images" gorm:"type:text[]"`
	Shop        Shops          `json:"shop"`
	ShopID      uint           `json:"shop_id"`
}

type Order struct {
	ID     uint        `gorm:"primary_key;index:idx_name,unique"`
	User   User        `json:"user"`
	UserID uint        `json:"user_id"`
	Items  []OrderItem `json:"order_items" gorm:"many2many:order_items;"`
	Price  float64     `json:"price"`
}

type OrderItem struct {
	ID        uint    `gorm:"primary_key;index:idx_name,unique"`
	User      User    `json:"user"`
	UserID    uint    `json:"user_id"`
	Order     Order   `json:"order"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  uint    `json:"quantity"`
}

type Categories struct {
	ID    uint   `gorm:"primary_key;index:idx_name,unique"`
	Title string `json:"title"`
}

type Shops struct {
	ID     uint   `gorm:"primary_key;index:idx_name,unique"`
	Name   string `json:"shop_name"`
	User   User   `json:"owner"`
	UserId uint   `json:"user_id"`
	Image  string `json:"image"`
}
