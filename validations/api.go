package validations

import (
	"encoding/json"
	"net/http"
)

type ValidationRequest struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ValidationResult struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

type CreditCardResult struct {
	Brand string `json:"brand,omitempty"`
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var req ValidationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var isValid bool
	var message string
	var brand string

	switch req.Type {
	case "cpf":
		isValid = IsValidCPF(req.Value)
		message = "CPF"
	case "cnpj":
		isValid = IsValidCNPJ(req.Value)
		message = "CNPJ"
	case "url":
		isValid = IsValidURL(req.Value)
		message = "URL"
	case "creditcard":
		// Using functions from creditcard.go
		isValid, brand = ValidateCreditCard(req.Value)
		message = "Credit Card"
	default:
		http.Error(w, "Invalid validation type", http.StatusBadRequest)
		return
	}

	var result interface{}
	if req.Type == "creditcard" {
		result = struct {
			ValidationResult
			CreditCard CreditCardResult `json:"creditCard,omitempty"`
		}{
			ValidationResult: ValidationResult{
				IsValid: isValid,
				Message: message,
			},
			CreditCard: CreditCardResult{
				Brand: brand,
			},
		}
	} else {
		result = ValidationResult{
			IsValid: isValid,
			Message: message,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
