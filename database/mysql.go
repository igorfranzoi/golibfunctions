package database

import (
	"strconv"

	"github.com/igorfranzoi/golibfunctions/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(cfgParameters *config.DBConfig) (*gorm.DB, error) {

	strDSN := cfgParameters.Username + ":"
	strDSN += cfgParameters.Password + "@tcp(" + cfgParameters.Host + ":" + strconv.Itoa(cfgParameters.Port) + ")/"
	strDSN += cfgParameters.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(strDSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
