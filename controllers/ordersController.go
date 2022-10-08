package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetOrders(ctx *gin.Context) {
	db := database.Connection()
	orders := []models.Order{}

	if err := ctx.ShouldBind(&orders); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := db.Model(&models.Order{}).Preload("Items").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func CreateOrder(ctx *gin.Context) {
	db := database.Connection()
	order := models.Order{}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	newOrder := models.Order{CustomerName: order.CustomerName, Items: order.Items}

	if err := db.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Insert Data Order Success", "data": newOrder})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.Connection()
	order := models.Order{}
	item := models.Item{}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := db.Unscoped().Where("order_id = ?", order.ID).Delete(item).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := db.Save(order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Update Order Success", "data": order})
}

func DeleteOrder(ctx *gin.Context) {
	db := database.Connection()
	order := models.Order{}

	if err := db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}
	orderDelete := &order
	if err := db.Select(clause.Associations).Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Data Order Success", "data": orderDelete})
}
