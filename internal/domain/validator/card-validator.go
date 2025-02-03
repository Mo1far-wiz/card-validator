package validator

import (
	"card-validator/internal/domain/models"
	luhn "card-validator/internal/utils/luhn"
	"fmt"
	"log"
	"time"
)

type CreditCardValidator struct {
	Logger *log.Logger
}

const expDateLayout = "2006-01"

func (cv *CreditCardValidator) Validate(c models.Card) error {
	cv.Logger.Printf("Validating card: %v\v", c)

	cardExpDate := fmt.Sprintf("%04d-%02d", c.ExpYear, c.ExpMonth)
	expDate, err := time.Parse(expDateLayout, cardExpDate)
	if err != nil {
		cv.Logger.Printf("Card: %v has problems with expiration date: %s", c, err.Error())
		return ErrorOnParsingExpirationDate
	}
	if !expDate.After(time.Now()) {
		cv.Logger.Printf("Card: %v has problems with expiration date", c)
		return ErrorCardExpired
	}

	if !luhn.CheckLuhn(c.CardNumber) {
		cv.Logger.Printf("Card: %v has problems with card number", c)
		return ErrorWrongCardNumber
	}

	cv.Logger.Printf("Card: %v passed all validation checks", c)
	return nil
}
