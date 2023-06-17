package orders

import (
	"github.com/gin-gonic/gin"
	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/connections/tokens"
	"github.com/steelthedev/go-commerce/package/accounts"
)

func (h handler) CreateOrder(c *gin.Context) {
	_, err := accounts.IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unathorized user",
			"state":   false,
		})
	}

	userId, err := tokens.ExtractTokenID(c)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Could not fetch ID",
			"state":   false,
		})
		return
	}

	//get userCart
	var userCart models.Cart

	if err := h.DB.Where("user_id=?", userId).Preload("User").Preload("Products.Product").First(&userCart).Error; err != nil {

		c.AbortWithStatusJSON(500, gin.H{
			"message": "Could not fetch cart",
			"state":   false,
		})
		return
	}

	// get or create orders
	order := models.Order{
		UserID:    userId,
		Items:     userCart.Products,
		Price:     userCart.Price,
		Delivered: false,
		Paid:      false,
	}

	if result := h.DB.Where(&order).FirstOrCreate(&order); result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Could not create order",
			"state":   false,
		})
		return
	}

	orderData := OrderSerializer{
		ID:        order.ID,
		Paid:      order.Paid,
		Delivered: order.Delivered,
		UserID:    order.UserID,
		Items:     order.Items,
	}
	c.IndentedJSON(200, &orderData)

}
