package main

import (
	"log"

	"github.com/ffalconesmera/payments-platform/merchants/config"
	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/helpers"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("initializating merchant api..")
	//ctx := context.Background()

	log.Println("set up http logger..")
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	log.Println("read enviroment config..")
	config := config.NewConfig()

	log.Println("declare database object..")
	database := database.NewDatabaseConnection(config)

	log.Println("initialize database connection")
	database.InitDatabase()

	log.Println("declare customs has generator object")
	customHash := helpers.NewCustomHash(config)

	log.Println("declare customs logs generator object")
	customLog := helpers.NewCustomLog(customHash)

	log.Println("declare customs logs generator object")
	customError := helpers.NewCustomError()
}
