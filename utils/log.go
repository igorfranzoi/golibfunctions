package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	InfoLevel = iota
	WarnLevel
	ErrorLevel
	FatalLevel
)

var logUseFile bool = true
var logUsePath string = "logs" + "/"
var logUseName string = "logapp"

var (
	MsgErrVldPath    = "Erro ao verificar existência do diretório:"
	MsgErrCreatePath = "Erro ao tentar criar o diretório"
)

func CreateLog() (*zerolog.Logger, bool) {

	var retFun bool = false
	var logger *zerolog.Logger

	if os.Getenv("USE_FILE") != "" {
		logUseFile = (os.Getenv("USE_FILE") == "true")
	}

	if os.Getenv("LOG_PATH") != "" {
		logUsePath = os.Getenv("LOG_PATH")
	}

	if os.Getenv("LOG_NAME") != "" {
		logUseName = os.Getenv("LOG_NAME")
	}

	// Criar um diretório com permissões padrão (0777 no Unix)
	if _, err := os.Stat(logUsePath); os.IsNotExist(err) {
		// Criar um diretório com permissões padrão (0777 no Unix)
		if err := os.Mkdir(logUsePath, 0755); err != nil {
			log.Warn().Str("Error:", err.Error()).Msg(MsgErrCreatePath)

			// Retorna  se não for possível criar o diretório
			return nil, retFun
		}
	} else if err != nil {
		log.Warn().Str("Erro:", err.Error()).Msg(MsgErrVldPath)

		return nil, retFun
	}

	if logUseFile {

		// Obter a data atual
		dateToday := time.Now()

		monthString := fmt.Sprintf("%02d", dateToday.Month())
		dayString := fmt.Sprintf("%02d", dateToday.Day())

		logUseName += "_" + strconv.Itoa(dateToday.Year()) + "_" + monthString + "_" + dayString

		logDir := logUsePath + logUseName + ".log"

		// Abrir ou criar o arquivo de log
		fileLog, err := os.OpenFile(logDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal().Err(err).Msg("Falha ao abrir/criar o arquivo de log")

			return nil, retFun
		}

		//defer fileLog.Close()

		// Configurar zerolog para escrever no arquivo
		newLogger := zerolog.New(fileLog).With().Timestamp().Logger()

		logger = &newLogger

		//Define o output geral para também utilizar o zerolog
		//log.SetOutput(logger)
		logger.Info().Msg("iniciando aplicação")

		retFun = true

		/*
			// Exemplo de logging
			log.Info().Msg("Mensagem de informação")
			log.Warn().Str("animal", "gato").Int("size", 10).Msg("Um animal apareceu")
			log.Error().Str("animal", "leão").Int("size", 100).Msg("Animal perigoso")
			log.Fatal().Str("animal", "tigre").Int("size", 80).Msg("Animal fatal encontrado")
		*/
	}

	return logger, retFun
}

// Escreve uma mensagem no log com base no nível fornecido
func WriteLog(logger *zerolog.Logger, level int, message string, args ...interface{}) {

	if logger == nil {
		log.Error().Msg("Logger não inicializado. Certifique-se de chamar CreateLog.")
		return
	}

	switch level {
	case InfoLevel:
		logger.Info().Msgf(message, args...)
	case WarnLevel:
		logger.Warn().Msgf(message, args...)
	case ErrorLevel:
		logger.Error().Msgf(message, args...)
	case FatalLevel:
		logger.Fatal().Msgf(message, args...)
	default:
		logger.Info().Msgf("Nível de log desconhecido: %d. Mensagem: %s", level, message)
	}
}
