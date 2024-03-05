package main

import (
	"log"

	"github.com/ffalconesmera/payments-platform/payments/config"
	"github.com/ffalconesmera/payments-platform/payments/controller"
	"github.com/ffalconesmera/payments-platform/payments/database"
	external_repository "github.com/ffalconesmera/payments-platform/payments/externals/repository"
	"github.com/ffalconesmera/payments-platform/payments/repository"
	"github.com/ffalconesmera/payments-platform/payments/routes"
	"github.com/ffalconesmera/payments-platform/payments/service"
	"github.com/sirupsen/logrus"
)

func main() {
	log.Println("initializating merchant api..")

	log.Println("reading variables environment..")
	config.InitConfig()

	log.Println("set up http logger..")
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	db := database.NewDatabaseConnection(config.GetDatabaseHost(), config.GetDatabasePort(), config.GetDatabaseName(), config.GetDatabaseUser(), config.GetDatabasePassword())

	log.Println("setting repository layer..")
	customerRepo := repository.NewCustomerRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	refundRepo := repository.NewRefundRepository(db)
	bankRepo := external_repository.NewBankRepository()
	merchantRepo := external_repository.NewMerchantRepository()

	log.Println("setting service layer..")
	paymentService := service.NewPaymentService(&paymentRepo, &refundRepo, &customerRepo, &bankRepo, &merchantRepo)

	log.Println("setting controller layer..")
	paymentController := controller.NewPaymentController(&paymentService)

	log.Println("init routes and listening..")
	routes.InitApiRouter(&paymentController).ListenAndServe()
}
