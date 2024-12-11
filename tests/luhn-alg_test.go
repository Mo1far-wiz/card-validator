package tests

import (
	"card-validator/internal/utils"
	"testing"
)

func TestIsCodeOk(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid card numbers
		{"Valid: American Express", "371449635398431", true},
		{"Valid: Diners Club", "30569309025904", true},
		{"Valid: Discover", "6011111111111117", true},
		{"Valid: JCB", "3530111333300000", true},
		{"Valid: Mastercard", "5555555555554444", true},
		{"Valid: Visa", "4111111111111111", true},

		// Invalid card numbers
		{"Invalid: American Express", "371449635398432", false},
		{"Invalid: Diners Club", "3056930902594", false},
		{"Invalid: Discover", "6011111111111118", false},
		{"Invalid: JCB", "3530111333301000", false},
		{"Invalid: Mastercard", "5555555555554344", false},
		{"Invalid: Visa", "4111111111111112", false},

		// Edge cases
		{"Empty string", "", false},
		{"Only dashes", "----", false},
		{"Mixed valid with dashes", "4111-1111-1111-1111", true},
		{"Mixed invalid with dashes", "4111-1111-1111-1112", false},
		{"Non-numeric characters", "4111-1111-1111-ABCD", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.CheckLuhn(tt.input)
			if result != tt.expected {
				t.Errorf("IsCodeOk(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
