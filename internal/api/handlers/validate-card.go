package handlers

import (
	"card-validator/internal/domain/models"
	"card-validator/internal/domain/validator"
	json "card-validator/internal/utils/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type (
	request struct {
		CardNumber string `json:"card_number" validate:"required"`
		ExpMonth   int    `json:"expiration_month" validate:"required"`
		ExpYear    int    `json:"expiration_year" validate:"required"`
	}
	errorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	response struct {
		Valid bool           `json:"valid"`
		Error *errorResponse `json:"error,omitempty"`
	}
)

var CardValidator validator.CardValidator

func init() {
	logger := log.New(os.Stdout, "Custom INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	v := &validator.CreditCardValidator{
		Logger: logger,
	}

	CardValidator = validator.NewCardValidator(v)
}

func ValidateCardHandler(w http.ResponseWriter, r *http.Request) {

	var req request
	if err := json.ReadJSON(w, r, &req); err != nil {
		badRequestError(w, r, err)
		return
	}

	if err := json.Validate.Struct(req); err != nil {
		badRequestError(w, r, err)
		return
	}

	card := mapToDomain(req)

	var resp response
	err := CardValidator.Validate(card)

	if err != nil {
		switch {
		case errors.Is(err, validator.ErrorWrongCardNumber):
			resp.Error = &errorResponse{Code: validator.ErrorWrongCardNumber.Code, Message: validator.ErrorWrongCardNumber.Message}
		case errors.Is(err, validator.ErrorCardExpired):
			resp.Error = &errorResponse{Code: validator.ErrorCardExpired.Code, Message: validator.ErrorCardExpired.Message}
		case errors.Is(err, validator.ErrorOnParsingExpirationDate):
			resp.Error = &errorResponse{Code: validator.ErrorOnParsingExpirationDate.Code, Message: validator.ErrorOnParsingExpirationDate.Message}
		default:
			resp.Error = &errorResponse{Code: 0, Message: "unknown validation error: " + err.Error()}
		}
	}

	resp.Valid = err == nil

	if err := json.ResponseJSON(w, http.StatusOK, resp); err != nil {
		internalServerError(w, r, err)
	}
}

func mapToDomain(r request) models.Card {
	return models.Card{
		CardNumber: r.CardNumber,
		ExpMonth:   r.ExpMonth,
		ExpYear:    r.ExpYear,
	}
}
