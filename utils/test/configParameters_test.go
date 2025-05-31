package tests

import (
	"golibfunctions/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnvMemSuccess(t *testing.T) {

	t.Parallel()

	err := utils.LoadEnvMem()

	assert.Equal(t, nil, err, "algo aconteceu de incorreto, pois o retorno deve ser NULO quando existir o arquivo .env!")

}
