package main

import (
	"net/http"

	"github.com/steelthedev/go-commerce/connections/db"

	"github.com/steelthedev/go-commerce/package/accounts"
	"github.com/steelthedev/go-commerce/package/products"
	"github.com/steelthedev/go-commerce/package/shops"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {

	// dbURL := string(os.Getenv("DB_URL"))
	dbURL := "postgres://steel:akinwumi1914@localhost:5432/marketplace"
	router := gin.Default()
	dbHandler := db.Init(dbURL)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome to marketplace"})
	})

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
	shops.RegisterRoutes(router, dbHandler)

	router.Run(":8000")
}
