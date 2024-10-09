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

	for _, unit := range units {
		result := database.Create(&unit)
		if result.Error != nil {
			log.Printf("Failed to seed unit %s: %v", unit.Name, result.Error)
		} else {
			log.Printf("Seeded unit: %s", unit.Name)
		}
	}

	log.Println("Seed data inserted successfully")
}
