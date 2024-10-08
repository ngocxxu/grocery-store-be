// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Mutation struct {
}

type Product struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Type          string          `json:"type"`
	Sku           string          `json:"sku"`
	Status        string          `json:"status"`
	Price         float64         `json:"price"`
	Discount      float64         `json:"discount"`
	Rating        int             `json:"rating"`
	Quantity      int             `json:"quantity"`
	WeightOptions []*WeightOption `json:"weightOptions"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
}

type ProductInput struct {
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Type          string               `json:"type"`
	Sku           string               `json:"sku"`
	Status        string               `json:"status"`
	Price         float64              `json:"price"`
	Discount      float64              `json:"discount"`
	Rating        int                  `json:"rating"`
	Quantity      int                  `json:"quantity"`
	WeightOptions []*WeightOptionInput `json:"weightOptions"`
}

type Query struct {
}

type Unit struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type WeightOption struct {
	ID     string  `json:"id"`
	Weight float64 `json:"weight"`
	Unit   *Unit   `json:"unit"`
}

type WeightOptionInput struct {
	Weight float64 `json:"weight"`
	UnitID string  `json:"unitId"`
}
