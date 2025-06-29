package tests

import (
	"fmt"

	"testing"

	"github.com/igorfranzoi/golibfunctions/utils"
	"github.com/stretchr/testify/assert"
)

func TestRemoveSpecialChars(t *testing.T) {

	strOrigin := `Arquivo\com/um/caminho\especial?*|<>:"`
	strRemoved := utils.RemoveSpecialChars(strOrigin)

	fmt.Println("Original:", strOrigin)
	fmt.Println("Limpo:", strRemoved)

	assert.NotEqual(t, strOrigin, strRemoved, "Algo aconteceu de incosistente, pois o retorno deve ser diferente!")
}
