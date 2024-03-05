package main

import (
	"context"
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
	ctx := context.Background()

	log.Println("read enviroment config..")
	config.Config().InitConfig()

	log.Println("set up http logger..")
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db := database.NewDatabaseConnection()
	db.InitDatabase(config.Config().GetDatabaseHost(), config.Config().GetDatabasePort(), config.Config().GetDatabaseName(), config.Config().GetDatabaseUser(), config.Config().GetDatabasePassword())
	defer db.Close()

	log.Println("setting repository layer..")
	merchantRepo := repository.NewMerchantRepository(db)
	userRepo := repository.NewUserRepository(db)

	log.Println("setting service layer..")
	merchantService := service.NewMerchantService(userRepo, merchantRepo)

	log.Println("setting controller layer..")
	userController := controller.NewMerchantController(ctx, merchantService)

	log.Println("init routes and listening..")
	routes.InitApiRouter(userController).ListenAndServe(ctx)
}
