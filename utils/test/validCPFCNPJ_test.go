package tests

import (
	"golibfunctions/utils"
	"testing"
)

func TestValidCPFCNPJSuccess(t *testing.T) {

	// Exemplo de CPF e CNPJ para teste
	strCPF := "332.945.830-52"
	strCNPJ := "61.695.227/0001-93"

	// Verificar se o número do CPF é válido
	if !utils.Valid_CNPJ_CPF(strCPF) {
		t.Errorf("O CPF '%s' é um número válido, porém não foi validado pela função! - \n", strCPF)
	}
	// Verificar se o número do CNPJ é válido
	if !utils.Valid_CNPJ_CPF(strCNPJ) {
		t.Errorf("O CNPJ '%s' é um número válido, porém não foi validado pela função! - \n", strCNPJ)
	}

	/*strCNPJ = "BR12A56780001Z5"

	if !utils.Valid_CNPJ_CPF(strCNPJ) {
		t.Errorf("O CNPJ '%s' é um número válido, porém não foi validado pela função! - \n", strCNPJ)
	}*/

}

func TestValidCPFCNPJFailed(t *testing.T) {

	// Exemplo de CPF e CNPJ para teste
	strCPF := "33294583000"
	strCNPJ := "61695227000100"

	// Verificar se o número do CPF é válido
	if utils.Valid_CNPJ_CPF(strCPF) {
		t.Errorf("O CPF '%s' é um número inválido, porém a verificação validou o número! - \n", strCPF)
	}
	// Verificar se o número do CNPJ é válido
	if utils.Valid_CNPJ_CPF(strCNPJ) {
		t.Errorf("O CNPJ '%s' é um número inválido, porém a verificação validou o número! - \n", strCNPJ)
	}

}
