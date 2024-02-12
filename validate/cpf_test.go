package validate

import (
	"fmt"
	"testing"
)

func TestIsValidCPFAllSameNumbers(t *testing.T) {
	// Test with CPFs where all numbers are the same (0 to 9)
	for digit := 0; digit <= 9; digit++ {
		allSameCPF := fmt.Sprintf("%d%d%d.%d%d%d.%d%d%d-%d%d", digit, digit, digit, digit, digit, digit, digit, digit, digit, digit, digit)
		result := IsValidCPF(allSameCPF)

		if result {
			t.Errorf("Expected CPF %s to be invalid, but got valid", allSameCPF)
		}
	}
}

func TestIsValidCPFValid(t *testing.T) {
	// Test a valid CPF
	validCPF := "123.456.789-09"
	result := IsValidCPF(validCPF)

	if !result {
		t.Errorf("Expected CPF %s to be valid, but got invalid", validCPF)
	}
}

func TestIsValidCPFInvalid(t *testing.T) {
	// Test an invalid CPF
	invalidCPF := "111.222.333-44"
	result := IsValidCPF(invalidCPF)

	if result {
		t.Errorf("Expected CPF %s to be invalid, but got valid", invalidCPF)
	}
}

func TestIsValidCPFShortInput(t *testing.T) {
	// Test with a short input string
	shortCPF := "123.456"
	result := IsValidCPF(shortCPF)

	if result {
		t.Errorf("Expected CPF %s to be invalid due to short input, but got valid", shortCPF)
	}
}

func TestIsValidCPFNonNumericInput(t *testing.T) {
	// Test with non-numeric input
	nonNumericCPF := "ABC.DEF.GHI-JK"
	result := IsValidCPF(nonNumericCPF)

	if result {
		t.Errorf("Expected CPF %s to be invalid due to non-numeric characters, but got valid", nonNumericCPF)
	}
}
