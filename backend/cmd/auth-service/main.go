package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amilcar-vasquez/auth-service/backend/config"
	"github.com/amilcar-vasquez/auth-service/backend/internal/middleware"
	"github.com/amilcar-vasquez/auth-service/backend/internal/models"
	"github.com/amilcar-vasquez/auth-service/backend/internal/utils"
	"github.com/amilcar-vasquez/auth-service/backend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize JWT utilities
	utils.InitJWT(cfg.JWTSecret)

	// Connect to database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate database schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("✓ Database migration completed")

	// Create router and setup global middleware first
	router := chi.NewRouter()

	// Add CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.CORSOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Add logger middleware
	router.Use(middleware.Logger)

	// Setup routes after middleware
	routes.SetupRoutes(router, db)

	// Start server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("🚀 Auth Service running on port %s", cfg.Port)
	log.Printf("📡 CORS enabled for: %s", cfg.CORSOrigin)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
