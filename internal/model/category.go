package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"size:50;uniqueIndex" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Products    []Product `gorm:"many2many:product_categories;" json:"products"`
}
