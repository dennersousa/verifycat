package validate

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IsValidCNPJHandler é um handler para a rota /validate com tipo "cnpj".
func IsValidCNPJHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid := IsValidCNPJ(req.Value)

	result := ValidationResult{
		IsValid: isValid,
		Message: "CNPJ",
	}

	c.JSON(http.StatusOK, result)
}

// IsValidCNPJ verifica se um CNPJ é válido.
func IsValidCNPJ(cnpj string) bool {
	// Remove não dígitos do CNPJ
	re := regexp.MustCompile(`[^\d]`)
	cnpj = re.ReplaceAllString(cnpj, "")

	// Verifique se o CNPJ está vazio ou não tem 14 dígitos
	if cnpj == "" || len(cnpj) != 14 {
		return false
	}

	// Verifique se há dígitos repetidos no CNPJ
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[i-1] {
			// Dígitos não são repetidos, continue com a validação
			break
		}
		if i == len(cnpj)-1 {
			// Todos os dígitos são repetidos, retorne false
			return false
		}
	}

	// Valide os dígitos de verificação do CNPJ
	size := len(cnpj) - 2
	numbers := cnpj[:size]
	digits := cnpj[size:]
	sum := 0
	pos := size - 7

	// Calcule o primeiro dígito de verificação
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

	// Compare o dígito calculado com o verdadeiro primeiro dígito de verificação
	firstDigit, _ := strconv.Atoi(string(digits[0]))
	if result != firstDigit {
		return false
	}

	// Calcule o segundo dígito de verificação
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

	// Compare o dígito calculado com o verdadeiro segundo dígito de verificação
	secondDigit, _ := strconv.Atoi(string(digits[1]))
	return result == secondDigit
}
