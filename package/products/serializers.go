package products

import "github.com/steelthedev/go-commerce/connections/models"

type CreateProduct struct {
	ID          uint                `gorm:"primary_key;index:idx_name,unique"`
	Title       string              `json:"title" validate:"required, min=2 max=40"`
	CategoryID  uint64              `json:"product_category"`
	Price       float64             `json:"price"`
	MainImage   string              `json:"main_image"`
	SubImages   []string            `json:"sub_images"`
	Shop        models.Stores       `json:"store"`
	Category    []models.Categories `json:"category"`
	Description string              `json:"description"`
}

type ProductSerializer struct {
	ID          uint                `gorm:"primary_key;index:idx_name,unique"`
	Title       string              `json:"title" validate:"required, min=2 max=40"`
	Category    []models.Categories `json:"product_category"`
	Price       uint64              `json:"price"`
	MainImage   string              `json:"main_image"`
	SubImages   []string            `json:"sub_images"`
	Shop        models.Stores       `json:"store"`
	Description string              `json:"description"`
}

type CategoriesSerializer struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
