package config

import "os"

var (
	DBUrl         string
	AuthorizerUrl string
)

func GetEnvOrDefault(envVar, defaultValue string) string {
	if value := os.Getenv(envVar); value != "" {
		return value
	}

	return defaultValue
}

func Load() {
	DBUrl = GetEnvOrDefault("DB_URL", "root:password@tcp(127.0.0.1:3306)/q2_database?parseTime=true")
	AuthorizerUrl = GetEnvOrDefault("AUTHORIZER_URL", "https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
}
