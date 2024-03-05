package repository

import (
	"errors"
	"fmt"

	"github.com/ffalconesmera/payments-platform/payments/config"
	"github.com/ffalconesmera/payments-platform/payments/externals/dto"
)

// MerchantRepository is an interface for receive information from merchants microservice
type MerchantRepository interface {
	FindMerchantByCode(paymentCode string) (*dto.Merchant, error)
}

type merchantRepositoryImpl struct {
}

func NewMerchantRepository() MerchantRepository {
	return &merchantRepositoryImpl{}
}

// FindMerchantByCode: request merchant information by code
func (m *merchantRepositoryImpl) FindMerchantByCode(paymentCode string) (*dto.Merchant, error) {
	var jsonMerchant dto.JSONMerchant
	err := SendRequestApiExternal(fmt.Sprintf("%s/%s", config.GetMerchantEndpoint(), paymentCode), "GET", "", &jsonMerchant)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if jsonMerchant.Merchant.MerchantCode == "" {
		return nil, errors.New("merchant could not be found")
	}

	return &jsonMerchant.Merchant, nil
}
