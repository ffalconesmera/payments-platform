package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config is an interface for get environment vars
type Config interface {
	InitConfig()
	GetDatabaseHost() string
	GetDatabasePort() string
	GetDatabaseName() string
	GetDatabaseUser() string
	GetDatabasePassword() string
	GetJWTExpiration() int64
	GetJWTSecretKey() string
}

type configImpl struct{}

func NewConfig() *configImpl {
	return &configImpl{}
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

func (c *configImpl) GetJWTExpiration() int64 {
	minutes, err := strconv.ParseInt(os.Getenv("JWT_EXPIRATION_MINUTES"), 10, 64)
	if err != nil {
		minutes = 60
	}

	return minutes
}

func (c *configImpl) GetJWTSecretKey() string {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("jwt secret key is undefined")
	}

	return secret
}
