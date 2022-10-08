package models

type Item struct {
	ID          uint   `gorm:"primaryKey;column:item_id" json:"-" validate:"required"`
	ItemCode    string `gorm:"not null;type:VARCHAR(50)" form:"itemCode" json:"itemCode" validate:"required" binding:"required"`
	Description string `gorm:"type:TEXT" form:"description" json:"description" binding:"required"`
	Quantity    uint   `gorm:"not null" form:"quantity" json:"quantity" validate:"required" binding:"required"`
	OrderID     uint   `gorm:"not null" json:"-" validate:"required"`
	Orders      Order  `gorm:"foreignKey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
