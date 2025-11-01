package routes

import (
	"github.com/amilcar-vasquez/auth-service/backend/internal/handlers"
	"github.com/amilcar-vasquez/auth-service/backend/internal/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// SetupRoutes configures all application routes
func SetupRoutes(r *chi.Mux, db *gorm.DB) {

	// Initialize handlers
	authHandler := &handlers.AuthHandler{DB: db}
	userHandler := &handlers.UserHandler{DB: db}

	// Public routes
	r.Route("/api", func(r chi.Router) {
		// Health check
		r.Get("/health", handlers.Health)

		// Authentication routes (public)
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)

			// User profile routes
			r.Get("/profile", userHandler.GetProfile)
			r.Put("/profile", userHandler.UpdateProfile)
			r.Delete("/profile", userHandler.DeleteProfile)
		})
	})
}
