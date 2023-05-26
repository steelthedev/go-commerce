package shops

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

	routes := router.Group("shops")

	routes.POST("/create", h.CreateShop)
	routes.GET("/get-all", h.GetAllShops)
}
