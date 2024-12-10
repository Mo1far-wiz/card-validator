package main

import (
	"card-validator/internal/models"
	"net/http"
)

func (app *application) validateCardHandler(w http.ResponseWriter, r *http.Request) {
	type ValidationResult struct {
		Valid bool   `json:"valid"`
		Error string `json:"error,omitempty"`
	}

	var card models.Card
	if err := readJSON(w, r, &card); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	if err := Validate.Struct(card); err != nil {
		app.badRequestError(w, r, err)
		return
	}

	err := card.Validate()
	var result ValidationResult
	if err != nil {
		result.Error = err.Error()
	}

	result.Valid = err == nil

	if err := app.jsonResponse(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
