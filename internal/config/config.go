package config

type Config struct {
	Addr    string
	DSN     string // Data Source Name for database connection
	BaseURL string // Base URL for the application
	Port    string // Port for the application
}
