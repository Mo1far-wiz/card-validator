package models

import (
	"card-validator/internal/utils"
	"errors"
	"fmt"
	"time"
)

type Card struct {
	CardNumber string `json:"card_number" validate:"required"`
	ExpMonth   int    `json:"expiration_month" validate:"required"`
	ExpYear    int    `json:"expiration_year" validate:"required"`
}

const expDateLayout = "2006-01"

func (c Card) Validate() error {
	cardExpDate := fmt.Sprintf("%04d-%02d", c.ExpYear, c.ExpMonth)
	expDate, err := time.Parse(expDateLayout, cardExpDate)
	if err != nil {
		return err
	}
	if !expDate.After(time.Now()) {
		return errors.New("card is expired")
	}

	if !utils.CheckLuhn(c.CardNumber) {
		return errors.New("card number is wrong")
	}

	return nil
}
