package main

import (
	"log"

	"github.com/ngocxxu/grocery-store-svelte-be/internal/config"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/db"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"gorm.io/gorm"
)

func main() {
	cfg := config.New()

	database, err := db.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	seedData(database)
	log.Println("Seed data inserted successfully")
}

func seedData(database *gorm.DB) {
	units := []model.Unit{
		{Name: "Kilogram", Abbreviation: "Kg"},
		{Name: "Gram", Abbreviation: "g"},
	}

	categories := []model.Category{
		{Name: "Fruits & Vegetables", Description: "The list of Fruits & vegetables"},
		{Name: "Meat & Fish", Description: "The list of Meat & fish"},
	}

	products := []model.Product{
		{
			Name:        "Ground Nuts Oil Pack 43g",
			Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			Sku:         "WH12",
			Status:      "Sale",
			Price:       99.99,
			Discount:    5.00,
			Rating:      4,
			Quantity:    10,
			Image: []string{
				"https://res.cloudinary.com/ngocxxu/image/upload/v1728530367/grocery-store-svelte/back-3_pb7kvm.jpg",
				"https://res.cloudinary.com/ngocxxu/image/upload/v1728530288/grocery-store-svelte/3_ppbbw7.jpg",
			},
			WeightOptions: []model.WeightOption{
				{Weight: 1.5, UnitID: 1},
				{Weight: 2.0, UnitID: 2},
			},
			Categories: []model.Category{
				{Name: "Snacks & Branded Foods"},
			},
		},
		{
			Name:        "Organic Litchi Juice Pack",
			Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit.",
			Sku:         "WH13",
			Status:      "Sale",
			Price:       24.99,
			Discount:    10.00,
			Rating:      3,
			Quantity:    50,
			Image: []string{
				"https://res.cloudinary.com/ngocxxu/image/upload/v1728530367/grocery-store-svelte/back-3_pb7kvm.jpg",
				"https://res.cloudinary.com/ngocxxu/image/upload/v1728530288/grocery-store-svelte/3_ppbbw7.jpg",
			},
			WeightOptions: []model.WeightOption{
				{Weight: 1.5, UnitID: 1},
			},
			Categories: []model.Category{
				{Name: "Fruits & Vegetables"},
			},
		},
	}

	seedUnits(database, units)
	seedProducts(database, products)
	seedCategories(database, categories)
}

func seedUnits(database *gorm.DB, units []model.Unit) {
	for _, unit := range units {
		if result := database.Create(&unit); result.Error != nil {
			log.Printf("Failed to seed unit %s: %v", unit.Name, result.Error)
		} else {
			log.Printf("Seeded unit: %s", unit.Name)
		}
	}
}

func seedProducts(database *gorm.DB, products []model.Product) {
	for _, product := range products {
		if result := database.Create(&product); result.Error != nil {
			log.Printf("Failed to seed product %s: %v", product.Name, result.Error)
		} else {
			log.Printf("Seeded product: %s", product.Name)
		}
	}
}

func seedCategories(database *gorm.DB, categories []model.Category) {
	for _, category := range categories {
		if result := database.Create(&category); result.Error != nil {
			log.Printf("Failed to seed category %s: %v", category.Name, result.Error)
		} else {
			log.Printf("Seeded category: %s", category.Name)
		}
	}
}
