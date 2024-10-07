package model

import (
	"gorm.io/gorm"
)



type Unit struct {
	gorm.Model
	Name        string `gorm:"size:50;uniqueIndex" json:"name"`
	Abbreviation string `gorm:"size:10;uniqueIndex" json:"abbreviation"`
}