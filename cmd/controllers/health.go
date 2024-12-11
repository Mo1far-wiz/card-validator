package controllers

import (
	"card-validator/internal/utils"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "healthy",
	}

	if err := utils.ResponseJSON(w, http.StatusOK, data); err != nil {
		internalServerError(w, r, err)
	}
}
