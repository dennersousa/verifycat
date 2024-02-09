package validate

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// IsValidCPFHandler é um handler para a rota /validate com tipo "cpf".
func IsValidCPFHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid := IsValidCPF(req.Value)

	result := ValidationResult{
		IsValid: isValid,
		Message: "CPF",
	}

	c.JSON(http.StatusOK, result)
}

// IsValidCPF checks if a CPF (Brazilian ID number) is valid.
// The input should be a string containing only numeric digits.
func IsValidCPF(cpf string) bool {
	// Remove non-numeric characters from CPF
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Check for known invalid CPFs
	if strings.Count(cpf, string(cpf[0])) == 11 {
		return false
	}

	// Extract digits from CPF
	digits := make([]int, 11)
	for i := 0; i < 11; i++ {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			// Handle error as needed
		}
		digits[i] = digit
	}

	// Calculate the first verification digit
	sum1 := digits[0]*10 + digits[1]*9 + digits[2]*8 + digits[3]*7 + digits[4]*6 + digits[5]*5 + digits[6]*4 + digits[7]*3 + digits[8]*2
	remainder1 := (sum1 * 10) % 11
	if remainder1 == 10 {
		remainder1 = 0
	}

	// Calculate the second verification digit
	sum2 := digits[0]*11 + digits[1]*10 + digits[2]*9 + digits[3]*8 + digits[4]*7 + digits[5]*6 + digits[6]*5 + digits[7]*4 + digits[8]*3 + digits[9]*2
	remainder2 := (sum2 * 10) % 11
	if remainder2 == 10 {
		remainder2 = 0
	}

	// Check verification digits
	return remainder1 == digits[9] && remainder2 == digits[10]
}
