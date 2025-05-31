package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const strTopPath = "../"
const strEnvFile = ".env"
const strIdeRoot = "go.mod"

var (
	ErrEnvNotFound = errors.New("no .env file found")
	ErrEnvPathLoad = errors.New("não foi possível localizar o diretório da aplicação")
	ErrEnvRootPath = errors.New("não foi possível localizar o diretório principal")
)

// Função para carregar o arquivo env na memória, e devolver os parâmetros
func LoadEnvMem() error {

	// Verifica se o arquivo .env está na mesma pasta que a aplicação está executando
	_, err := os.Stat(strEnvFile)

	if errors.Is(err, os.ErrNotExist) {
		//Caso não tenha localizado o arquivo na pasta raiz, usa o comando de voltar
		//uma pasta para tentar localizar o arquivo .env
		//Esse tratamento é uma forma de os testes funcionarem normalmente, pois as funções
		//de teste rodam na própria pasta, o que pode acabar gerando erros de execução caso
		//não encontrem o arquivo.env
		if err := godotenv.Load(strTopPath + strEnvFile); err != nil {

			//Caso falhe, tenta localizar a pasta principal do projeto, e verificar diretamente a partir da mesma
			//Pega o diretório atual da aplicação/ diretório que estou rodando a aplicação
			err := LoadSysEnv()

			if err != nil {
				return err
			}
		}

	} else {
		if err := godotenv.Load(strEnvFile); err != nil {
			return ErrEnvNotFound
		}
	}

	return nil
}

// Procura/carrega o arquivo '.env', a partir do diretório mais acima da aplicação
func LoadSysEnv() error {
	rootPath, err := GetRootPath()

	if err != nil {
		return err
	}

	//Monta a pesquisa pelo diretório, e tenta fazer o load do arquivo '.env'
	if err := godotenv.Load(filepath.Join(rootPath, strEnvFile)); err != nil {
		// Não localizou o '.env' - verificar os diretórios, e voltar até o diretório que contém um identificador
		// que é considerado como diretório raiz da aplicação, para tentar localizar o '.env'
		envFilePath := GetSysPath(rootPath)

		if err := godotenv.Load(envFilePath); err != nil {
			return ErrEnvNotFound
		}
	}

	return nil
}

// Retorna o nível de diretório da aplicação que está rodando
func GetRootPath() (string, error) {

	var rootPath string = ""

	strPathApp, err := os.Getwd()

	if err != nil {
		return rootPath, ErrEnvPathLoad
	}

	// Retorna até o diretório da aplicação
	rootPath, err = filepath.Abs(filepath.Dir(strPathApp))

	if err != nil {
		return rootPath, ErrEnvRootPath
	}

	return rootPath, nil
}

// Retorna o nível de diretório mais alto a partir do caminho recebido
func GetSysPath(strPath string) string {

	var topDir string = ""

	for {
		// Verificar se o diretório atual contém um arquivo ou diretório que identifica a raiz do projeto (exemplo: go.mod)
		if _, err := os.Stat(filepath.Join(strPath, strIdeRoot)); err == nil {
			topDir = filepath.Join(strPath, strEnvFile)
			break
		}

		//Subir um nível de diretório
		topPath := filepath.Dir(strPath)
		if topPath == strPath {
			// Está no limite de pastas e não localizou o arquivo '.env'
			break
		}

		strPath = topPath
	}

	return topDir
}
