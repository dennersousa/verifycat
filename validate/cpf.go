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

// IsValidCPF verifica se um CPF (Cadastro de Pessoa Física brasileiro) é válido.
// A entrada deve ser uma string contendo apenas dígitos numéricos.
func IsValidCPF(cpf string) bool {
	// Remove caracteres não numéricos do CPF
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

<<<<<<< Updated upstream
	// Check for known invalid CPFs
=======
	// Verifique se o CPF contém apenas dígitos numéricos
	if _, err := strconv.Atoi(cpf); err != nil {
		return false
	}

	// Verifique CPFs inválidos conhecidos
>>>>>>> Stashed changes
	if strings.Count(cpf, string(cpf[0])) == 11 {
		return false
	}

	// Extraia os dígitos do CPF
	digits := make([]int, 11)
	for i := 0; i < 11; i++ {
<<<<<<< Updated upstream
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			// Handle error as needed
=======
		if i < len(cpf) {
			digit, err := strconv.Atoi(string(cpf[i]))
			if err != nil {
				// Trate o erro conforme necessário
			}
			digits[i] = digit
		} else {
			// Trate o caso em que a string é muito curta
			return false
>>>>>>> Stashed changes
		}
		digits[i] = digit
	}

	// Calcule o primeiro dígito de verificação
	sum1 := digits[0]*10 + digits[1]*9 + digits[2]*8 + digits[3]*7 + digits[4]*6 + digits[5]*5 + digits[6]*4 + digits[7]*3 + digits[8]*2
	remainder1 := (sum1 * 10) % 11
	if remainder1 == 10 {
		remainder1 = 0
	}

	// Calcule o segundo dígito de verificação
	sum2 := digits[0]*11 + digits[1]*10 + digits[2]*9 + digits[3]*8 + digits[4]*7 + digits[5]*6 + digits[6]*5 + digits[7]*4 + digits[8]*3 + digits[9]*2
	remainder2 := (sum2 * 10) % 11
	if remainder2 == 10 {
		remainder2 = 0
	}

	// Verifique os dígitos de verificação
	return remainder1 == digits[9] && remainder2 == digits[10]
}
