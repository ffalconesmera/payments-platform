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
	GetJWTExpiration() int64
	GetJWTSecretKey() string
}

type configImpl struct {
}

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
