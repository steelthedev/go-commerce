package cart

import (
	"log"

	"github.com/steelthedev/go-commerce/connections/models"
)

func CalculateCartTotal(cart *models.Cart) float64 {
	var total float64

	for _, item := range cart.Products {
		subtotal := float64(item.Quantity) * item.Product.Price
		log.Fatal(subtotal)
		total += subtotal
	}

	return total
}
