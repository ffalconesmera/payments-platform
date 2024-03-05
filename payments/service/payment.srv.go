package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ffalconesmera/payments-platform/payments/config"
	ext_repository "github.com/ffalconesmera/payments-platform/payments/externals/repository"
	"github.com/ffalconesmera/payments-platform/payments/helpers"
	"github.com/ffalconesmera/payments-platform/payments/model"
	"github.com/ffalconesmera/payments-platform/payments/model/dto"
	"github.com/ffalconesmera/payments-platform/payments/repository"
	"github.com/gin-gonic/gin"
)

// PaymentService is an inteface for process payments

// ProcessPayment: receive card information and get send to bank provider for process charge
// RefundPayment: send to bank an order of refund
// CheckPayment: returns the information of a previous payment
type PaymentService interface {
	CheckoutPayment(ctxt context.Context, c *gin.Context)
	ProcessPayment(ctxt context.Context, c *gin.Context)
	RefundPayment(ctxt context.Context, c *gin.Context)
	CheckPayment(ctxt context.Context, c *gin.Context)
}

type paymentServiceImpl struct {
	paymentRepository  repository.PaymentRepository
	refundRepository   repository.RefundRepository
	customerRepository repository.CustomerRepository
	bankRepository     ext_repository.BankRepository
	merchantRepository ext_repository.MerchantRepository
}

func NewPaymentService(paymentRepository *repository.PaymentRepository, refundRepository *repository.RefundRepository, customerRepository *repository.CustomerRepository, bankRepository *ext_repository.BankRepository, merchantRepository *ext_repository.MerchantRepository) PaymentService {
	return &paymentServiceImpl{
		paymentRepository:  *paymentRepository,
		refundRepository:   *refundRepository,
		customerRepository: *customerRepository,
		bankRepository:     *bankRepository,
		merchantRepository: *merchantRepository,
	}
}

// CheckoutPayment: generate an payment order and return a payment code for future process
func (cp *paymentServiceImpl) CheckoutPayment(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute sing up method..")

	helpers.PrintInfo(ctxt, c, "finding merchant by code..")
	merchantCode := c.Params.ByName("merchant_code")

	merchant, err := cp.merchantRepository.FindMerchantByCode(merchantCode)

	if err != nil {
		helpers.PrintInfo(ctxt, c, fmt.Sprintf("merchant not found %s. Error: %s", merchantCode, err.Error()))
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("merchant not found %s", merchantCode))
		return
	}

	helpers.PrintInfo(ctxt, c, "mapping payment data..")
	var payment *dto.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error mapping payment information. Error: %s", err), false)
		helpers.JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	helpers.PrintInfo(ctxt, c, "validating merchant data..")

	if helpers.InvalidFloat(payment.Amount) {
		helpers.PrintError(ctxt, c, "amount could not be zero", false)
		helpers.JsonFail(c, http.StatusBadRequest, "amount could not be zero")
		return
	}

	if helpers.EmptyString(payment.Description) {
		helpers.PrintError(ctxt, c, "description could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "description could not be empty")
		return
	}

	_, okCurrency := model.Currency[payment.Currency]

	if !okCurrency {
		helpers.PrintError(ctxt, c, fmt.Sprintf("currency: is not a valid value. %v", model.Currency), false)
		helpers.JsonFail(c, http.StatusBadRequest, fmt.Sprintf("currency: is not a valid value. %v", model.Currency))
		return
	}

	helpers.PrintInfo(ctxt, c, "validating customer data..")
	if helpers.EmptyString(payment.Customer.DNI) {
		helpers.PrintError(ctxt, c, "customer dni could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "customer dni could not be empty")
		return
	}

	if helpers.EmptyString(payment.Customer.Name) {
		helpers.PrintError(ctxt, c, "customer name could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "customer name could not be empty")
		return
	}

	if helpers.EmptyString(payment.Customer.Email) {
		helpers.PrintError(ctxt, c, "customer email could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "customer email could not be empty")
		return
	}

	if helpers.EmptyString(payment.Customer.Phone) {
		helpers.PrintError(ctxt, c, "customer phone could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "customer phone could not be empty")
		return
	}

	helpers.PrintInfo(ctxt, c, "parsing dto data to customer model..")
	customerModel := model.PayCustomer{
		UUID:    helpers.NewUUIDString(),
		DNI:     payment.Customer.DNI,
		Name:    payment.Customer.Name,
		Email:   payment.Customer.Email,
		Phone:   payment.Customer.Phone,
		Address: payment.Customer.Address,
	}

	helpers.PrintInfo(ctxt, c, "parsing dto data to payment model..")
	exp := time.Now().Add(time.Minute * time.Duration(config.GetPaymentExpiration()))

	bankName := "simulator"

	paymentModel := model.PayTransaction{
		UUID:                     helpers.NewUUIDString(),
		PaymentCode:              helpers.NewUUIDString(),
		MerchantCode:             merchant.MerchantCode,
		Amount:                   payment.Amount,
		Description:              payment.Description,
		Currency:                 payment.Currency,
		Status:                   model.TransactionStatusPending,
		ExpirationProcess:        &exp,
		NaturalExpirationProcess: exp.Format("2006-01-02 15:04:05"),
		BankName:                 &bankName,
		CustomerUUID:             customerModel.UUID,
	}

	helpers.PrintInfo(ctxt, c, "saving customer data..")
	errInsertCustomer := cp.customerRepository.CreateCustomer(ctxt, &customerModel)

	if errInsertCustomer != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error saving customer information. Error: %s", errInsertCustomer.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	helpers.PrintInfo(ctxt, c, "saving payment data..")
	errInsertPayment := cp.paymentRepository.CreatePayment(ctxt, &paymentModel)

	if errInsertPayment != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error saving payment information. Error: %s", errInsertPayment.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	payment.PaymentCode = paymentModel.PaymentCode
	payment.MerchantCode = paymentModel.MerchantCode
	payment.Status = string(paymentModel.Status)
	payment.NaturalExpirationProcess = paymentModel.NaturalExpirationProcess
	payment.BankName = paymentModel.BankName

	helpers.JsonSuccess(c, payment)
	helpers.PrintInfo(ctxt, c, "execution of payment checkout finished.")
}

func (cp paymentServiceImpl) ProcessPayment(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute process payment method..")

	helpers.PrintInfo(ctxt, c, "finding payment by code..")
	paymentCode := c.Params.ByName("payment_code")

	payment, findPayment, err := cp.paymentRepository.FindPaymentByCode(ctxt, paymentCode)

	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error finding payment %s. Error: %s", paymentCode, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding payment: %s", paymentCode))
		return
	}

	if !findPayment {
		helpers.PrintError(ctxt, c, "payment not found", false)
		helpers.JsonFail(c, http.StatusNotFound, "payment not found")
		return
	}

	if payment.Status != model.TransactionStatusPending {
		helpers.PrintError(ctxt, c, fmt.Sprintf("could not process. payment is %s.", payment.Status), false)
		helpers.JsonFail(c, http.StatusConflict, fmt.Sprintf("could not process. payment is %s.", payment.Status))
		return
	}

	bankResp, err := cp.bankRepository.ProcessPayment(paymentCode)
	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("payment could not be processed. Error: %s.", err.Error()), false)
		helpers.PrintError(ctxt, c, err, false)
		helpers.JsonFail(c, http.StatusInternalServerError, "payment could not be processed")
		return
	}

	if bankResp.Code == 1000 {
		helpers.PrintInfo(ctxt, c, "payment processed")
		payment.Status = model.TransactionStatusSucceeded
	} else if bankResp.Code == 2000 {
		helpers.PrintInfo(ctxt, c, "insufficient funds")
		payment.Status = model.TransactionStatusFailure
	}

	if bankResp.Code != 1000 {
		payment.FailureReason = &bankResp.Message
	}

	if bankResp.Reference != "" {
		payment.BankReference = &bankResp.Reference
	}

	helpers.PrintInfo(ctxt, c, "saving payment data..")
	errInsertPayment := cp.paymentRepository.SavePayment(ctxt, payment)

	if errInsertPayment != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error saving payment information. Error: %s", errInsertPayment.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	if bankResp.Code != 1000 {
		helpers.PrintError(ctxt, c, fmt.Sprintf("payment no processed by the bank. info: %s", bankResp.Message), false)
		helpers.JsonFail(c, http.StatusForbidden, bankResp.Message)
	} else {
		helpers.JsonSuccess(c, bankResp)
	}

	helpers.PrintInfo(ctxt, c, "execution of payment processing finished.")
}

func (cp paymentServiceImpl) RefundPayment(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute refund method..")

	helpers.PrintInfo(ctxt, c, "finding payment by code..")
	paymentCode := c.Params.ByName("payment_code")

	payment, findPayment, err := cp.paymentRepository.FindPaymentByCode(ctxt, paymentCode)

	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error finding payment %s. Error: %s", paymentCode, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding payment: %s", paymentCode))
		return
	}

	if !findPayment {
		helpers.PrintError(ctxt, c, "payment not found", false)
		helpers.JsonFail(c, http.StatusNotFound, "payment not found")
		return
	}

	if payment.Status != model.TransactionStatusPending && payment.Status != model.TransactionStatusSucceeded {
		helpers.PrintError(ctxt, c, fmt.Sprintf("refund could not be processed process. payment is %s.", payment.Status), false)
		helpers.JsonFail(c, http.StatusConflict, fmt.Sprintf("refund could not be processed process. payment is %s.", payment.Status))
		return
	}

	helpers.PrintInfo(ctxt, c, "processing refund data..")
	bankResp, err := cp.bankRepository.ProcessRefund(paymentCode)
	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("refund could not be processed process. Error: %s.", err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "refund could not be processed")
		return
	}

	if bankResp.Code == 1000 {
		helpers.PrintInfo(ctxt, c, "refund processed")

		t := time.Now()
		refund := model.PayRefund{
			UUID:          helpers.NewUUIDString(),
			Code:          helpers.NewUUIDString(),
			BankReference: bankResp.Reference,
			Date:          &t,
		}

		helpers.PrintInfo(ctxt, c, "saving refund data..")
		errInsertRefund := cp.refundRepository.CreateRefund(ctxt, &refund)

		if errInsertRefund != nil {
			helpers.PrintError(ctxt, c, fmt.Sprintf("error saving refund information. Error: %s.", errInsertRefund.Error()), false)
			helpers.JsonFail(c, http.StatusInternalServerError, "error saving refund information")
			return
		}

		payment.RefundUUID = &refund.UUID
		payment.Status = model.TransactionStatusRefunded
	}

	if bankResp.Code != 1000 {
		payment.FailureReason = &bankResp.Message
	}

	helpers.PrintInfo(ctxt, c, "saving payment data..")
	errInsertPayment := cp.paymentRepository.SavePayment(ctxt, payment)

	if errInsertPayment != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error saving payment information. Error: %s.", errInsertPayment.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	if bankResp.Code != 1000 {
		helpers.PrintError(ctxt, c, fmt.Sprintf("refund no processed by the bank. info. Error: %s.", bankResp.Message), false)
		helpers.JsonFail(c, http.StatusBadRequest, bankResp.Message)
	} else {
		helpers.JsonSuccess(c, bankResp)
	}

	helpers.PrintInfo(ctxt, c, "execution of payment refund finished.")
}

func (cp paymentServiceImpl) CheckPayment(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute checking payment method..")

	helpers.PrintInfo(ctxt, c, "finding merchant by code..")
	paymentCode := c.Params.ByName("payment_code")

	payment, findPayment, err := cp.paymentRepository.FindPaymentByCode(ctxt, paymentCode)

	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error finding payment %s. Error: %s", paymentCode, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding payment: %s", paymentCode))
		return
	}

	if !findPayment {
		helpers.PrintError(ctxt, c, "payment not found", false)
		helpers.JsonFail(c, http.StatusNotFound, "payment not found")
		return
	}

	helpers.JsonSuccess(c, payment)
	helpers.PrintInfo(ctxt, c, "execution of payment checking finished.")
}
