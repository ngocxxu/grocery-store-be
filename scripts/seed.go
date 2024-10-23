package main

import (
	"log"

	"github.com/ngocxxu/grocery-store-svelte-be/internal/config"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/db"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
)

func main() {
	cfg := config.New()

	database, err := db.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	units := []model.Unit{
		{Name: "Kilogram", Abbreviation: "Kg"},
		{Name: "Gram", Abbreviation: "g"},
	}

	categories := []model.Category{
		{Name: "Fruits & vegetables", Description: "The list of Fruits & vegetables"},
		{Name: "Meat & fish", Description: "The list of Meat & fish"},
	}

	for _, unit := range units {
		result := database.Create(&unit)
		if result.Error != nil {
			log.Printf("Failed to seed unit %s: %v", unit.Name, result.Error)
		} else {
			log.Printf("Seeded unit: %s", unit.Name)
		}
	}

	for _, category := range categories {
		result := database.Create(&category)
		if result.Error != nil {
			log.Printf("Failed to seed category %s: %v", category.Name, result.Error)
		} else {
			log.Printf("Seeded category: %s", category.Name)
		}
	}

	log.Println("Seed data inserted successfully")
}
