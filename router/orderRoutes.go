package router

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func RouterOrder() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", controllers.GetOrders)
	router.PUT("/orders/:orderId", controllers.UpdateOrder)
	router.POST("/orders", controllers.CreateOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)

	return router
}
