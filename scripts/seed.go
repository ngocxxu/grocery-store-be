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

    users := []model.User{
        {Name: "John Doe", Email: "john@example.com"},
        {Name: "Jane Smith", Email: "jane@example.com"},
    }

    for _, user := range users {
        result := database.Create(&user)
        if result.Error != nil {
            log.Printf("Failed to seed user %s: %v", user.Name, result.Error)
        } else {
            log.Printf("Seeded user: %s", user.Name)
        }
    }

    log.Println("Seed data inserted successfully")
}