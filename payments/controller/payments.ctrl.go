package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	ext_dto "github.com/ffalconesmera/payments-platform/payments/externals/dto"
	"github.com/ffalconesmera/payments-platform/payments/helpers"
	"github.com/ffalconesmera/payments-platform/payments/model"
	"github.com/ffalconesmera/payments-platform/payments/model/dto"
	"github.com/ffalconesmera/payments-platform/payments/service"
	"github.com/gin-gonic/gin"
)

// PaymentController is an interface to comunicate with service layer and define context for identify each request
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

// CheckoutPayment: execute CheckoutPayment defined in service
func (cp paymentControllerImpl) CheckoutPayment(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())
	merchantCode := c.Params.ByName("merchant_code")

	var payment dto.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		helpers.ResponseJson(c, nil, err)
		return
	}

	if helpers.InvalidFloat(payment.Amount) {
		helpers.ResponseJson(c, nil, errors.New("amount could not be zero"))
		return
	}

	if helpers.EmptyString(payment.Description) {
		helpers.ResponseJson(c, nil, errors.New("description could not be empty"))
		return
	}

	_, okCurrency := model.Currency[payment.Currency]

	if !okCurrency {
		helpers.ResponseJson(c, nil, fmt.Errorf("currency: is not a valid value. %v", model.Currency))
		return
	}

	if helpers.EmptyString(payment.Customer.DNI) {
		helpers.ResponseJson(c, nil, fmt.Errorf("customer dni could not be empty"))
		return
	}

	if helpers.EmptyString(payment.Customer.Name) {
		helpers.ResponseJson(c, nil, fmt.Errorf("customer name could not be empty"))
		return
	}

	if helpers.EmptyString(payment.Customer.Email) {
		helpers.ResponseJson(c, nil, fmt.Errorf("customer email could not be empty"))
		return
	}

	if helpers.EmptyString(payment.Customer.Phone) {
		helpers.ResponseJson(c, nil, fmt.Errorf("customer phone could not be empty"))
		return
	}
	pay, err := cp.paymentService.CheckoutPayment(c, merchantCode, payment)
	helpers.ResponseJson(c, pay, err)
}

// ProcessPayment: execute ProcessPayment defined in service
func (cp paymentControllerImpl) ProcessPayment(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())
	paymentCode := c.Params.ByName("payment_code")

	var cardInfoDto *ext_dto.BankCardInput
	if err := c.ShouldBindJSON(&cardInfoDto); err != nil {
		helpers.ResponseJson(c, nil, err)
		return
	}

	cardInfoDto.PaymentReference = paymentCode

	cardInfo, err := json.Marshal(cardInfoDto)
	if err != nil {
		helpers.ResponseJson(c, nil, err)
		return
	}

	resp, err := cp.paymentService.ProcessPayment(c, paymentCode, string(cardInfo))
	helpers.ResponseJson(c, resp, err)
}

// RefundPayment: execute RefundPayment defined in service
func (cp paymentControllerImpl) RefundPayment(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())
	paymentCode := c.Params.ByName("payment_code")

	var refundDto *ext_dto.RefundInput
	if err := c.ShouldBindJSON(&refundDto); err != nil {
		helpers.ResponseJson(c, nil, err)
	}

	refundDto.PaymentReference = paymentCode

	refundInfo, err := json.Marshal(refundDto)
	if err != nil {
		helpers.ResponseJson(c, nil, err)
		return
	}
	ref, err := cp.paymentService.RefundPayment(c, paymentCode, string(refundInfo))
	helpers.ResponseJson(c, ref, err)
}

// CheckPayment: execute CheckPayment defined in service
func (cp paymentControllerImpl) CheckPayment(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())
	paymentCode := c.Params.ByName("payment_code")
	pay, err := cp.paymentService.CheckPayment(c, paymentCode, c.Value("MERCHANT_CODE").(string))
	helpers.ResponseJson(c, pay, err)
}
