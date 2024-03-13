package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ffalconesmera/payments-platform/payments/config"
	ext_dto "github.com/ffalconesmera/payments-platform/payments/externals/dto"
	ext_repository "github.com/ffalconesmera/payments-platform/payments/externals/repository"
	"github.com/ffalconesmera/payments-platform/payments/helpers"
	"github.com/ffalconesmera/payments-platform/payments/model"
	"github.com/ffalconesmera/payments-platform/payments/model/dto"
	"github.com/ffalconesmera/payments-platform/payments/repository"
)

// PaymentService is an inteface for process payments

// ProcessPayment: receive card information and get send to bank provider for process charge
// RefundPayment: send to bank an order of refund
// CheckPayment: returns the information of a previous payment
type PaymentService interface {
	CheckoutPayment(ctxt context.Context, merchantCode string, payment dto.Payment) (*dto.Payment, error)
	ProcessPayment(ctxt context.Context, paymentCode string, cardInfo string) (*ext_dto.BankPayment, error)
	RefundPayment(ctxt context.Context, paymentCode string, merchantCode string, refundInfo string) (*ext_dto.BankRefund, error)
	CheckPayment(ctxt context.Context, paymentCode string, merchantCode string) (*dto.Payment, error)
}

type paymentServiceImpl struct {
	paymentRepository  repository.PaymentRepository
	refundRepository   repository.RefundRepository
	customerRepository repository.CustomerRepository
	bankRepository     ext_repository.BankRepository
	merchantRepository ext_repository.MerchantRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository, refundRepository repository.RefundRepository, customerRepository repository.CustomerRepository, bankRepository ext_repository.BankRepository, merchantRepository ext_repository.MerchantRepository) PaymentService {
	return &paymentServiceImpl{
		paymentRepository:  paymentRepository,
		refundRepository:   refundRepository,
		customerRepository: customerRepository,
		bankRepository:     bankRepository,
		merchantRepository: merchantRepository,
	}
}

// CheckoutPayment: generate an payment order and return a payment code for future process
func (cp *paymentServiceImpl) CheckoutPayment(ctxt context.Context, merchantCode string, payment dto.Payment) (*dto.Payment, error) {
	var merchant *ext_dto.Merchant
	merchant, err := cp.merchantRepository.FindMerchantByCode(merchantCode)

	if err != nil {
		return nil, err
	}

	customerModel := model.PayCustomer{
		UUID:    helpers.NewUUIDString(),
		DNI:     payment.Customer.DNI,
		Name:    payment.Customer.Name,
		Email:   payment.Customer.Email,
		Phone:   payment.Customer.Phone,
		Address: payment.Customer.Address,
	}

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

	if err := cp.customerRepository.CreateCustomer(ctxt, &customerModel); err != nil {
		return nil, err
	}

	if err := cp.paymentRepository.CreatePayment(ctxt, &paymentModel); err != nil {
		return nil, err
	}

	payment.PaymentCode = paymentModel.PaymentCode
	payment.MerchantCode = paymentModel.MerchantCode
	payment.Status = string(paymentModel.Status)
	payment.NaturalExpirationProcess = paymentModel.NaturalExpirationProcess
	payment.BankName = paymentModel.BankName

	return &payment, nil
}

func (cp paymentServiceImpl) ProcessPayment(ctxt context.Context, paymentCode string, cardInfo string) (*ext_dto.BankPayment, error) {
	payment, findPayment, err := cp.paymentRepository.FindPaymentByCode(ctxt, paymentCode)

	if err != nil {
		return nil, err
	}

	if !findPayment {
		return nil, errors.New("payment not found")
	}

	if payment.Status != model.TransactionStatusPending {
		return nil, fmt.Errorf("could not process. payment is %s", payment.Status)
	}

	bankResp, err := cp.bankRepository.ProcessPayment(paymentCode, cardInfo)
	if err != nil {
		return nil, err
	}

	if bankResp.Code == 1000 {
		payment.Status = model.TransactionStatusSucceeded
	} else if bankResp.Code == 2000 {
		payment.Status = model.TransactionStatusFailure
	}

	if bankResp.Code != 1000 {
		payment.FailureReason = &bankResp.Message
	}

	if bankResp.Reference != "" {
		payment.BankReference = &bankResp.Reference
	}

	if err := cp.paymentRepository.SavePayment(ctxt, payment); err != nil {
		return nil, err
	}

	if bankResp.Code != 1000 {
		return nil, errors.New(bankResp.Message)
	}

	return bankResp, nil
}

func (cp paymentServiceImpl) RefundPayment(ctxt context.Context, paymentCode string, merchantCode string, refundInfo string) (*ext_dto.BankRefund, error) {
	payment, findPayment, err := cp.paymentRepository.FindPaymentByCode(ctxt, paymentCode)

	if err != nil {
		return nil, err
	}

	if !findPayment {
		return nil, errors.New("payment not found")
	}

	if payment.MerchantCode != merchantCode {
		return nil, errors.New("payment not found")
	}

	if payment.Status != model.TransactionStatusPending && payment.Status != model.TransactionStatusSucceeded {
		return nil, errors.New("payment not found")
	}

	bankResp, err := cp.bankRepository.ProcessRefund(paymentCode, refundInfo)
	if err != nil {
		return nil, err
	}

	if bankResp.Code == 1000 {
		t := time.Now()
		refund := model.PayRefund{
			UUID:          helpers.NewUUIDString(),
			Code:          helpers.NewUUIDString(),
			BankReference: bankResp.Reference,
			Date:          &t,
		}

		errInsertRefund := cp.refundRepository.CreateRefund(ctxt, &refund)

		if errInsertRefund != nil {
			return nil, errInsertRefund
		}

		payment.RefundUUID = &refund.UUID
		payment.Status = model.TransactionStatusRefunded
	}

	if bankResp.Code != 1000 {
		payment.FailureReason = &bankResp.Message
	}

	errInsertPayment := cp.paymentRepository.SavePayment(ctxt, payment)

	if errInsertPayment != nil {
		return nil, errInsertPayment
	}

	if bankResp.Code != 1000 {
		return nil, errors.New(bankResp.Message)
	}

	return bankResp, nil
}

func (cp paymentServiceImpl) CheckPayment(ctxt context.Context, paymentCode string, merchantCode string) (*dto.Payment, error) {
	pay, findPayment, err := cp.paymentRepository.FindPaymentByCode(ctxt, paymentCode)

	if err != nil {
		return nil, err
	}

	if !findPayment {
		return nil, errors.New("payment not found")
	}

	if pay.MerchantCode != merchantCode {
		return nil, errors.New("payment not found")
	}

	customer, findCustomer, err := cp.customerRepository.FindCustomerById(ctxt, pay.CustomerUUID)

	if err != nil {
		return nil, err
	}

	if !findCustomer {
		return nil, errors.New("customer not found")
	}

	payment := dto.Payment{
		PaymentCode:              pay.PaymentCode,
		MerchantCode:             pay.MerchantCode,
		Amount:                   pay.Amount,
		Description:              pay.Description,
		Currency:                 pay.Currency,
		Status:                   string(pay.Status),
		NaturalExpirationProcess: pay.NaturalExpirationProcess,
		FailureReason:            pay.FailureReason,
		BankReference:            pay.BankReference,
		BankName:                 pay.BankName,
		Customer: dto.Customer{
			DNI:     customer.DNI,
			Name:    customer.Name,
			Email:   customer.Email,
			Phone:   customer.Phone,
			Address: customer.Address,
		},
	}

	return &payment, nil
}
