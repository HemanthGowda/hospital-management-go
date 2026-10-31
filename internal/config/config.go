package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port              string
	DatabaseURL       string
	JWTSecret         string
	JWTExpiration     time.Duration
	JWTIssuer         string
	SeedAdminEmail    string
	SeedAdminPassword string
	SeedAdminFullName string
}

func Load() (Config, error) {
	expirationMS, err := strconv.ParseInt(getenv("JWT_EXPIRATION_MS", "3600000"), 10, 64)
	if err != nil || expirationMS <= 0 {
		return Config{}, fmt.Errorf("JWT_EXPIRATION_MS must be a positive integer")
	}

	secret := getenv("JWT_SECRET", "change-me-to-a-long-random-secret-at-least-32-chars")
	if len(secret) < 32 {
		return Config{}, fmt.Errorf("JWT_SECRET must be at least 32 characters")
	}

	return Config{
		Port:              getenv("SERVER_PORT", "8080"),
		DatabaseURL:       getenv("DATABASE_URL", "postgres://hospital:hospital@localhost:5433/hospital_management?sslmode=disable"),
		JWTSecret:         secret,
		JWTExpiration:     time.Duration(expirationMS) * time.Millisecond,
		JWTIssuer:         getenv("JWT_ISSUER", "hospital-management"),
		SeedAdminEmail:    getenv("SEED_ADMIN_EMAIL", "admin@hospital.local"),
		SeedAdminPassword: getenv("SEED_ADMIN_PASSWORD", "ChangeMe123!"),
		SeedAdminFullName: getenv("SEED_ADMIN_FULL_NAME", "Hospital Admin"),
	}, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
