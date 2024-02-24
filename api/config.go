package api

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	Address    string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	JWTSecret  string
}

func LoadConfig() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("PORT not set, defaulting to 8000")
		port = "8000"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return Config{}, errors.New("DB_USER not set")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return Config{}, errors.New("DB_PASSWORD not set")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return Config{}, errors.New("DB_NAME not set")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return Config{}, errors.New("DB_HOST not set")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return Config{}, errors.New("DB_PORT not set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return Config{}, errors.New("JWT_SECRET not set")
	}

	return Config{
		Address:    ":" + port,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBHost:     dbHost,
		DBPort:     dbPort,
		JWTSecret:  jwtSecret,
	}, nil
}
