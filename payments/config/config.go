package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config is an singleton for get environment vars
func InitConfig() {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Sprintf("failed to load .env file: %s", err.Error()))
	}

	log.Println(".env load successful")
}

func GetDatabaseHost() string {
	return os.Getenv("DB_HOST")
}

func GetDatabasePort() string {
	return os.Getenv("DB_PORT")
}

func GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}

func GetDatabaseUser() string {
	return os.Getenv("DB_USER")
}

func GetDatabasePassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetJWTSecretKey() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("jwt secret key is undefined")
	}

	return secret
}

func GetPaymentExpiration() int64 {
	minutes, err := strconv.ParseInt(os.Getenv("PAYMENT_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		minutes = 5
	}

	return minutes
}

func GetBankEndpoint() string {
	return os.Getenv("BANK_SIMULATOR_ENDPOINT")
}

func GetMerchantEndpoint() string {
	return os.Getenv("MERCHANTS_ENDPOINT")
}
