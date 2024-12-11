package tests

import (
	"card-validator/internal/models"
	"card-validator/internal/utils"
	"card-validator/internal/validator"
	"fmt"
	"time"
)

type TestCreditCardValidator struct{}

const expDateLayout = "2006-01"

var TimeNow = time.Date(2024, 12, 12, 10, 0, 0, 0, time.UTC)

func (cv *TestCreditCardValidator) Validate(c models.Card) error {
	cardExpDate := fmt.Sprintf("%04d-%02d", c.ExpYear, c.ExpMonth)
	expDate, err := time.Parse(expDateLayout, cardExpDate)
	if err != nil {
		return validator.ErrorOnParsingExpirationDate
	}
	if !expDate.After(TimeNow) {
		return validator.ErrorCardExpired
	}

	if !utils.CheckLuhn(c.CardNumber) {
		return validator.ErrorWrongCardNumber
	}
	return nil
}
