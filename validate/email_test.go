package validate

import (
	"testing"
)

func TestIsValidEmail_ValidEmail(t *testing.T) {
	email := "test@example.com"
	result := IsValidEmail(email)

	if !result {
		t.Errorf("Expected email '%s' to be valid, but it was invalid", email)
	}
}

func TestIsValidEmail_InvalidEmail(t *testing.T) {
	email := "invalid-email"
	result := IsValidEmail(email)

	if result {
		t.Errorf("Expected email '%s' to be invalid, but it was valid", email)
	}
}

func TestIsValidEmail_EmptyEmail(t *testing.T) {
	email := ""
	result := IsValidEmail(email)

	if result {
		t.Errorf("Expected empty email to be invalid, but it was valid")
	}
}

func TestIsValidEmail_InvalidCharacters(t *testing.T) {
	email := "test!@example.com"
	result := IsValidEmail(email)

	if result {
		t.Errorf("Expected email '%s' with invalid characters to be invalid, but it was valid", email)
	}
}

func TestIsValidEmail_InvalidDomain(t *testing.T) {
	email := "test@example"
	result := IsValidEmail(email)

	if result {
		t.Errorf("Expected email '%s' with invalid domain to be invalid, but it was valid", email)
	}
}

func TestIsValidEmail_InvalidTLD(t *testing.T) {
	email := "test@example.c"
	result := IsValidEmail(email)

	if result {
		t.Errorf("Expected email '%s' with invalid top-level domain to be invalid, but it was valid", email)
	}
}
