package controller

import (
	"context"

	"github.com/ffalconesmera/payments-platform/payments/helpers"
	"github.com/ffalconesmera/payments-platform/payments/service"
	"github.com/gin-gonic/gin"
)

// PaymentController is an interface to comunicate with service layer and define context for identify each request
// CheckoutPayment: execute CheckoutPayment defined in service
// ProcessPayment: execute ProcessPayment defined in service
// RefundPayment: execute RefundPayment defined in service
// CheckPayment: execute CheckPayment defined in service
type PaymentController interface {
	CheckoutPayment(c *gin.Context)
	ProcessPayment(c *gin.Context)
	RefundPayment(c *gin.Context)
	CheckPayment(c *gin.Context)
}

type paymentControllerImpl struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService *service.PaymentService) PaymentController {
	if paymentService == nil {
		return nil
	}

	return &paymentControllerImpl{
		paymentService: *paymentService,
	}
}

func (p paymentControllerImpl) CheckoutPayment(c *gin.Context) {
	checkoutCtx := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	p.paymentService.CheckoutPayment(checkoutCtx, c)
}

func (p paymentControllerImpl) ProcessPayment(c *gin.Context) {
	processCtx := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	p.paymentService.ProcessPayment(processCtx, c)
}

func (p paymentControllerImpl) RefundPayment(c *gin.Context) {
	refundCtx := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	p.paymentService.RefundPayment(refundCtx, c)
}

func (p paymentControllerImpl) CheckPayment(c *gin.Context) {
	checkPaymentCtx := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	p.paymentService.CheckPayment(checkPaymentCtx, c)
}
