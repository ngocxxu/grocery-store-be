package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string         `gorm:"size:255" json:"name"`
	Description   string         `gorm:"type:text" json:"description"`
	Price         float64        `gorm:"type:decimal(10,2)" json:"price"`
	Rating        uint           `json:"rating"`
	Image         pq.StringArray `gorm:"type:text[]" json:"image"`
	WeightOptions []WeightOption `gorm:"foreignKey:ProductID" json:"weight_options"`
	Quantity      uint           `json:"quantity"`
	Sku           string         `gorm:"uniqueIndex;size:50" json:"sku"`
	Status        string         `gorm:"size:20" json:"status"`
	Discount      float64        `gorm:"type:decimal(5,2)" json:"discount"`
	Categories    []Category     `gorm:"many2many:product_categories;" json:"categories"`
}
