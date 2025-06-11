package database

import (
	"errors"

	"golibfunctions/config"

	"gorm.io/gorm"
)

var connectionDB *gorm.DB
var repository *GormRepository

var (
	ErrInvalidDriver = errors.New("invalid - driver SQL")
)

type GormRepository struct {
	Db *gorm.DB
}

func GetGormDB() *gorm.DB {
	return connectionDB
}

func GetGormRepository() *GormRepository {
	return repository
}

func ConnectionDatabase(cfgParameters *config.DBConfig) (*gorm.DB, error) {
	var err error

	switch cfgParameters.Driver {
	case "sqlite":
		connectionDB, err = InitSQLite(cfgParameters)
	case "mysql":
		connectionDB, err = InitMySQL(cfgParameters)
	case "postgres":
		connectionDB, err = InitPostgres(cfgParameters)
	case "mssql":
		connectionDB, err = InitMSSQL(cfgParameters)
	case "oracle":
		//connectionDB, err = InitOracle(cfgParameters)
	case "firebase":
		// Firebase not yet implemented here
		return connectionDB, nil
	default:
		return nil, ErrInvalidDriver
	}
	if err != nil {
		return nil, err
	}

	return connectionDB, nil
}

func GormRepositoryInstance(db *gorm.DB) (*GormRepository, error) {
	var repository GormRepository

	if db == nil {
		return nil, errors.New("database is not valid")
	}

	repository.Db = db

	return &repository, nil
}
