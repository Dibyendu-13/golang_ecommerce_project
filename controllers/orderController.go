package controllers

import (
	"ecommerce-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	// Example order creation logic
	orderID := 1 // Replace with actual logic to generate ID
	orderDetails := "Sample order"
	services.ProcessOrder(orderID, orderDetails)

	c.JSON(http.StatusOK, gin.H{"message": "Order created and processed"})
}
