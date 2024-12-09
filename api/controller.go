package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Product created"})
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of products"})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
