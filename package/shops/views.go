package shops

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/connections/tokens"
	"github.com/steelthedev/go-commerce/package/accounts"
	"github.com/steelthedev/go-commerce/package/helpers"
)

// @Summary      Create a new shop using
// @Description  New shop using json request body
// @Tags         shops
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "Shops Name"
// @Success      200  {object} ShopsSerializer
// @Router       /shops/create [post]
func (h handler) CreateShop(c *gin.Context) {

	body := ShopsSerializer{}

	var shop models.Shops
	var err error

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body request",
			"state":   false,
		})
		return
	}

	_, err = accounts.IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to access this request. Requires valid authentication",
			"state":   false,
			"error":   err.Error(),
		})

		return
	}

	var user models.User
	user_id, err := tokens.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You need to be logged in"})
		return
	}

	if err := h.DB.Where("id=?", user_id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "User could not be found",
			"state":   false,
		})
		return
	}

	shop.Name = body.Name
	shop.User = user
	shop.UserId = user_id

	image, err := helpers.AddSingleImage(c, "image")

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
			"state":   false,
			"error":   err.Error(),
		})
	}

	shop.Image = image

	if result := h.DB.Create(&shop); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create shop",
			"state":   false,
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, &shop)

}

// GetAllShops godoc
// @Summary Get all shops
// @Description Retrieve a list of all shops
// @Tags shops
// @Produce json
// @Success 200 {array} ShopsSerializer
// @Failure 404
// @Router /shops/get-all [get]
func (h handler) GetAllShops(c *gin.Context) {

	var shops []models.Shops

	if result := h.DB.Find(&shops); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Could not be found",
			"state":   false,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, &shops)
}
