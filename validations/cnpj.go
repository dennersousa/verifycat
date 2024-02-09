package validations

import (
	"regexp"
	"strconv"
)

// IsValidCNPJ checks if a given CNPJ (Brazilian National Register of Legal Entities) is valid.
func IsValidCNPJ(cnpj string) bool {
	// Remove non-digits from the input CNPJ
	re := regexp.MustCompile(`[^\d]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	// Check if the CNPJ is empty or does not have 14 digits
	if cnpj == "" || len(cnpj) != 14 {
		return false
	}

	// Check for repeated digits in the CNPJ
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[i-1] {
			// Digits are not repeated, continue with the validation
			break
		}
		if i == len(cnpj)-1 {
			// All digits are repeated, return false
			return false
		}
	}

	// Validate CNPJ's verification digits
	size := len(cnpj) - 2
	numbers := cnpj[:size]
	digits := cnpj[size:]
	sum := 0
	pos := size - 7

	// Calculate the first verification digit
	for i := size; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numbers[size-i]))
		sum += num * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}
	result := sum % 11
	if result < 2 {
		result = 0
	} else {
		result = 11 - result
	}

	// Compare the calculated digit with the actual first verification digit
	firstDigit, _ := strconv.Atoi(string(digits[0]))
	if result != firstDigit {
		return false
	}

	// Calculate the second verification digit
	size++
	numbers = cnpj[:size]
	sum = 0
	pos = size - 7

	for i := size; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numbers[size-i]))
		sum += num * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}

	result = sum % 11
	if result < 2 {
		result = 0
	} else {
		result = 11 - result
	}

	// Compare the calculated digit with the actual second verification digit
	secondDigit, _ := strconv.Atoi(string(digits[1]))
	return result == secondDigit
}
