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

// Mapeamento de tipos para funções de validação
var validationFuncMap = map[string]func(string) (bool, string){
	"cpf":        func(value string) (bool, string) { return IsValidCPF(value), "" },
	"cnpj":       func(value string) (bool, string) { return IsValidCNPJ(value), "" },
	"url":        func(value string) (bool, string) { return IsValidURL(value), "" },
	"creditcard": ValidateCreditCard,
	"email":      func(value string) (bool, string) { return IsValidEmail(value), "" },
}

func ValidateHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationFunc, exists := validationFuncMap[req.Type]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid validation type"})
		return
	}

	var isValid bool
	var brand string

	if req.Type == "creditcard" {
		isValid, brand = validationFunc(req.Value)
	} else {
		isValid, _ = validationFunc(req.Value)
	}

	var result interface{}
	if req.Type == "creditcard" {
		result = struct {
			ValidationResult
			CreditCard CreditCardResult `json:"creditcard,omitempty"`
		}{
			ValidationResult: ValidationResult{
				IsValid: isValid,
				Message: req.Type,
			},
			CreditCard: CreditCardResult{
				Brand: brand,
			},
		}
	} else {
		result = ValidationResult{
			IsValid: isValid,
			Message: req.Type,
		}
	}

	c.JSON(http.StatusOK, result)
}
