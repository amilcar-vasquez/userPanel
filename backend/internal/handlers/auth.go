package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/amilcar-vasquez/auth-service/backend/internal/models"
	"github.com/amilcar-vasquez/auth-service/backend/internal/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

// RegisterRequest represents the registration payload
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest represents the login payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

// Register creates a new user account
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate input
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.Name = strings.TrimSpace(req.Name)

	if req.Email == "" || req.Name == "" || req.Password == "" {
		utils.RespondError(w, http.StatusBadRequest, "Name, email, and password are required")
		return
	}

	if len(req.Password) < 6 {
		utils.RespondError(w, http.StatusBadRequest, "Password must be at least 6 characters")
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := h.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.RespondError(w, http.StatusConflict, "Email already registered")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to process password")
		return
	}

	// Create user
	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hashedPassword,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondSuccess(w, AuthResponse{
		Token: token,
		User:  &user,
	})
}

// Login authenticates a user and returns a JWT token
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Normalize email
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))

	if req.Email == "" || req.Password == "" {
		utils.RespondError(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	// Find user by email
	var user models.User
	if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Verify password
	if !utils.VerifyPassword(user.PasswordHash, req.Password) {
		utils.RespondError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondSuccess(w, AuthResponse{
		Token: token,
		User:  &user,
	})
}
