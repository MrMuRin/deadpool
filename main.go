package main

import (
	"deadpool/adapters/auth"
	"deadpool/adapters/http"
	"deadpool/adapters/persistence"
	"deadpool/config"
	"deadpool/core/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := "http://localhost:8080/api/auth/google/callback"

	// ตั้งค่า Database
	db := config.InitDB()

	googleAuth := auth.NewGoogleAuth(clientID, clientSecret, redirectURL)

	// ตั้งค่า Repository
	userRepo := persistence.NewUserRepository(db)

	// ตั้งค่า Service
	authService := services.AuthService{UserRepo: userRepo}

	// ตั้งค่า Handlers
	authHandler := http.AuthHandler{AuthService: &authService, GoogleAuth: googleAuth}

	// ตั้งค่า Fiber
	app := fiber.New()

	// Public Routes
	app.Get("/api/auth/google/login", authHandler.GoogleLogin)
	app.Get("/api/auth/google/callback", authHandler.GoogleCallback)

	// Start Server
	log.Println("Starting server on :8080")
	log.Fatal(app.Listen(":8080"))
}
