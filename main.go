package main

import (
	"net/http"
	"os"

	"github.com/steelthedev/go-commerce/connections/db"

	"github.com/steelthedev/go-commerce/package/accounts"
	"github.com/steelthedev/go-commerce/package/products"
	"github.com/steelthedev/go-commerce/package/stores"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	_ "github.com/steelthedev/go-commerce/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title An ecommerce Api
// @version 1.0
// @description A golang microservice for ecommerce

// @host 	localhost:8000
// @BasePath /
func main() {

	dbURL := string(os.Getenv("DB_URL"))

	router := gin.Default()
	dbHandler := db.Init(dbURL)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to marketplace"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(func(c *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		}).ServeHTTP(c.Writer, c.Request, func(w http.ResponseWriter, r *http.Request) {
		})
	})
	router.Static("/images", "../images")
	accounts.RegisterRoutes(router, dbHandler)
	products.RegisterRoutes(router, dbHandler)
	stores.RegisterRoutes(router, dbHandler)
	router.Run(":8000")
}
