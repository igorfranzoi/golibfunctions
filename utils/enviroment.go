package utils

import (
	"log"

	"github.com/igorfranzoi/golibfunctions/config/validation"
	"github.com/igorfranzoi/golibfunctions/database"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func InitEnviroment() (*zerolog.Logger, *gorm.DB, error) {

	var err error

	// Disponibiliza o arquivo de log do aplicativo
	logger, valid := CreateLog()

	if !valid {
		log.Fatal("Error ocurrer in log instance")

		return nil, nil, err
	}

	err = LoadEnvMem()

	if err != nil {
		WriteLog(logger, InfoLevel, "Error loading environment variables", err)

		return logger, nil, err
	}

	cfgStruct, err := validation.InitDatabaseVars()

	if err != nil {
		WriteLog(logger, ErrorLevel, "Error loading database variables", err)

		return logger, nil, err
	}

	connDB, err := database.ConnectionDatabase(cfgStruct)

	if err != nil {
		WriteLog(logger, ErrorLevel, "Error initializing database", err)

		return logger, nil, err
	}

	/*gormRepo, err := database.GormRepositoryInstance(connDB)

	if err != nil {
		WriteLog(logger, ErrorLevel, "Error initializing database", err)

		return logger, nil, nil, err
	}*/

	return logger, connDB, err
}
