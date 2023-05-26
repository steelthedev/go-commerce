package accounts

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/accounts")
	routes.POST("/signup", h.SignUp)
	routes.POST("/login", h.Login)
	routes.GET("/profile", h.GetUser)
}
