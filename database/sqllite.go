package database

import (
	"github.com/igorfranzoi/golibfunctions/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite(cfg *config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
