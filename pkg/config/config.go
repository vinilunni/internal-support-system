package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Environment string
	ServerPort  string
	GRPCPort    string

	// Database Configuration
	DatabaseType     string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMode  string
	DatabaseURL      string

	// JWT Configuration
	JWTSecret           string
	JWTExpiresIn        string
	JWTRefreshExpiresIn string

	// OAuth Configuration
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	AllowedDomains     []string

	// Other Services
	RedisURL string

	// CORS Configuration
	CORSAllowedOrigins []string
	CORSAllowedMethods []string
	CORSAllowedHeaders []string
}

func Load() *Config {
	cfg := &Config{
		Environment: getEnv("ENV", "development"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		GRPCPort:    getEnv("GRPC_PORT", "9090"),

		// Database Configuration
		DatabaseType:     getEnv("DB_TYPE", "postgres"),
		DatabaseHost:     getEnv("DB_HOST", "localhost"),
		DatabasePort:     getEnv("DB_PORT", "5432"),
		DatabaseUser:     getEnv("DB_USER", "postgres"),
		DatabasePassword: getEnv("DB_PASSWORD", ""),
		DatabaseName:     getEnv("DB_NAME", "internal_support"),
		DatabaseSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		// JWT Configuration
		JWTSecret:           getEnv("JWT_SECRET", "your-super-secret-key-change-in-production"),
		JWTExpiresIn:        getEnv("JWT_EXPIRES_IN", "24h"),
		JWTRefreshExpiresIn: getEnv("JWT_REFRESH_EXPIRES_IN", "7d"),

		// OAuth Configuration
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "http://localhost:8080/auth/google/callback"),
		AllowedDomains:     getEnvAsSlice("ALLOWED_DOMAINS", []string{"yourcompany.com"}),

		// Other Services
		RedisURL: getEnv("REDIS_URL", "redis://localhost:6379"),

		// CORS Configuration
		CORSAllowedOrigins: getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000"}),
		CORSAllowedMethods: getEnvAsSlice("CORS_ALLOWED_METHODS", []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		CORSAllowedHeaders: getEnvAsSlice("CORS_ALLOWED_HEADERS", []string{"Content-Type", "Authorization"}),
	}

	// Build database URL based on type
	cfg.DatabaseURL = cfg.buildDatabaseURL()

	return cfg
}

func (c *Config) buildDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
		c.DatabaseSSLMode,
	)
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}
	return defaultValue
}
