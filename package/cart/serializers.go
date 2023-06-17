package cart

import "github.com/steelthedev/go-commerce/connections/models"

type AddProduct struct {
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type CartSerializer struct {
	ID       uint              `json:"id"`
	UserID   uint              `json:"user_id"`
	Products []models.CartItem `json:"products"`
	Price    float64           `json:"price"`
}

type ProductCartSerializer struct {
	ID          uint                `gorm:"primary_key;index:idx_name,unique"`
	Title       string              `json:"title" validate:"required, min=2 max=40"`
	Category    []models.Categories `json:"product_category"`
	Price       uint64              `json:"price"`
	MainImage   string              `json:"main_image"`
	SubImages   []string            `json:"sub_images"`
	Description string              `json:"description"`
}
