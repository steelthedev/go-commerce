package accounts

import (
	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/connections/tokens"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyPassword(password, hashedpwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpwd), []byte(password))
}

func LoginCheck(db *gorm.DB, email string, password string) (string, error) {
	var err error

	user := models.User{}

	err = db.Where("Email=?", email).First(&user).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {

		return "", err
	}

	token, err := tokens.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func IsAuthenticated(c *gin.Context) (bool, error) {

	var err error

	// check if token is valid

	err = tokens.TokenValid(c)

	if err != nil {
		return false, err
	}

	return true, nil

}
