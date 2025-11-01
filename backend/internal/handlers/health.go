package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amilcar-vasquez/auth-service/backend/internal/utils"
)

var startTime = time.Now()

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

// Health returns the health status of the service
func Health(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)
	utils.RespondSuccess(w, HealthResponse{
		Status: "ok",
		Uptime: formatDuration(uptime),
	})
}

// formatDuration formats a duration into a human-readable string
func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}
	return fmt.Sprintf("%ds", seconds)
}
