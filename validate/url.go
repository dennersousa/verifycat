package validate

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
)

// IsValidURLHandler é um handler para a rota /validate com tipo "url".
func IsValidURLHandler(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid := IsValidURL(req.Value)

	result := ValidationResult{
		IsValid: isValid,
		Message: "URL",
	}

	c.JSON(http.StatusOK, result)
}

// IsValidURL verifica se uma URL é válida.
func IsValidURL(urlStr string) bool {
	// Tente fazer o parsing da URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Verifique se o esquema é http ou https
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Use uma expressão regular para validar o host
	hostRegex := regexp.MustCompile(`^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !hostRegex.MatchString(parsedURL.Host) {
		return false
	}

	return true
}
