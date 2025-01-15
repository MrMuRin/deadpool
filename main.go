package main

import (
	"deadpool/adapters/auth"
	"deadpool/adapters/http"
	"deadpool/adapters/http/middlewares"
	"deadpool/adapters/persistence"
	"deadpool/config"
	"deadpool/core/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
        log.Println("No .env file found, loading environment variables from system")
    }

    clientID := os.Getenv("GOOGLE_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
    redirectURL := "http://localhost:8080/api/auth/google/callback"

    db := config.InitDB()
    googleAuth := auth.NewGoogleAuth(clientID, clientSecret, redirectURL)
    userRepo := persistence.NewUserRepository(db)
    authService := services.AuthService{JWTSecret: "test", UserRepo: userRepo}
	userService := services.UserService{UserRepo: userRepo}
    authMiddleware := middlewares.NewAuthMiddleware(&authService)


    userHandler := http.NewUserHandler(&userService)
    authHandler := http.NewAuthHandler(&authService, googleAuth)

    app := fiber.New()

    // Public Routes
    app.Get("/api/auth/google/login", authHandler.GoogleLogin)
    app.Get("/api/auth/google/callback", authHandler.GoogleCallback)

    // Protected Routes
    api := app.Group("/api/auth")
    api.Use(authMiddleware.Handle)
    api.Get("/me", userHandler.GetMe)

    log.Println("Starting server on :8080")
    log.Fatal(app.Listen(":8080"))
}
