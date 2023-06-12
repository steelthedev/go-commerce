package stores

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/connections/tokens"
	"github.com/steelthedev/go-commerce/package/accounts"
	"github.com/steelthedev/go-commerce/package/helpers"
)

// @Summary      Create a new store using
// @Description  New store using json request body
// @Tags         Stores
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "Stores Name"
// @Success      200  {object} StoresSerializer
// @Router       /stores/create [post]
func (h handler) CreateStore(c *gin.Context) {

	var store models.Stores
	var err error

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

	//check if user already has a store
	if err := h.DB.Where("user_id=?", user.ID).First(&store).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User already has a store",
			"state":   false,
		})
		return
	}

	store.Name = c.Request.FormValue("store_name")
	store.User = user
	store.UserId = user_id

	image, err := helpers.AddSingleImage(c, "image")

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Bad request",
			"state":   false,
			"error":   err.Error(),
		})
	}

	store.Image = image

	if result := h.DB.Create(&store); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create store",
			"state":   false,
		})
		return
	}
	c.IndentedJSON(http.StatusCreated, &store)

}

// GetAllStores godoc
// @Summary Get all Stores
// @Description Retrieve a list of all Stores
// @Tags Stores
// @Produce json
// @Success 200 {array} StoresSerializer
// @Failure 404
// @Router /stores/get-all [get]
func (h handler) GetAllStores(c *gin.Context) {

	var Stores []models.Stores

	if result := h.DB.Preload("User").Find(&Stores); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Could not be found",
			"state":   false,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, &Stores)
}

// GetUserStores godoc
// @Summary Get a user store
// @Description Retrieve a the store belonging to a logged in user
// @Tags Stores
// @Produce json
// @Success 200 {array} StoresSerializer
// @Headers
// @Failure 404
// @Failure 401
// @Router /stores/get-user-store [get]
func (h handler) GetUserStore(c *gin.Context) {
	var store models.Stores
	var user models.User

	_, err := accounts.IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to access this request. Requires valid authentication",
			"state":   false,
			"error":   err.Error(),
		})
		return
	}

	//get user from user id

	user_id, err := tokens.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User Id does not exist",
			"state":   false,
			"error":   err.Error(),
		})
		return
	}

	//check user exists
	if result := h.DB.Where("ID=?", user_id).First(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "User could not be fetched",
			"state":   false,
			"error":   err.Error(),
		})
		return
	}

	//get user store

	if result := h.DB.Where("user_id=?", user.ID).First(&store); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Store with user could not be found",
			"state":   false,
			"error":   err.Error(),
		})
		return
	}

	body := StoresSerializer{
		ID:    store.ID,
		Name:  store.Name,
		Image: store.Image,
		User: accounts.UserSerializer{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		},
	}

	c.IndentedJSON(200, &body)

}
