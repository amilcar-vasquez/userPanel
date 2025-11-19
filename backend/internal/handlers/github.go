package handlers

import (
	"net/http"

	"github.com/amilcar-vasquez/auth-service/backend/internal/github"
	"github.com/amilcar-vasquez/auth-service/backend/internal/middleware"
	"github.com/amilcar-vasquez/auth-service/backend/internal/models"
	"github.com/amilcar-vasquez/auth-service/backend/internal/utils"
	"gorm.io/gorm"
)

// GetGithubProfile fetches GitHub profile statistics for the authenticated user
func GetGithubProfile(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user from context (set by auth middleware)
		userID, ok := middleware.GetUserIDFromContext(r)
		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, "User ID not found in context")
			return
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			utils.RespondError(w, http.StatusNotFound, "User not found")
			return
		}

		// Check if user has GitHub credentials configured
		if user.GithubUsername == "" || user.GithubToken == "" {
			utils.RespondError(w, http.StatusBadRequest, "GitHub username or token not configured. Please update your profile first.")
			return
		}

		// Fetch GitHub profile stats
		stats, err := github.FetchUserProfile(user.GithubUsername, user.GithubToken)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch GitHub profile: "+err.Error())
			return
		}

		// Calculate developer rank
		rank := github.CalculateRank(*stats)

		// Return stats with rank information
		response := map[string]interface{}{
			"profile": stats,
			"rank":    rank,
		}

		utils.RespondSuccess(w, response)
	}
}

// UpdateGithubCredentials updates the user's GitHub username and token
func UpdateGithubCredentials(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user from context
		userID, ok := middleware.GetUserIDFromContext(r)
		if !ok {
			utils.RespondError(w, http.StatusUnauthorized, "User ID not found in context")
			return
		}

		var req struct {
			GithubUsername string `json:"github_username"`
			GithubToken    string `json:"github_token"`
		}

		if err := utils.ParseJSON(r, &req); err != nil {
			utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		// Validate that at least one field is provided
		if req.GithubUsername == "" && req.GithubToken == "" {
			utils.RespondError(w, http.StatusBadRequest, "GitHub username or token required")
			return
		}

		// Update user's GitHub credentials
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			utils.RespondError(w, http.StatusNotFound, "User not found")
			return
		}

		// Update fields if provided
		if req.GithubUsername != "" {
			user.GithubUsername = req.GithubUsername
		}
		if req.GithubToken != "" {
			user.GithubToken = req.GithubToken
		}

		if err := db.Save(&user).Error; err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to update GitHub credentials")
			return
		}

		utils.RespondSuccess(w, map[string]interface{}{
			"message":         "GitHub credentials updated successfully",
			"github_username": user.GithubUsername,
		})
	}
}
