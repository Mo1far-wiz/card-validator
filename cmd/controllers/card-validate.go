package controllers

import (
	"card-validator/internal/models"
	"card-validator/internal/utils"
	"card-validator/internal/validator"
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
	if err := utils.ReadJSON(w, r, &req); err != nil {
		badRequestError(w, r, err)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
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
		default:
			resp.Error = &errorResponse{Code: validator.ErrorOnParsingExpirationDate.Code, Message: validator.ErrorOnParsingExpirationDate.Message}
		}
	}

	resp.Valid = err == nil

	if err := utils.ResponseJSON(w, http.StatusOK, resp); err != nil {
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
