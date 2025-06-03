package config

type DBConfig struct {
	DBType   string
	Driver   string
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	DBApiKey string
	SSLMode  string
}
