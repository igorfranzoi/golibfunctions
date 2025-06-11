package validation

import (
	"errors"
	"os"
	"strconv"

	"github.com/igorfranzoi/golibfunctions/config"
)

var (
	ErrEnvParNotFound   = errors.New("alguma variável do database não foi configurada")
	ErrEnvParPortNotInt = errors.New("o número da porta de conexão não é um inteiro")
	ErrInvalidDBType    = errors.New("DB_TYPE - não é válido, verifique o arquivo '.env'")
)

// Carrega as variáveis de ambiente e retorna um mapa de configurações
func InitDatabaseVars() (*config.DBConfig, error) {

	var retPar config.DBConfig

	envVars := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_USERNAME": os.Getenv("DB_USERNAME"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_DRIVER":   os.Getenv("DB_DRIVER"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_SSLMODE":  os.Getenv("DB_SSLMODE"),
		"DB_API_KEY":  os.Getenv("DB_API_KEY"),
	}

	err := validateEnvVariables(envVars)

	if err != nil {
		return nil, err
	}

	retPar.Driver = envVars["DB_DRIVER"]
	retPar.Username = envVars["DB_USERNAME"]
	retPar.Password = envVars["DB_PASSWORD"]
	retPar.Host = envVars["DB_HOST"]
	retPar.Port, _ = strconv.Atoi(envVars["DB_PORT"])
	retPar.DBName = envVars["DB_NAME"]
	retPar.SSLMode = envVars["DB_SSLMODE"]
	retPar.DBApiKey = envVars["DB_API_KEY"]

	return &retPar, nil
}

// Mapeia os tipos de bancos de dados para as variáveis necessárias
func requiredVarsForDBType(dbType string) ([]string, error) {
	switch dbType {
	case "postgres", "mysql", "mssql", "oracle":
		return []string{"DB_HOST", "DB_PORT", "DB_USERNAME", "DB_PASSWORD", "DB_NAME"}, nil
	case "sqlite":
		return []string{"DB_NAME"}, nil
	case "firebase":
		return []string{"DB_API_KEY"}, nil
	default:
		return nil, ErrInvalidDBType
	}
}

func validateEnvVariables(varsEnv map[string]string) error {
	dbType := varsEnv["DB_DRIVER"]

	requiredVars, err := requiredVarsForDBType(dbType)

	if err != nil {
		return err
	}

	// Verificar se todas as variáveis obrigatórias estão presentes
	for _, varName := range requiredVars {
		if varsEnv[varName] == "" {
			return ErrEnvParNotFound
		}
	}

	// Validar se a porta, é um número inteiro
	if _, err := strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		return ErrEnvParPortNotInt
	}

	return nil
}
