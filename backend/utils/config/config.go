package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgreDatabase DatabaseConfig
	MongoDatabase   string
	DBType          string
	Port            string
	JwtSecret       string
	Redis           RedisConfig
}

type DatabaseConfig struct {
	Host           string
	DBPort         string
	User           string
	Password       string
	Name           string
	SSLMode        string
	ChannelBinding string
}

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error on loading env")
	}

	cfg := &Config{
		PostgreDatabase: DatabaseConfig{
			Host:     getEnv("HOST", "localhost"),
			DBPort:   getEnv("DB_PORT", "8000"),
			User:     getEnv("USER", "postgres"),
			Password: getEnv("PASSWORD", ""),
			Name:     getEnv("NAME", "socialmedia"),
			SSLMode:  getEnv("SSLMODE", "disable"),
			// ChannelBinding: getEnv("CHANNEL_BINDING", "disable"),
		},
		MongoDatabase: getEnv("MONGO_CONNECTION_STRING", ""),
		DBType:        getEnv("DB_TYPE", ""),
		Port:          getEnv("PORT", "8000"),
		JwtSecret:     getEnv("JWT_SECRET", ""),
		Redis: RedisConfig{
			RedisHost:     getEnv("REDIS_HOST", "localhost"),
			RedisPort:     getEnv("REDIS_PORT", "6379"),
			RedisPassword: getEnv("REDIS_PASSWORD", ""),
			RedisDB:       getEnvInt("REDIS_DB", 0),
		},
	}

	if cfg.JwtSecret == "" {
		return nil, errors.New("jwt secret is null")
	}

	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Print("Error on loading env INT")
			return defaultVal
		}

		return i
	}

	return defaultVal
}
