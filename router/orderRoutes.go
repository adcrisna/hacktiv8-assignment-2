package router

import (
	"assignment-2/controllers"
	"assignment-2/database"

	"github.com/gin-gonic/gin"
)

func RouterOrder() *gin.Engine {
	router := gin.Default()
	db := database.Connection()

	orderController := controllers.NewOrderController(db)

	router.GET("/orders", orderController.GetOrders)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.POST("/orders", orderController.CreateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)

	return router
}
