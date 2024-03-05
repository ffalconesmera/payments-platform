package repository

import (
	"errors"
	"fmt"

	"github.com/ffalconesmera/payments-platform/payments/config"
	"github.com/ffalconesmera/payments-platform/payments/externals/dto"
)

// BankRepository is an interface for receive information from merchants microservice
type BankRepository interface {
	ProcessPayment(paymentCode string) (*dto.BankPayment, error)
	ProcessRefund(paymentCode string) (*dto.BankRefund, error)
}

type bankRepositoryImpl struct {
}

func NewBankRepository() BankRepository {
	return &bankRepositoryImpl{}
}

// ProcessPayment: send an order of charge to the bank
func (m *bankRepositoryImpl) ProcessPayment(paymentCode string) (*dto.BankPayment, error) {
	var bankPayment dto.BankPayment
	err := SendRequestApiExternal(fmt.Sprintf("%s/%s", config.Config().GetMerchantEndpoint(), paymentCode), "GET", "", &bankPayment)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &bankPayment, nil
}

func (m *bankRepositoryImpl) ProcessRefund(paymentCode string) (*dto.BankRefund, error) {
	var bankRefund dto.BankRefund
	err := SendRequestApiExternal(fmt.Sprintf("%s/%s", config.Config().GetMerchantEndpoint(), paymentCode), "GET", "", &bankRefund)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &bankRefund, nil
}
