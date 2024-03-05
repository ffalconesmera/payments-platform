package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config is an singleton for get environment vars
type configImpl struct{}

var conf *configImpl

func Config() *configImpl {
	if conf == nil {
		conf = &configImpl{}
	}

	return conf
}

func (c *configImpl) InitConfig() {
	err := godotenv.Load()

	if err != nil {
		panic(fmt.Sprintf("failed to load .env file: %s", err.Error()))
	}

	log.Println(".env load successful")
}

func (c *configImpl) GetDatabaseHost() string {
	return os.Getenv("DB_HOST")
}

func (c *configImpl) GetDatabasePort() string {
	return os.Getenv("DB_PORT")
}

func (c *configImpl) GetDatabaseName() string {
	return os.Getenv("DB_NAME")
}

func (c *configImpl) GetDatabaseUser() string {
	return os.Getenv("DB_USER")
}

func (c *configImpl) GetDatabasePassword() string {
	return os.Getenv("DB_PASSWORD")
}

func (c *configImpl) GetJWTSecretKey() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("jwt secret key is undefined")
	}

	return secret
}

func (c *configImpl) GetPaymentExpiration() int64 {
	minutes, err := strconv.ParseInt(os.Getenv("PAYMENT_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		minutes = 5
	}

	return minutes
}

func (c *configImpl) GetBankEndpoint() string {
	return os.Getenv("BANK_SIMULATOR_ENDPOINT")
}

func (c *configImpl) GetMerchantEndpoint() string {
	return os.Getenv("MERCHANTS_ENDPOINT")
}