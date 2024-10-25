package main

import (
	"log"
	"net/http"

	"github.com/ngocxxu/grocery-store-svelte-be/internal/config"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/db"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/handler"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/repository"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/service"
	"github.com/ngocxxu/grocery-store-svelte-be/migrations"
)

func main() {
	cfg := config.New()
	database, err := db.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := migrations.RunMigrations(database); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)

	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)

	http.Handle("/", handler.NewPlaygroundHandler())
	http.Handle("/query", handler.NewGraphQLHandler(userService, productService))

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
