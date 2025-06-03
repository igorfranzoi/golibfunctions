package models

import "reflect"

var migrates []reflect.Type

type MigrateServiceRepository interface {
	MigrateApply() error
	MigrateRevert() error
	MigrateName() string
	MigratePremises() []MigrateServiceRepository
}

func AddMigrate(migrate reflect.Type) {
	migrates = append(migrates, migrate)
}

func GetAllMigrates() []reflect.Type {
	return migrates
}
