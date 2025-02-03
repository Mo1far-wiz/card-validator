package tests

import (
	"card-validator/internal/api/handlers"
	"card-validator/internal/domain/models"
	"card-validator/internal/domain/validator"
	"card-validator/tests/mock"
	"errors"
	"testing"
	"time"
)

func SetupValidator() {
	handlers.CardValidator = validator.NewCardValidator(&mock.Validator{
		TimeNow: time.Date(2024, 12, 12, 10, 0, 0, 0, time.UTC),
	})
}

func TestValidCards(t *testing.T) {
	SetupValidator()

	tests := []models.Card{
		{
			CardNumber: "3714-4963-5398-431",
			ExpMonth:   1,
			ExpYear:    2025,
		},
		{
			CardNumber: "3550998650131033",
			ExpMonth:   12,
			ExpYear:    2028,
		},
		{
			CardNumber: "4457010000000009",
			ExpMonth:   4,
			ExpYear:    2027,
		},
		{
			CardNumber: "5167-0010-2023-6549",
			ExpMonth:   8,
			ExpYear:    2026,
		},
	}

	for _, tt := range tests {
		t.Run("Valid Card", func(t *testing.T) {
			err := handlers.CardValidator.Validate(tt)
			if err != nil {
				t.Errorf("Card %v expected to be Valid; Err: %s", tt, err.Error())
			}
		})
	}
}

func TestExpiredCards(t *testing.T) {
	SetupValidator()

	tests := []models.Card{
		{
			CardNumber: "3714-4963-5398-431",
			ExpMonth:   1,
			ExpYear:    2023,
		},
		{
			CardNumber: "3550998650131033",
			ExpMonth:   11,
			ExpYear:    2024,
		},
		{
			CardNumber: "4457010000000009",
			ExpMonth:   12,
			ExpYear:    2022,
		},
		{
			CardNumber: "5167-0010-2023-6549",
			ExpMonth:   7,
			ExpYear:    2016,
		},
	}

	for _, tt := range tests {
		t.Run("Expired Card", func(t *testing.T) {
			err := handlers.CardValidator.Validate(tt)
			if !errors.Is(err, validator.ErrorCardExpired) {
				t.Errorf("Card %v expected to have error %s; Err: %s", tt, validator.ErrorCardExpired.Error(), err.Error())
			}
		})
	}
}

func TestWrongCardNumbers(t *testing.T) {
	SetupValidator()

	tests := []models.Card{
		{
			CardNumber: "3715-4964-5398-432",
			ExpMonth:   1,
			ExpYear:    2025,
		},
		{
			CardNumber: "3551998650131033",
			ExpMonth:   12,
			ExpYear:    2028,
		},
		{
			CardNumber: "445701000000009",
			ExpMonth:   4,
			ExpYear:    2027,
		},
		{
			CardNumber: "5167-0011-2023-6549",
			ExpMonth:   8,
			ExpYear:    2026,
		},
	}

	for _, tt := range tests {
		t.Run("Wrong Card Number", func(t *testing.T) {
			err := handlers.CardValidator.Validate(tt)
			if !errors.Is(err, validator.ErrorWrongCardNumber) {
				t.Errorf("Card %v expected to have error %s; Err: %s", tt, validator.ErrorWrongCardNumber.Error(), err.Error())
			}
		})
	}
}

func TestMonthOrYearOutOfRange(t *testing.T) {
	SetupValidator()

	tests := []models.Card{
		{
			CardNumber: "3714-4963-5398-431",
			ExpMonth:   -1,
			ExpYear:    2025,
		},
		{
			CardNumber: "3550998650131033",
			ExpMonth:   1,
			ExpYear:    -1,
		},
		{
			CardNumber: "4457010000000009",
			ExpMonth:   14,
			ExpYear:    2027,
		},
		{
			CardNumber: "5167-0010-2023-6549",
			ExpMonth:   18,
			ExpYear:    2026,
		},
	}

	for _, tt := range tests {
		t.Run("Month Out Of Range", func(t *testing.T) {
			err := handlers.CardValidator.Validate(tt)
			if !errors.Is(err, validator.ErrorOnParsingExpirationDate) {
				t.Errorf("Card %v expected to have error %s; Err: %s", tt, validator.ErrorOnParsingExpirationDate.Error(), err.Error())
			}
		})
	}
}
