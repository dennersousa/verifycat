package validations

import (
	"regexp"
)

func IsValidCreditCard(cardNumber string) bool {
	regex := regexp.MustCompile(`^\d{16}$`)
	return regex.MatchString(cardNumber)
}

func ValidateCreditCard(cardNumber string) (bool, string) {
	// Remove spaces and dashes from the card number
	re := regexp.MustCompile(`\s|-`)
	cardNumber = re.ReplaceAllString(cardNumber, "")

	// Check if the card number has the correct length
	if ok, _ := regexp.MatchString(`^\d{13,16}$`, cardNumber); !ok {
		return false, ""
	}

	// Identify the card brand
	brand := IdentifyCardBrand(cardNumber)

	// If the brand is empty, the card number is invalid
	if brand == "" {
		return false, ""
	}

	// If it reaches here, the card number is valid, and the brand is identified
	return true, brand
}

func IdentifyCardBrand(cardNumber string) string {
	// Some brand identification patterns
	switch {
	case regexp.MustCompile(`^4`).MatchString(cardNumber):
		return "Visa"
	case regexp.MustCompile(`^5[1-5]`).MatchString(cardNumber):
		return "MasterCard"
	case regexp.MustCompile(`^3[47]`).MatchString(cardNumber):
		return "American Express"
	case regexp.MustCompile(`^(36|38|30[0-5])`).MatchString(cardNumber):
		return "Diners Club / Carte Blanche"
	case regexp.MustCompile(`^6011`).MatchString(cardNumber):
		return "Discover"
	case regexp.MustCompile(`^(3[0-5]|21[23]|1800)`).MatchString(cardNumber):
		return "JCB"
	// Add more identification patterns as needed
	default:
		return ""
	}
}
