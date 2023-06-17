package cart

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/package/helpers"
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

func (h handler) GetCart(c *gin.Context) (models.Cart, error) {

	userId, err := helpers.GetUserId(c)

	//get cart
	var cart models.Cart

	if result := h.DB.Find(userId, cart); result.Error != nil {
		return cart, err
	}

	return cart, nil
}
