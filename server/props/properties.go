package props

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	AppPort string `yaml:"app_port"`
	AppHost string `yaml:"app_host"`
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DBName    string `yaml:"dbname"`
	SSLMode   string `yaml:"sslmode"`
	IsMigrate bool   `yaml:"is_migrate"`
}
