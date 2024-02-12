package validate

import (
	"testing"
)

// TestIsValidCreditCard verifica se a função IsValidCreditCard funciona corretamente.
func TestIsValidCreditCard(t *testing.T) {
	// Teste para um número de cartão válido
	validCard := "1234567812345670"
	if !IsValidCreditCard(validCard) {
		t.Errorf("Erro para cartão válido. Esperado true, obteve false.")
	}

	// Teste para um número de cartão inválido
	invalidCard := "1234-5678-1234-5678"
	if IsValidCreditCard(invalidCard) {
		t.Errorf("Erro para cartão inválido. Esperado false, obteve true.")
	}
}

// TestValidateCreditCard verifica se a função ValidateCreditCard funciona corretamente.
func TestValidateCreditCard(t *testing.T) {
	// Teste para um número de cartão válido
	validCard := "4539445433274775"
	isValid, brand := ValidateCreditCard(validCard)
	if !isValid || brand != "Visa" {
		t.Errorf("Erro para cartão válido. Esperado (true, 'Visa'), obteve (%t, '%s').", isValid, brand)
	}

	// Teste para um número de cartão inválido
	invalidCard := "1234-5678-1234-5678"
	isValid, brand = ValidateCreditCard(invalidCard)
	if isValid || brand != "" {
		t.Errorf("Erro para cartão inválido. Esperado (false, ''), obteve (%t, '%s').", isValid, brand)
	}
}

// TestIdentifyCardBrand verifica se a função IdentifyCardBrand funciona corretamente.
func TestIdentifyCardBrand(t *testing.T) {
	// Teste para diferentes marcas de cartão
	visaCard := "4123456781234567"
	if brand := IdentifyCardBrand(visaCard); brand != "Visa" {
		t.Errorf("Erro para cartão Visa. Esperado 'Visa', obteve '%s'.", brand)
	}

	masterCard := "5212345678901234"
	if brand := IdentifyCardBrand(masterCard); brand != "MasterCard" {
		t.Errorf("Erro para cartão MasterCard. Esperado 'MasterCard', obteve '%s'.", brand)
	}

	amexCard := "371234567890123"
	if brand := IdentifyCardBrand(amexCard); brand != "American Express" {
		t.Errorf("Erro para cartão American Express. Esperado 'American Express', obteve '%s'.", brand)
	}

}
