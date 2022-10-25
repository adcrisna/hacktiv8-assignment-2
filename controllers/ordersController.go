package controllers

import (
	"assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (o *OrderController) GetOrders(ctx *gin.Context) {
	orders := []models.Order{}

	if err := ctx.ShouldBind(&orders); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	if err := o.db.Model(&models.Order{}).Preload("Items").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	order := models.Order{}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return
	}

	newOrder := models.Order{CustomerName: order.CustomerName, Items: order.Items}

	if err := o.db.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Insert Data Order Success", "Data ID": newOrder})
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	order := models.Order{}
	item := models.Item{}

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := o.db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := o.db.Unscoped().Where("order_id = ?", order.ID).Delete(item).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := o.db.Save(order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Update Order Success", "data": order})
}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	order := models.Order{}

	if err := o.db.Where("order_id = ?", ctx.Param("orderId")).First(&order).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}
	orderDelete := &order
	if err := o.db.Select(clause.Associations).Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Data Order Success", "data": orderDelete.ID})
}
