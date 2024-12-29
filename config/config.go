package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EthereumRPCURL string
	Port           string
}

var GlobalConfig *Config

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	config := &Config{
		EthereumRPCURL: getEnv("ETHEREUM_RPC_URL", "http://127.0.0.1:8545"),
		Port:           getEnv("PORT", "8080"),
	}

	if config.EthereumRPCURL == "" {
		log.Fatal("ETHEREUM_RPC_URL is required but not set")
	}

	GlobalConfig = config
	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
