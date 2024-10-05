package main

import (
	"log"
	"net/http"

	"github.com/ngocxxu/grocery-store-svelte-be/internal/config"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/db"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/handler"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/model"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/repository"
	"github.com/ngocxxu/grocery-store-svelte-be/internal/service"
)

func main() {
    cfg := config.New()

    database, err := db.New(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto Migrate
    err = database.AutoMigrate(&model.User{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    userRepo := repository.NewUserRepository(database)
    userService := service.NewUserService(userRepo)

    http.Handle("/", handler.NewPlaygroundHandler())
    http.Handle("/query", handler.NewGraphQLHandler(userService))

		log.Printf("Connect to http://localhost:%s/ for GraphQL playground", cfg.Port)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}