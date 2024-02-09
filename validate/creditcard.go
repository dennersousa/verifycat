package validate

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// ValidateCreditCardHandler é um handler para a rota /validate com tipo "creditcard".
func ValidateCreditCardHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, brand := ValidateCreditCard(req.Value)

	result := gin.H{
		"ValidationResult": ValidationResult{
			IsValid: isValid,
			Message: "Credit Card",
		},
		"CreditCard": CreditCardResult{
			Brand: brand,
		},
	}

	c.JSON(http.StatusOK, result)
}

// IsValidCreditCard verifica se um número de cartão de crédito é válido.
func IsValidCreditCard(cardNumber string) bool {
	regex := regexp.MustCompile(`^\d{16}$`)
	return regex.MatchString(cardNumber)
}

// ValidateCreditCard valida um número de cartão de crédito e identifica a marca.
func ValidateCreditCard(cardNumber string) (bool, string) {
	// Remove espaços e traços do número do cartão
	re := regexp.MustCompile(`\s|-`)
	cardNumber = re.ReplaceAllString(cardNumber, "")

	// Verifique se o número do cartão tem o comprimento correto
	if ok, _ := regexp.MatchString(`^\d{13,16}$`, cardNumber); !ok {
		return false, ""
	}

	// Identificar a marca do cartão
	brand := IdentifyCardBrand(cardNumber)

	// Se a marca estiver vazia, o número do cartão é inválido
	if brand == "" {
		return false, ""
	}

	// Se chegar aqui, o número do cartão é válido, e a marca é identificada
	return true, brand
}

// IdentifyCardBrand identifica a marca de um cartão de crédito com base no número.
func IdentifyCardBrand(cardNumber string) string {
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
	// Adicione mais padrões de identificação conforme necessário
	default:
		return ""
	}
}
