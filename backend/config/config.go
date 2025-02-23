package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Github   GithubConfig
	Server   ServerConfig
	Proxy    ProxyConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type GithubConfig struct {
	ClientID     string
	ClientSecret string
	RepoOwner    string
	RepoName     string
}

type ServerConfig struct {
	Port string
}

type ProxyConfig struct {
	Enable bool
	URL    string
}

var AppConfig Config

func Init() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	AppConfig = Config{
		Database: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "3306"),
			User:     getEnvOrDefault("DB_USER", "root"),
			Password: getEnvOrDefault("DB_PASSWORD", ""),
			Name:     getEnvOrDefault("DB_NAME", "note_sync"),
		},
		Github: GithubConfig{
			ClientID:     getEnvOrDefault("GITHUB_CLIENT_ID", ""),
			ClientSecret: getEnvOrDefault("GITHUB_CLIENT_SECRET", ""),
			RepoOwner:    getEnvOrDefault("GITHUB_REPO_OWNER", ""),
			RepoName:     getEnvOrDefault("GITHUB_REPO_NAME", ""),
		},
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
		},
		Proxy: ProxyConfig{
			Enable: getEnvOrDefault("PROXY_ENABLE", "true") == "true",
			URL:    getEnvOrDefault("PROXY_URL", "socks5://127.0.0.1:7989"),
		},
	}

	validateConfig()

	// 初始化OAuth配置
	InitOAuth()
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func validateConfig() {
	if AppConfig.Github.ClientID == "" || AppConfig.Github.ClientSecret == "" {
		log.Fatal("GitHub OAuth credentials are required")
	}

	if AppConfig.Github.RepoOwner == "" || AppConfig.Github.RepoName == "" {
		log.Fatal("GitHub repository configuration is required")
	}

	log.Println("Configuration loaded successfully")
}

func GetMySQLDSN() string {
	db := AppConfig.Database
	return db.User + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" +
		db.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
}
