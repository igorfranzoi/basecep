package utils

import "regexp"

// Validar se o CEP fornecido é válido
func ValidCEPNumber(cepNumber string) bool {
	// Regex - CEP está no formato correto
	var cepRegex = regexp.MustCompile(`^\d{5}-?\d{3}$`)

	return cepRegex.MatchString(cepNumber)
}
