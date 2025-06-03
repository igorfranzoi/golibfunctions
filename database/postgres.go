package database

import (
	"strconv"

	"github.com/igorfranzoi/base-lib-functions/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(cfgParameters *config.DBConfig) (*gorm.DB, error) {
	strDSN := "host=" + cfgParameters.Host
	strDSN += " user=" + cfgParameters.Username
	strDSN += " password=" + cfgParameters.Password
	strDSN += " dbname=" + cfgParameters.DBName
	strDSN += " port=" + strconv.Itoa(cfgParameters.Port)
	strDSN += " sslmode=" + cfgParameters.SSLMode

	db, err := gorm.Open(postgres.Open(strDSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
