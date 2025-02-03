package handlers

import (
	json "card-validator/internal/utils/json"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "healthy",
	}

	if err := json.ResponseJSON(w, http.StatusOK, data); err != nil {
		internalServerError(w, r, err)
	}
}
