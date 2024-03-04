package main

import (
	"fmt"
	"log"

	"github.com/ffalconesmera/bank-simulator/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("initializing bank simulator..")

	g := gin.Default()

	routes_bank := g.Group("/bank_simulator/api/v1")
	{
		bankController := controller.NewBankController()
		routes_merchants := routes_bank.Group("/")
		routes_merchants.POST("/payments", bankController.ProcessPayment)
		routes_merchants.POST("/refunds", bankController.RefundPayment)
	}

	g.Run(":8083")

	log.Println("BANK SIMUlATOR: Running..!!")
}
