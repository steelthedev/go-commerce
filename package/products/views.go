package products

import (
	"net/http"
	"strconv"

	"github.com/steelthedev/go-commerce/connections/models"
	"github.com/steelthedev/go-commerce/connections/tokens"
	"github.com/steelthedev/go-commerce/package/accounts"
	"github.com/steelthedev/go-commerce/package/helpers"

	"github.com/gin-gonic/gin"
)

func (h handler) GetCategory(c *gin.Context) {

	category_id, _ := c.Params.Get("id")

	var category models.Categories

	if category_id != "" {
		if result := h.DB.Where("ID=?", category_id).First(&category); result.Error == nil {
			c.IndentedJSON(http.StatusOK, gin.H{

				"data": category,
			})

			return

		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "could not be found, empty feedback",
		"state":   false,
	})

}

func (h handler) CreateCategory(c *gin.Context) {
	body := CategoriesSerializer{}

	var category models.Categories

	var err error

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid body request",
			"state":   false,
		})
		return

	}

	_, err = accounts.IsAuthenticated(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to access this request. Requires authentication",
			"state":   false,
			"error":   err.Error(),
		})

		return
	}

	category.Title = body.Title

	if result := h.DB.Create(&category); result.Error == nil {
		c.IndentedJSON(http.StatusCreated, gin.H{
			"message": "Created successfully",
			"state":   true,
		})
	}
}

func (h handler) GetAllcategory(c *gin.Context) {
	var categories []models.Categories

	if err := h.DB.Find(&categories).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Not found",
			"state":   false,
		})

		return
	}

	c.IndentedJSON(http.StatusOK, &categories)
}

func (h handler) CreateProduct(c *gin.Context) {

	var product models.Product
	var category models.Categories
	var shop models.Shops
	var err error
	var user models.User

	_, err = accounts.IsAuthenticated(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to access this request. Requires authentication",
			"state":   false,
			"error":   err.Error(),
		})

		return
	}

	err = c.Request.ParseMultipartForm(32 << 20)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"state":   "false",
			"error":   err.Error(),
		})
		return
	}

	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	category_id, err := strconv.ParseUint(c.Request.FormValue(("product_category")), 8, 64)
	price, err := strconv.ParseFloat(c.Request.FormValue("price"), 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state": "false",
			"error": err.Error(),
		})
		return
	}

	mainImagepath, err := helpers.AddSingleImage(c, "mainImage")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state": "false",
			"error": err.Error(),
		})
		return
	}

	subImages := c.Request.MultipartForm.File["subImages"]

	subImagesPath, err := helpers.AddMultipleImage(c, subImages)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state": "false",
			"error": err.Error(),
		})
		return
	}

	user_id, err := tokens.ExtractTokenID(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state": "false",
			"error": err.Error(),
		})

		return
	}

	// get user

	category_ := make([]models.Categories, 0)

	category_ = append(category_, category)

	body := models.Product{
		Title:       title,
		Price:       price,
		MainImage:   mainImagepath,
		SubImages:   subImagesPath,
		Description: description,
	}
	if err := h.DB.First(&user, user_id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state": "false",
			"error": err.Error(),
		})
		return
	}

	if err := h.DB.Where("ID=?", category_id).First(&category).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
			"state":   false,
		})

		return

	}

	if err := h.DB.Where("user_id=?", user.ID).First(&shop).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Shop not found",
			"state":   false,
		})

		return
	}

	body.Category = append(category_, category)
	body.Shop = shop

	if result := h.DB.Create(&body); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
			"state":   false,
		})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Product successfully created",
		"state":   true,
		"product": product,
	})

}

func (h handler) GetAllProducts(c *gin.Context) {
	var products []models.Product

	if err := h.DB.Preload("Category").Preload("Shop").Preload(("Shop.User")).Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Not found",
			"state":   false,
		})

		return
	}

	c.IndentedJSON(http.StatusOK, &products)
}

func (h handler) GetSingleProduct(c *gin.Context) {
	product_id, _ := c.Params.Get("id")

	var product models.Product

	if product_id != "" {
		if result := h.DB.Where("ID=?", product_id).First(&product); result.Error == nil {
			c.IndentedJSON(http.StatusOK, &product)
			return
		}
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "could not be found, empty feedback",
		"state":   false,
	})
}
