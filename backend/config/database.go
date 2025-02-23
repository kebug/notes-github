package config

var (
	DBHost     = getEnvOrDefault("DB_HOST", "localhost")
	DBPort     = getEnvOrDefault("DB_PORT", "3306")
	DBUser     = getEnvOrDefault("DB_USER", "root")
	DBPassword = getEnvOrDefault("DB_PASSWORD", "password")
	DBName     = getEnvOrDefault("DB_NAME", "note_sync")
)
