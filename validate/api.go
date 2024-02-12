package validate

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func ValidateHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		isValid, brand = ValidateCreditCard(req.Value)
		message = "Credit Card"
	case "email":
		isValid = IsValidEmail(req.Value)
		message = "Email"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid validation type"})
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

	c.JSON(http.StatusOK, result)
}
