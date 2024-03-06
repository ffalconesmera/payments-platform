package main

import (
	"log"

	"github.com/ffalconesmera/payments-platform/merchants/config"
	"github.com/ffalconesmera/payments-platform/merchants/controller"
	"github.com/ffalconesmera/payments-platform/merchants/database"
	"github.com/ffalconesmera/payments-platform/merchants/repository"
	"github.com/ffalconesmera/payments-platform/merchants/routes"
	"github.com/ffalconesmera/payments-platform/merchants/service"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("initializating merchant api..")

	log.Println("read enviroment config..")
	config.InitConfig()

	log.Println("set up http logger..")
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db := database.NewDatabaseConnection(config.GetDatabaseHost(), config.GetDatabasePort(), config.GetDatabaseName(), config.GetDatabaseUser(), config.GetDatabasePassword())

	log.Println("setting repository layer..")
	merchantRepo := repository.NewMerchantRepository(db)
	userRepo := repository.NewUserRepository(db)

	log.Println("setting service layer..")
	merchantService := service.NewMerchantService(userRepo, merchantRepo)

	log.Println("setting controller layer..")
	merchantController := controller.NewMerchantController(&merchantService)

	log.Println("init routes and listening..")
	routes.InitApiRouter(&merchantController).ListenAndServe()
}
