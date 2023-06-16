package cart

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

	routes := router.Group("/cart")
	routes.POST("/add-to-cart", h.AddToCart)
	routes.GET("/get-user-cart", h.GetUserCart)
}
