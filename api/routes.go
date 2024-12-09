package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	productGroup := router.Group("/api/products")
	{
		productGroup.POST("/", CreateProduct)
		productGroup.GET("/", GetProducts)
		productGroup.PUT("/:id", UpdateProduct)
		productGroup.DELETE("/:id", DeleteProduct)
	}
}
