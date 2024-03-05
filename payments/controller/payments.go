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
	CheckoutPayment(ctx context.Context, c *gin.Context)
	ProcessPayment(ctx context.Context, c *gin.Context)
	RefundPayment(ctx context.Context, c *gin.Context)
	CheckPayment(ctx context.Context, c *gin.Context)
}

type paymentControllerImpl struct {
	paymentService service.PaymentService
}

func NewPaymentController(ctx context.Context, paymentService service.PaymentService) PaymentController {
	return &paymentControllerImpl{
		paymentService: paymentService,
	}
}

func (p paymentControllerImpl) CheckoutPayment(ctx context.Context, c *gin.Context) {
	checkoutCtx := context.WithValue(ctx, "REQUEST_ID", helpers.CustomHash().NewUUIDString())
	p.paymentService.CheckoutPayment(checkoutCtx, c)
}

func (p paymentControllerImpl) ProcessPayment(ctx context.Context, c *gin.Context) {
	processCtx := context.WithValue(ctx, "REQUEST_ID", helpers.CustomHash().NewUUIDString())
	p.paymentService.ProcessPayment(processCtx, c)
}

func (p paymentControllerImpl) RefundPayment(ctx context.Context, c *gin.Context) {
	refundCtx := context.WithValue(ctx, "REQUEST_ID", helpers.CustomHash().NewUUIDString())
	p.paymentService.RefundPayment(refundCtx, c)
}

func (p paymentControllerImpl) CheckPayment(ctx context.Context, c *gin.Context) {
	checkPaymentCtx := context.WithValue(ctx, "REQUEST_ID", helpers.CustomHash().NewUUIDString())
	p.paymentService.CheckPayment(checkPaymentCtx, c)
}
