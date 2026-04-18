package config

import "os"

type Config struct {
	Port              string
	DBURL             string
	JWTSecretKey      string
	CloudinaryURL     string
	OpenAIAPIKey      string
	ElevenLabsAPIKey  string
	ElevenLabsVoiceID string
}

func LoadConfig() Config {
	return Config{
		Port:              getEnv("PORT", "8080"),
		DBURL:             getEnv("DB_URL", ""),
		JWTSecretKey:      getEnv("JWT_SECRET_KEY", ""),
		CloudinaryURL:     getEnv("CLOUDINARY_URL", ""),
		OpenAIAPIKey:      getEnv("OPENAI_API_KEY", ""),
		ElevenLabsAPIKey:  getEnv("ELEVENLABS_API_KEY", ""),
		ElevenLabsVoiceID: getEnv("ELEVENLABS_VOICE_ID", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
