package main

import (
	"card-validator/internal/models"
	"testing"
)

func TestCardExpDateValidation(t *testing.T) {
	tests := []struct {
		name     string
		card     models.Card
		wantErr  bool
		errorMsg string
	}{
		{
			name: "Valid card",
			card: models.Card{
				CardNumber: "3714-4963-5398-431",
				ExpMonth:   12,
				ExpYear:    2025,
			},
			wantErr: false,
		},
		{
			name: "Expired card",
			card: models.Card{
				CardNumber: "3714-4963-5398-431",
				ExpMonth:   12,
				ExpYear:    2020,
			},
			wantErr:  true,
			errorMsg: "card is expired",
		},
		{
			name: "Invalid expiration date",
			card: models.Card{
				CardNumber: "3714-4963-5398-431",
				ExpMonth:   13,
				ExpYear:    2024,
			},
			wantErr:  true,
			errorMsg: "parsing time",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.card.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err != nil && !contains(err.Error(), tt.errorMsg) {
				t.Errorf("Expected error message to contain '%s', got '%s'", tt.errorMsg, err.Error())
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s[:len(substr)] == substr || contains(s[1:], substr))
}
