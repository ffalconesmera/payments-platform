package repository

import (
	"errors"
	"fmt"

	"github.com/ffalconesmera/payments-platform/payments/config"
	"github.com/ffalconesmera/payments-platform/payments/externals/dto"
)

// BankRepository is an interface for receive information from merchants microservice
type BankRepository interface {
	ProcessPayment(paymentCode, body string) (*dto.BankPayment, error)
	ProcessRefund(paymentCode, body string) (*dto.BankRefund, error)
}

type bankRepositoryImpl struct {
}

func NewBankRepository() BankRepository {
	return &bankRepositoryImpl{}
}

// ProcessPayment: send an order of charge to the bank
func (c *bankRepositoryImpl) ProcessPayment(paymentCode, body string) (*dto.BankPayment, error) {
	var bankPayment dto.BankPayment
	err := SendRequestApiExternal(fmt.Sprintf("%s/payments", config.GetBankEndpoint()), "POST", body, &bankPayment)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &bankPayment, nil
}

func (c *bankRepositoryImpl) ProcessRefund(paymentCode, body string) (*dto.BankRefund, error) {
	var bankRefund dto.BankRefund
	err := SendRequestApiExternal(fmt.Sprintf("%s/refunds", config.GetBankEndpoint()), "POST", body, &bankRefund)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &bankRefund, nil
}
