package models

import "github.com/lib/pq"

type User struct {
	ID        uint   `gorm:"primary_key;index:idx_name,unique"`
	FirstName string `json:"first_name" validate:"required,min=2,max=40"`
	LastName  string `json:"last_name" validate:"required,min=2,max=40"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
}

type Product struct {
	ID          uint           `gorm:"primary_key;index:idx_name,unique"`
	Title       string         `json:"title" validate:"required, min=2 max=40"`
	Description string         `json:"description" validate:"required"`
	Category    []Categories   `json:"product_category" gorm:"many2many:product_categories;"`
	Price       float64        `json:"price"`
	MainImage   string         `json:"main_image"`
	SubImages   pq.StringArray `json:"sub_images" gorm:"type:text[]"`
	Store       Stores         `json:"store"`
	StoreID     uint           `json:"store_id"`
}

type Order struct {
	ID        uint       `gorm:"primary_key;index:idx_name,unique"`
	User      User       `json:"user"`
	UserID    uint       `json:"user_id"`
	Items     []CartItem `json:"order_items" gorm:"many2many:order_items;"`
	Price     float64    `json:"price"`
	Paid      bool       `json:"paid"`
	Delivered bool       `json:"delivered"`
	Payment   Payment    `json:"payment"`
	PaymentID uint       `json:"payment_id"`
}

type CartItem struct {
	ID        uint    `gorm:"primary_key;index:idx_name,unique"`
	UserID    uint    `json:"user_id"`
	ProductID uint    `json:"product_id" gorm:"index;default:2"`
	Product   Product `json:"product"`
	Quantity  uint    `json:"quantity"`
}

type Cart struct {
	ID       uint       `gorm:"primary_key;index:idx_name,unique"`
	User     User       `json:"user"`
	UserID   uint       `json:"user_id"`
	Products []CartItem `json:"products" gorm:"many2many:product_carts;"`
	Price    float64    `json:"price"`
}

type Categories struct {
	ID    uint   `gorm:"primary_key;index:idx_name,unique"`
	Title string `json:"title"`
}

type Stores struct {
	ID     uint   `gorm:"primary_key;index:idx_name,unique"`
	Name   string `json:"store_name"`
	User   User   `json:"owner"`
	UserId uint   `json:"user_id"`
	Image  string `json:"image"`
}

type PaymentType string

const (
	PaymentTypeCard   PaymentType = "Card"
	PaymentTypePaypal PaymentType = "Paypal"
)

type Payment struct {
	ID     uint        `gorm:"primary_key;index:idx_name,unique"`
	Type   PaymentType `json:"type"`
	Amount float64     `json:"amount"`
}
