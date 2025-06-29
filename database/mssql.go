package database

import (
	"strconv"

	"github.com/igorfranzoi/golibfunctions/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitMSSQL(cfgParameters *config.DBConfig) (*gorm.DB, error) {
	strDSN := "sqlserver://" + cfgParameters.Username
	strDSN += ":" + cfgParameters.Password
	strDSN += "@" + cfgParameters.Host
	strDSN += ":" + strconv.Itoa(cfgParameters.Port)
	strDSN += "?database=" + cfgParameters.DBName

	db, err := gorm.Open(sqlserver.Open(strDSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
