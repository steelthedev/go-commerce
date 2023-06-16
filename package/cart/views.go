package cart

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/connections/tokens"
	"github.com/steelthedev/go-commerce/package/accounts"
	"gorm.io/gorm"
)

// GetUserStores godoc
// @Summary Add to Cart
// @Description Add a product to cart
// @Tags Cart
// @Produce json
// @Success 200 {object} models.Cart
// @Param Authorization header string true "Token for Authorization"
// @Param body body AddProduct true "Token for Authorization"
// @Failure 404
// @Failure 401
// @Failure 500
// @Success 200
// @Router /cart/add-to-cart [post]
func (h handler) AddToCart(c *gin.Context) {

	//check if user is authenticated

	_, err := accounts.IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "User not authenticated",
			"state":   false,
		})
		return
	}

	body := AddProduct{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Invalid request body",
			"state":   false,
		})
		return
	}

	userId, err := tokens.ExtractTokenID(c)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Could not fetch ID",
			"state":   false,
		})
		return
	}

	//fetch user cart

	var userCart models.Cart

	if err := h.DB.Where("user_id=?", userId).Preload("Products").First(&userCart).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Cart not found, create a new cart for the user
			userCart.UserID = userId

		} else {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Could not fetch cart",
				"state":   false,
			})
			return
		}

	}

	//fetch the product to add to cart
	var product models.Product

	if result := h.DB.Preload("Category").Preload("Store").Preload("Store.User").First(&product, body.ProductID); result.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "Could not fetch product",
			"state":   false,
		})
		return
	}

	//check if the product is in cart; increase quantity if in cart

	for _, item := range userCart.Products {
		if item.ProductID == body.ProductID {

			item.Quantity += body.Quantity

			if result := h.DB.Save(&item); result.Error != nil {

				c.AbortWithStatusJSON(500, gin.H{
					"message": "Could not save cart item",
					"state":   false,
				})
				return
			}
			c.IndentedJSON(200, gin.H{
				"message": "Successfully updated quantity",
				"state":   true,
			})
			return
		}

	}

	//create cartItem

	cartItem := models.CartItem{
		Product:   product,
		UserID:    userId,
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
	}

	//add the cartItem to cart

	userCart.Products = append(userCart.Products, cartItem)

	if result := h.DB.Save(&userCart); result.Error != nil {

		c.AbortWithStatusJSON(500, gin.H{
			"message": "Could not save Cart",
			"state":   false,
			"error":   result.Error,
		})
		return
	}

	cartData := CartSerializer{
		Products: userCart.Products,
		UserID:   userCart.UserID,
		ID:       userCart.ID,
	}

	c.IndentedJSON(200, &cartData)

}

func (h handler) GetUserCart(c *gin.Context) {
	_, err := accounts.IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "User not authenticated",
			"state":   false,
		})
		return
	}

	userId, err := tokens.ExtractTokenID(c)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Could not fetch ID",
			"state":   false,
		})
		return
	}

	//fetch user cart

	var userCart models.Cart

	if err := h.DB.Where("user_id=?", userId).Preload("User").Preload("Products.Product").First(&userCart).Error; err != nil {

		c.AbortWithStatusJSON(500, gin.H{
			"message": "Could not fetch cart",
			"state":   false,
			"error":   err.Error(),
		})
		return
	}

	c.IndentedJSON(200, &userCart)
}
