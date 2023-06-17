package orders

import "github.com/steelthedev/go-commerce/connections/models"

type OrderSerializer struct {
	ID        uint              `json:"id"`
	UserID    uint              `json:"user_id"`
	Items     []models.CartItem `json:"order_items"`
	Price     float64           `json:"price"`
	Paid      bool              `json:"paid"`
	Delivered bool              `json:"delivered"`
}
