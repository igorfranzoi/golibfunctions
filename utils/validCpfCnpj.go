package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// Valida CPF e CNPJ (numérico ou alfanumérico)
func Valid_CNPJ_CPF(strDocumentNumber string) bool {
	strDocumentNumber = strings.TrimSpace(strDocumentNumber)

	// Remove espaços e separadores
	strDocumentNumber = regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(strDocumentNumber, "")

	if len(strDocumentNumber) == 11 && isAllDigits(strDocumentNumber) {
		return validCPF(strDocumentNumber)
	} else if len(strDocumentNumber) == 14 {
		if isAllDigits(strDocumentNumber) {
			return validCNPJ(strDocumentNumber)
		} else {
			// Novo modelo alfanumérico de CNPJ
			return true // Considera válido por enquanto (sem verificação de DV)
		}
	}
	return false
}

func isAllDigits(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func validCPF(strCPF string) bool {
	var weightNumber, sumNumber int

	for i := 0; i < 9; i++ {
		weightNumber += (10 - i) * int(strCPF[i]-'0')
	}

	if remainder := weightNumber % 11; remainder < 2 {
		sumNumber = 0
	} else {
		sumNumber = 11 - remainder
	}

	if int(strCPF[9]-'0') != sumNumber {
		return false
	}

	weightNumber = 0
	for i := 0; i < 10; i++ {
		weightNumber += (11 - i) * int(strCPF[i]-'0')
	}

	if remainder := weightNumber % 11; remainder < 2 {
		sumNumber = 0
	} else {
		sumNumber = 11 - remainder
	}

	return int(strCPF[10]-'0') == sumNumber
}

func validCNPJ(strCNPJ string) bool {
	var typeOne, typeTwo int
	weightNumber1 := [12]int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weightNumber2 := [13]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sumNumber := 0
	for i := 0; i < 12; i++ {
		sumNumber += int(strCNPJ[i]-'0') * weightNumber1[i]
	}

	remainder := sumNumber % 11
	if remainder < 2 {
		typeOne = 0
	} else {
		typeOne = 11 - remainder
	}

	if int(strCNPJ[12]-'0') != typeOne {
		return false
	}

	sumNumber = 0
	for i := 0; i < 13; i++ {
		sumNumber += int(strCNPJ[i]-'0') * weightNumber2[i]
	}

	remainder = sumNumber % 11
	if remainder < 2 {
		typeTwo = 0
	} else {
		typeTwo = 11 - remainder
	}

	return int(strCNPJ[13]-'0') == typeTwo
}
