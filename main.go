package main

import (
	"deadpool/adapters/http"
	"deadpool/adapters/persistence"
	"deadpool/config"
	"deadpool/core/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// ตั้งค่า Database
	db := config.InitDB()

	// ตั้งค่า Repository
	userRepo := persistence.NewUserRepository(db)

	// ตั้งค่า Service
	authService := services.AuthService{UserRepo: userRepo}

	// ตั้งค่า Handlers
	authHandler := http.AuthHandler{AuthService: &authService}

	// ตั้งค่า Fiber
	app := fiber.New()

	// Public Routes
	app.Get("/api/auth/google/login", authHandler.GoogleLogin)
	app.Get("/api/auth/google/callback", authHandler.GoogleCallback)

	// Start Server
	log.Println("Starting server on :8080")
	log.Fatal(app.Listen(":8080"))
}
