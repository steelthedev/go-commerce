package stores

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

	routes := router.Group("stores")

	routes.POST("/create", h.CreateStore)
	routes.GET("/get-all", h.GetAllStores)
	routes.GET("/get-user-store", h.GetUserStore)

}
