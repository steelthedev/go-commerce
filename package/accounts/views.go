package accounts

import (
	"net/http"

	"github.com/steelthedev/go-commerce/connections/tokens"

	"github.com/steelthedev/go-commerce/connections/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Create Users
// @Description Endpoint for creating all categories of users
// @Tags Accounts
// @Accept json
// @Produce json
// @Param body body CreateUser true "SignUp Body Payload"
// @Success 200 {object} UserSerializer
// @Failure 400
// @Failure 302
// @Success 201
// @Router       /accounts/signup [post]
func (h handler) SignUp(c *gin.Context) {

	body := CreateUser{}

	var user models.User

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "There is an error in your inputs, please try again"})
		return
	}

	if result := h.DB.Where("Email=?", body.Email).First(&user); result.Error == nil {
		c.AbortWithStatusJSON(http.StatusFound, gin.H{"message": "User already exists."})
		return
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	hashedpwd, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	user.Password = string(hashedpwd)
	user.Phone = body.Phone
	user.Email = body.Email

	if result := h.DB.Create(&user); result.Error == nil {
		user := UserSerializer{
			Email:     user.Email,
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"data": user})
	}
}

// @Summary Login
// @Description Endpoint for log in
// @Tags Accounts
// @Accept json
// @Produce json
// @Param body body LoginSerializer true "Login Body Payload"
// @Success 200 {object} UserSerializer
// @Failure 400
// @Success 200
// @Router       /accounts/login [post]
func (h handler) Login(c *gin.Context) {

	body := LoginSerializer{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body request",
			"state":   false,
		})
		return
	}

	user := models.User{}

	user.Email = body.Email
	user.Password = body.Password

	token, err := LoginCheck(h.DB, user.Email, user.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid details",
			"state":   false,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Get User Profile
// @Description Endpoint for getting a user based on id
// @Tags Accounts
// @Produce json
// @Param Authorization header string true "Token for Authorization"
// @Success 200 {object} UserSerializer
// @Failure 400
// @Router       /accounts/profile [get]
func (h handler) GetUser(c *gin.Context) {

	var user models.User

	_, err := IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to access this request. Requires valid authentication",
			"state":   false,
			"error":   err.Error(),
		})

		return
	}

	user_id, err := tokens.ExtractTokenID(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"state":   false,
		})
		return
	}

	err = h.DB.Where("ID=?", user_id).First(&user).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User could not be found",
			"state":   false,
		})
		return
	}

	response := UserSerializer{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
	}

	c.IndentedJSON(http.StatusOK, &response)

}
