package validate

import (
	"testing"
)

func TestIsValidURL_ValidURL_HTTP(t *testing.T) {
	urlStr := "http://www.example.com"
	result := IsValidURL(urlStr)

	if !result {
		t.Errorf("Expected URL '%s' to be valid, but it was invalid", urlStr)
	}
}

func TestIsValidURL_ValidURL_HTTPS(t *testing.T) {
	urlStr := "https://www.example.com"
	result := IsValidURL(urlStr)

	if !result {
		t.Errorf("Expected URL '%s' to be valid, but it was invalid", urlStr)
	}
}

func TestIsValidURL_InvalidScheme(t *testing.T) {
	urlStr := "ftp://www.example.com"
	result := IsValidURL(urlStr)

	if result {
		t.Errorf("Expected URL '%s' with invalid scheme to be invalid, but it was valid", urlStr)
	}
}

func TestIsValidURL_InvalidHost(t *testing.T) {
	urlStr := "http://invalid-host"
	result := IsValidURL(urlStr)

	if result {
		t.Errorf("Expected URL '%s' with invalid host to be invalid, but it was valid", urlStr)
	}
}

func TestIsValidURL_EmptyURL(t *testing.T) {
	urlStr := ""
	result := IsValidURL(urlStr)

	if result {
		t.Errorf("Expected empty URL to be invalid, but it was valid")
	}
}
