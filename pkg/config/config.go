package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port string
	DB_URL string
	JWT_SECRET_KEY string
}


func LoadConfig() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		DB_URL: getEnv("DB_URL", "root:Morowa@tcp(127.0.0.1:3306)/blogdb?charset=utf8mb4&parseTime=True&loc=Local"),
		JWT_SECRET_KEY: getEnv("JWT_SECRET_KEY", "SHABALALA"),
	}
}

func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("The env variable is: ", value)
		return value
	}
	return fallback
}