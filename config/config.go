package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort int
	ServerHost string
	JWTSecret  string
}

func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "microservice_db"),
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		JWTSecret:  getEnv("JWT_SECRET", "dev-secret-key"),
	}

	var err error
	cfg.DBPort, err = getEnvInt("DB_PORT", 5432)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	cfg.ServerPort, err = getEnvInt("SERVER_PORT", 8080)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %w", err)
	}

	return cfg, nil
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) (int, error) {
	valStr := getEnv(key, "")
	if valStr == "" {
		return defaultVal, nil
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}