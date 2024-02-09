package validations

import (
	"net/url"
	"regexp"
)

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
