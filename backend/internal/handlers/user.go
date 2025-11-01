package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/amilcar-vasquez/auth-service/backend/internal/middleware"
	"github.com/amilcar-vasquez/auth-service/backend/internal/models"
	"github.com/amilcar-vasquez/auth-service/backend/internal/utils"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

// UpdateProfileRequest represents the profile update payload
type UpdateProfileRequest struct {
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

// GetProfile retrieves the authenticated user's profile
func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.RespondError(w, http.StatusNotFound, "User not found")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	utils.RespondSuccess(w, user)
}

// UpdateProfile updates the authenticated user's profile
func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Retrieve user
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.RespondError(w, http.StatusNotFound, "User not found")
			return
		}
		utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = strings.TrimSpace(req.Name)
	}
	if req.Avatar != "" {
		user.Avatar = strings.TrimSpace(req.Avatar)
	}

	// Save changes
	if err := h.DB.Save(&user).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to update profile")
		return
	}

	utils.RespondSuccess(w, user)
}

// DeleteProfile deletes the authenticated user's account
func (h *UserHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserIDFromContext(r)
	if !ok {
		utils.RespondError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	// Soft delete the user
	if err := h.DB.Delete(&models.User{}, userID).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to delete account")
		return
	}

	utils.RespondSuccessWithMessage(w, "Account deleted successfully")
}
