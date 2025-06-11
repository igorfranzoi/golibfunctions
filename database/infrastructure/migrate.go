package infrastructure

import (
	"fmt"

	"reflect"

	"github.com/igorfranzoi/golibfunctions/database"
	"github.com/igorfranzoi/golibfunctions/database/models"
)

func RunningMigrate(constructMethod string) {

	migrates := models.GetAllMigrates()

	ifaceMigrate := reflect.TypeOf((*models.MigrateServiceRepository)(nil)).Elem()

	migrateMap := make(map[string]int)

	for _, migrateType := range migrates {
		fmt.Printf("processando migration - %s\n", migrateType)

		migrate := reflect.New(migrateType).Elem()

		if !migrate.Addr().Type().Implements(ifaceMigrate) {
			fmt.Println("interface não implementada ....")
			continue
		}

		field := migrate.FieldByName("Db")
		field.Set(reflect.ValueOf(database.GetGormDB()))

		fmt.Printf("iniciando aplicação 'migration': %s\n", migrateType)

		migrateProcess(migrate.Addr(), constructMethod, migrateMap)
	}
}

func migrateProcess(migrate reflect.Value, methodExecute string, migrateMap map[string]int) {
	fmt.Printf("método: %s\n", methodExecute)

	methodMigrateName := migrate.MethodByName("MigrateName")

	if !methodMigrateName.IsValid() {
		fmt.Printf("método %s não localizado\n", methodExecute)
		return
	}

	migrateName := methodMigrateName.Call(nil)[0].String()

	if _, ok := migrateMap[migrateName]; ok {
		return
	}

	methodMigratePremises := migrate.MethodByName("MigratePremises")

	if !methodMigratePremises.IsValid() {
		return
	}

	premises := methodMigratePremises.Call(nil)[0]

	for i := 0; i < premises.Len(); i++ {
		migrateProcess(premises.Index(i), methodExecute, migrateMap)
	}

	method := migrate.MethodByName(methodExecute)

	if !method.IsValid() {
		return
	}

	method.Call(nil)

	migrateMap[migrateName] = 1
}
