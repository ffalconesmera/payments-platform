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
	CheckoutPayment(ctx context.Context, c *gin.Context)
	ProcessPayment(ctx context.Context, c *gin.Context)
	RefundPayment(ctx context.Context, c *gin.Context)
	CheckPayment(ctx context.Context, c *gin.Context)
}

type paymentServiceImpl struct {
	paymentRepository  repository.PaymentRepository
	refundRepository   repository.RefundRepository
	customerRepository repository.CustomerRepository
	bankRepository     ext_repository.BankRepository
	merchantRepository ext_repository.MerchantRepository
}

func NewPaymentService(ctx context.Context, paymentRepository repository.PaymentRepository, refundRepository repository.RefundRepository, customerRepository repository.CustomerRepository, bankRepository ext_repository.BankRepository, merchantRepository ext_repository.MerchantRepository) PaymentService {
	return &paymentServiceImpl{
		paymentRepository:  paymentRepository,
		refundRepository:   refundRepository,
		customerRepository: customerRepository,
		bankRepository:     bankRepository,
		merchantRepository: merchantRepository,
	}
}

// CheckoutPayment: generate an payment order and return a payment code for future process
func (p *paymentServiceImpl) CheckoutPayment(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute sing up method..")

	helpers.CustomLog().PrintInfo(ctx, c, "finding merchant by code..")
	merchantCode := c.Params.ByName("merchant_code")

	merchant, err := p.merchantRepository.FindMerchantByCode(merchantCode)

	if err != nil {
		helpers.CustomLog().PrintInfo(ctx, c, fmt.Sprintf("merchant not found %s. Error: %s", merchantCode, err.Error()))
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("merchant not found %s", merchantCode))
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "mapping payment data..")
	var payment *dto.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error mapping payment information. Error: %s", err), false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "validating merchant data..")

	if helpers.CustomValidation().InvalidFloat(payment.Amount) {
		helpers.CustomLog().PrintError(ctx, c, "amount could not be zero", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "amount could not be zero")
		return
	}

	if helpers.CustomValidation().EmptyString(payment.Description) {
		helpers.CustomLog().PrintError(ctx, c, "description could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "description could not be empty")
		return
	}

	_, okCurrency := model.Currency[payment.Currency]

	if !okCurrency {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("currency: is not a valid value. %v", model.Currency), false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, fmt.Sprintf("currency: is not a valid value. %v", model.Currency))
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "validating customer data..")
	if helpers.CustomValidation().EmptyString(payment.Customer.DNI) {
		helpers.CustomLog().PrintError(ctx, c, "customer dni could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "customer dni could not be empty")
		return
	}

	if helpers.CustomValidation().EmptyString(payment.Customer.Name) {
		helpers.CustomLog().PrintError(ctx, c, "customer name could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "customer name could not be empty")
		return
	}

	if helpers.CustomValidation().EmptyString(payment.Customer.Email) {
		helpers.CustomLog().PrintError(ctx, c, "customer email could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "customer email could not be empty")
		return
	}

	if helpers.CustomValidation().EmptyString(payment.Customer.Phone) {
		helpers.CustomLog().PrintError(ctx, c, "customer phone could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "customer phone could not be empty")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "parsing dto data to customer model..")
	customerModel := model.PayCustomer{
		UUID:    helpers.CustomHash().NewUUIDString(),
		DNI:     payment.Customer.DNI,
		Name:    payment.Customer.Name,
		Email:   payment.Customer.Email,
		Phone:   payment.Customer.Phone,
		Address: payment.Customer.Address,
	}

	helpers.CustomLog().PrintInfo(ctx, c, "parsing dto data to payment model..")
	exp := time.Now().Add(time.Minute * time.Duration(config.Config().GetPaymentExpiration()))

	bankName := "simulator"

	paymentModel := model.PayTransaction{
		UUID:                     helpers.CustomHash().NewUUIDString(),
		PaymentCode:              helpers.CustomHash().NewUUIDString(),
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

	tx := p.paymentRepository.BeginTransaction()

	helpers.CustomLog().PrintInfo(ctx, c, "saving customer data..")
	errInsertCustomer := p.customerRepository.CreateCustomer(&customerModel)

	if errInsertCustomer != nil {
		p.paymentRepository.RollbackTransaction(tx)
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving customer information. Error: %s", errInsertCustomer.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "saving payment data..")
	errInsertPayment := p.paymentRepository.CreatePayment(&paymentModel)

	if errInsertPayment != nil {
		p.paymentRepository.RollbackTransaction(tx)
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving payment information. Error: %s", errInsertPayment.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	p.paymentRepository.CommitTransaction(tx)

	payment.PaymentCode = paymentModel.PaymentCode
	payment.MerchantCode = paymentModel.MerchantCode
	payment.Status = string(paymentModel.Status)
	payment.NaturalExpirationProcess = paymentModel.NaturalExpirationProcess
	payment.BankName = paymentModel.BankName

	helpers.JsonResponse().JsonSuccess(c, payment)
	helpers.CustomLog().PrintInfo(ctx, c, "execution of payment checkout finished.")
}

func (p paymentServiceImpl) ProcessPayment(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute process payment method..")

	helpers.CustomLog().PrintInfo(ctx, c, "finding payment by code..")
	paymentCode := c.Params.ByName("payment_code")

	payment, findPayment, err := p.paymentRepository.FindPaymentByCode(paymentCode)

	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error finding payment %s. Error: %s", paymentCode, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding payment: %s", paymentCode))
		return
	}

	if !findPayment {
		helpers.CustomLog().PrintError(ctx, c, "payment not found", false)
		helpers.JsonResponse().JsonFail(c, http.StatusNotFound, "payment not found")
		return
	}

	if payment.Status != model.TransactionStatusPending {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("could not process. payment is %s.", payment.Status), false)
		helpers.JsonResponse().JsonFail(c, http.StatusConflict, fmt.Sprintf("could not process. payment is %s.", payment.Status))
		return
	}

	bankResp, err := p.bankRepository.ProcessPayment(paymentCode)
	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("payment could not be processed. Error: %s.", err.Error()), false)
		helpers.CustomLog().PrintError(ctx, c, err, false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "payment could not be processed")
		return
	}

	if bankResp.Code == 1000 {
		helpers.CustomLog().PrintInfo(ctx, c, "payment processed")
		payment.Status = model.TransactionStatusSucceeded
	} else if bankResp.Code == 2000 {
		helpers.CustomLog().PrintInfo(ctx, c, "insufficient funds")
		payment.Status = model.TransactionStatusFailure
	}

	if bankResp.Code != 1000 {
		payment.FailureReason = &bankResp.Message
	}

	if bankResp.Reference != "" {
		payment.BankReference = &bankResp.Reference
	}

	helpers.CustomLog().PrintInfo(ctx, c, "saving payment data..")
	errInsertPayment := p.paymentRepository.SavePayment(payment)

	if errInsertPayment != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving payment information. Error: %s", errInsertPayment.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	if bankResp.Code != 1000 {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("payment no processed by the bank. info: %s", bankResp.Message), false)
		helpers.JsonResponse().JsonFail(c, http.StatusForbidden, bankResp.Message)
	} else {
		helpers.JsonResponse().JsonSuccess(c, bankResp)
	}

	helpers.CustomLog().PrintInfo(ctx, c, "execution of payment processing finished.")
}

func (p paymentServiceImpl) RefundPayment(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute refund method..")

	helpers.CustomLog().PrintInfo(ctx, c, "finding payment by code..")
	paymentCode := c.Params.ByName("payment_code")

	payment, findPayment, err := p.paymentRepository.FindPaymentByCode(paymentCode)

	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error finding payment %s. Error: %s", paymentCode, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding payment: %s", paymentCode))
		return
	}

	if !findPayment {
		helpers.CustomLog().PrintError(ctx, c, "payment not found", false)
		helpers.JsonResponse().JsonFail(c, http.StatusNotFound, "payment not found")
		return
	}

	if payment.Status != model.TransactionStatusPending && payment.Status != model.TransactionStatusSucceeded {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("refund could not be processed process. payment is %s.", payment.Status), false)
		helpers.JsonResponse().JsonFail(c, http.StatusConflict, fmt.Sprintf("refund could not be processed process. payment is %s.", payment.Status))
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "processing refund data..")
	bankResp, err := p.bankRepository.ProcessRefund(paymentCode)
	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("refund could not be processed process. Error: %s.", err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "refund could not be processed")
		return
	}

	tx := p.paymentRepository.BeginTransaction()

	if bankResp.Code == 1000 {
		helpers.CustomLog().PrintInfo(ctx, c, "refund processed")

		t := time.Now()
		refund := model.PayRefund{
			UUID:          helpers.CustomHash().NewUUIDString(),
			Code:          helpers.CustomHash().NewUUIDString(),
			BankReference: bankResp.Reference,
			Date:          &t,
		}

		helpers.CustomLog().PrintInfo(ctx, c, "saving refund data..")
		errInsertRefund := p.refundRepository.CreateRefund(&refund)

		if errInsertRefund != nil {
			p.paymentRepository.RollbackTransaction(tx)
			helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving refund information. Error: %s.", errInsertRefund.Error()), false)
			helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving refund information")
			return
		}

		payment.RefundUUID = &refund.UUID
		payment.Status = model.TransactionStatusRefunded
	}

	if bankResp.Code != 1000 {
		payment.FailureReason = &bankResp.Message
	}

	helpers.CustomLog().PrintInfo(ctx, c, "saving payment data..")
	errInsertPayment := p.paymentRepository.SavePayment(payment)

	if errInsertPayment != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving payment information. Error: %s.", errInsertPayment.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving payment information")
		return
	}

	p.paymentRepository.CommitTransaction(tx)

	if bankResp.Code != 1000 {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("refund no processed by the bank. info. Error: %s.", bankResp.Message), false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, bankResp.Message)
	} else {
		helpers.JsonResponse().JsonSuccess(c, bankResp)
	}

	helpers.CustomLog().PrintInfo(ctx, c, "execution of payment refund finished.")
}

func (p paymentServiceImpl) CheckPayment(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute checking payment method..")

	helpers.CustomLog().PrintInfo(ctx, c, "finding merchant by code..")
	paymentCode := c.Params.ByName("payment_code")

	payment, findPayment, err := p.paymentRepository.FindPaymentByCode(paymentCode)

	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error finding payment %s. Error: %s", paymentCode, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding payment: %s", paymentCode))
		return
	}

	if !findPayment {
		helpers.CustomLog().PrintError(ctx, c, "payment not found", false)
		helpers.JsonResponse().JsonFail(c, http.StatusNotFound, "payment not found")
		return
	}

	helpers.JsonResponse().JsonSuccess(c, payment)
	helpers.CustomLog().PrintInfo(ctx, c, "execution of payment checking finished.")
}
