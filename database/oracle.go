package database

/*
func InitOracle(cfgParameters *config.DBConfig) (*gorm.DB, error) {

	strDSN := cfgParameters.Username + "/"
	strDSN += cfgParameters.Password
	strDSN += "@" + cfgParameters.Host + ":" //"@localhost:1521/"
	strDSN += cfgParameters.Port + "/"
	strDSN += cfgParameters.DBName

	portNumber, err := strconv.Atoi(cfgParameters.Port)

	if err != nil {
		return nil, err
	}

	connStr := go_ora.BuildUrl(cfgParameters.Host,
		portNumber,
		cfgParameters.DBName,
		cfgParameters.Username, cfgParameters.Password, nil)

	//db, err := gorm.Open(go_ora.Open(strDSN), &gorm.Config{})
	db := connectionDB.Statement.ReflectValue.IsNil()
	db, err = sql.Open("oracle", connStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}
*/
