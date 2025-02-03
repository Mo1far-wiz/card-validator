package validator

import "card-validator/internal/domain/models"

var (
	ErrorWrongCardNumber         = NewValidationError(1, "card number is wrong")
	ErrorCardExpired             = NewValidationError(2, "card is expired")
	ErrorOnParsingExpirationDate = NewValidationError(3, "encountered error on expiration date parsing")
)

type CardValidatorInterface interface {
	Validate(c models.Card) error
}

type CardValidator struct {
	validator CardValidatorInterface
}

func NewCardValidator(v CardValidatorInterface) CardValidator {
	return CardValidator{
		validator: v,
	}
}

func (v *CardValidator) Validate(c models.Card) error {
	return v.validator.Validate(c)
}
