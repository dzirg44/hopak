package config

// DBConfig database configuration structure
type DBConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// GetConfig returns list of DB configuration
func GetConfig() *DBConfig {
	return &DBConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "root",
		DBName:   "hopak",
	}
}
