package validate

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// IsValidEmailHandler é um handler para a rota /validate com tipo "email".
func IsValidEmailHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid := IsValidEmail(req.Value)

	result := ValidationResult{
		IsValid: isValid,
		Message: "Email",
	}

	c.JSON(http.StatusOK, result)
}

// IsValidEmail verifica se um endereço de email é válido.
func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
