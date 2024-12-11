package validator

import "card-validator/internal/models"

type CardValidatorInterface interface {
	Validate(c models.Card) error
}

type CardValidator struct {
	validator CardValidatorInterface
}

func NewCardValidator(cv CardValidatorInterface) CardValidator {
	return CardValidator{
		validator: cv,
	}
}

func (v *CardValidator) Validate(c models.Card) error {
	return v.validator.Validate(c)
}
