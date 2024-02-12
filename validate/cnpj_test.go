package validate

import (
	"testing"
)

func TestIsValidCNPJ(t *testing.T) {
	// Caso de teste para CNPJ válido
	validCNPJ := "11.222.333/0001-81"
	if !IsValidCNPJ(validCNPJ) {
		t.Errorf("Expected CNPJ %s to be valid, but got invalid", validCNPJ)
	}

	// Caso de teste para CNPJ inválido
	invalidCNPJ := "11.222.333/0001-82"
	if IsValidCNPJ(invalidCNPJ) {
		t.Errorf("Expected CNPJ %s to be invalid, but got valid", invalidCNPJ)
	}

	// Caso de teste para CNPJ vazio
	emptyCNPJ := ""
	if IsValidCNPJ(emptyCNPJ) {
		t.Errorf("Expected empty CNPJ to be invalid, but got valid")
	}

	// Caso de teste para CNPJ com tamanho incorreto
	invalidSizeCNPJ := "11.222.333"
	if IsValidCNPJ(invalidSizeCNPJ) {
		t.Errorf("Expected CNPJ %s with incorrect size to be invalid, but got valid", invalidSizeCNPJ)
	}

	// Caso de teste para CNPJ com dígitos repetidos
	repeatedDigitsCNPJ := "11.111.111/1111-11"
	if IsValidCNPJ(repeatedDigitsCNPJ) {
		t.Errorf("Expected CNPJ %s with repeated digits to be invalid, but got valid", repeatedDigitsCNPJ)
	}

	// Caso de teste para CNPJ sem caracteres especiais
	withoutSpecialCharsCNPJ := "11222333000181"
	if !IsValidCNPJ(withoutSpecialCharsCNPJ) {
		t.Errorf("Expected CNPJ %s without special characters to be valid, but got invalid", withoutSpecialCharsCNPJ)
	}
}
