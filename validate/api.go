package validate

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ValidationRequest representa a estrutura de dados para a requisição de validação.
type ValidationRequest struct {
	Type  string `json:"type"`  // Tipo de dado a ser validado (ex: "cpf", "cnpj", "url", "creditcard", "email").
	Value string `json:"value"` // Valor do dado a ser validado.
}

// ValidationResult representa o resultado da validação.
type ValidationResult struct {
	IsValid bool   `json:"isValid"` // Indica se o dado é válido ou não.
	Message string `json:"message"` // Mensagem relacionada à validação.
}

// CreditCardResult representa o resultado específico para validação de cartão de crédito.
type CreditCardResult struct {
	Brand string `json:"brand,omitempty"` // Marca (bandeira) do cartão de crédito, se aplicável.
}

// ValidateHandler é o manipulador para a rota de validação.
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
		message = "Cartão de Crédito"
	case "email":
		isValid = IsValidEmail(req.Value)
		message = "Email"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de validação inválido"})
		return
	}

	var result interface{}
	if req.Type == "creditcard" {
		result = struct {
			ValidationResult
			CreditCard CreditCardResult `json:"creditcard,omitempty"`
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
