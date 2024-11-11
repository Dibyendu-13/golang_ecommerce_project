package routes

import (
	"ecommerce-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	// Product routes
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products/:id", controllers.GetProduct)

	// Order routes
	router.POST("/orders", controllers.CreateOrder)
}
