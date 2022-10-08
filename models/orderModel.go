package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey;column:order_id" json:"-" validate:"required"`
	OrderedAt    time.Time `gorm:"not null;autoCreateTime" json:"orderedAt" validate:"required"`
	CustomerName string    `gorm:"type:VARCHAR(50)" form:"customerName" json:"customerName" validate:"required" binding:"required"`
	Items        []Item    `form:"items" json:"items" validate:"required" binding:"required"`
}
