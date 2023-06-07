package products

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

	routes := router.Group("products")
	routes.POST("/create", h.CreateProduct)
	routes.GET("/get-all", h.GetAllProducts)
	routes.GET("/get/:id", h.GetSingleProduct)
	routes.POST("/category/create", h.CreateCategory)
	routes.GET("category/get/:id", h.GetCategory)
	routes.GET("category/get-all", h.GetAllcategory)
	routes.GET("/get-user-products", h.GetUserProduct)

}
