package tests

import (
	"card-validator/internal/domain/models"
	"card-validator/internal/domain/validator"
	luhn "card-validator/internal/utils/luhn"
	"fmt"
	"time"
)

type TestCreditCardValidator struct {
	// it is needed to be independent from time.Now() and don't change tests each month
	TimeNow time.Time
}

const expDateLayout = "2006-01"

func (cv *TestCreditCardValidator) Validate(c models.Card) error {
	cardExpDate := fmt.Sprintf("%04d-%02d", c.ExpYear, c.ExpMonth)
	expDate, err := time.Parse(expDateLayout, cardExpDate)
	if err != nil {
		return validator.ErrorOnParsingExpirationDate
	}
	if !expDate.After(cv.TimeNow) {
		return validator.ErrorCardExpired
	}

	if !luhn.CheckLuhn(c.CardNumber) {
		return validator.ErrorWrongCardNumber
	}
	return nil
}
