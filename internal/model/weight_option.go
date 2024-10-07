package model

import (
	"gorm.io/gorm"
)



type WeightOption struct {
	gorm.Model
	ProductID uint    `json:"product_id"` // Foreign key
	Weight    float64 `gorm:"type:decimal(10,3)" json:"weight"`
	UnitID    uint    `json:"unit_id"` // Foreign key
	Unit      Unit    `gorm:"foreignKey:UnitID" json:"unit"` // Relationship to Unit
}

