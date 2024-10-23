package model

import (
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ProductID  uint `gorm:"not null;index" json:"product_id"`
	CategoryID uint `gorm:"not null;index" json:"category_id"`
}
