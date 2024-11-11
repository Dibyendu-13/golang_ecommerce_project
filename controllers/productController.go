package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	// Logic for creating a product
	c.JSON(http.StatusOK, gin.H{"message": "Product created"})
}

func GetProduct(c *gin.Context) {
	// Logic for getting a product by ID
	c.JSON(http.StatusOK, gin.H{"message": "Product details"})
}
