package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ffalconesmera/payments-platform/merchants/config"
	"github.com/ffalconesmera/payments-platform/merchants/controller"
	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/helpers"
	"github.com/ffalconesmera/payments-platform/merchants/repository"
	"github.com/ffalconesmera/payments-platform/merchants/routes"
	"github.com/ffalconesmera/payments-platform/merchants/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("initializating merchant api..")
	ctx := context.Background()

	log.Println("reading variables environment..")
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load .env file: %s", err.Error()))
	}

	log.Println("set up http logger..")
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	log.Println("read enviroment config..")
	config := config.NewConfig()

	db := database.NewDatabaseConnection()
	db.InitDatabase(config.GetDatabaseHost(), config.GetDatabasePort(), config.GetDatabaseName(), config.GetDatabaseUser(), config.GetDatabasePassword())
	defer db.Close()

	log.Println("declare customs has generator object")
	customHash := helpers.NewCustomHash(config)

	log.Println("declare customs logs generator object")
	customLog := helpers.NewCustomLog(customHash)

	log.Println("declare customs validations object")
	customValidation := helpers.NewCustomValidation()

	log.Println("setting repository layer..")
	merchantRepo := repository.NewMerchantRepository(db)
	userRepo := repository.NewUserRepository(db)

	log.Println("setting service layer..")
	merchantService := service.NewMerchantService(customLog, customValidation, customHash, userRepo, merchantRepo)

	log.Println("setting controller layer..")
	userController := controller.NewMerchantController(ctx, merchantService, customHash)

	log.Println("init routes and listening..")
	routes.InitApiRouter(userController).ListenAndServe(ctx)
}
