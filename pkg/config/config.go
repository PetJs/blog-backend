package config

import "os"

type Config struct {
	Port string
	DB_URL string
}


func LoadConfig() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DB_URL: getEnv("DB_URL", "root:Morowa@tcp(127.0.0.1:3306)/blogdb?charset=utf8mb4&parseTime=True&loc=Local"),
	}
}

func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}