package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ffalconesmera/payments-platform/merchants/helpers"
	"github.com/ffalconesmera/payments-platform/merchants/model"
	"github.com/ffalconesmera/payments-platform/merchants/model/dto"
	"github.com/ffalconesmera/payments-platform/merchants/repository"
)

// MerchantService is an inteface sing up and login as merchant
type MerchantService interface {
	SingUp(ctxt context.Context, merchant dto.Merchant) (*dto.Merchant, error)
	Login(ctxt context.Context, login dto.LoginInput) (*dto.LoginOutput, error)
	FindMerchantByCode(ctxt context.Context, merchantCode string) (*dto.Merchant, error)
}

type merchantServiceImpl struct {
	userRepository     repository.UserRepository
	merchantRepository repository.MerchantRepository
}

func NewMerchantService(userRepository repository.UserRepository, merchantRepository repository.MerchantRepository) MerchantService {
	return &merchantServiceImpl{
		userRepository:     userRepository,
		merchantRepository: merchantRepository,
	}
}

// SingUp: register a new merchant
func (m *merchantServiceImpl) SingUp(ctxt context.Context, merchant dto.Merchant) (*dto.Merchant, error) {
	_, findUser, err := m.userRepository.FindUserByUsername(ctxt, merchant.User.Username)

	if err != nil {
		return nil, err
	}

	if findUser {
		return nil, fmt.Errorf("username %s is alreay exits", merchant.User.Username)
	}

	hash := helpers.GenerateHashPassword(merchant.User.Password)
	password := hash

	merchant.MerchantCode = helpers.NewUUIDString()
	merchantModel := model.PayMerchant{
		UUID:         helpers.NewUUIDString(),
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	userModel := model.PayUser{
		UUID:         helpers.NewUUIDString(),
		Username:     merchant.User.Username,
		Email:        merchant.User.Email,
		Password:     password,
		MerchantUUID: merchantModel.UUID,
	}

	errInsertMerchant := m.merchantRepository.CreateMerchant(ctxt, &merchantModel)

	if errInsertMerchant != nil {
		return nil, errInsertMerchant
	}

	errUserMerchant := m.userRepository.CreateUser(ctxt, &userModel)

	if errUserMerchant != nil {
		return nil, errUserMerchant
	}

	merchant.User.Password = ""

	return &merchant, nil
}

// Login: generate a authorization token for access as merchant
func (m *merchantServiceImpl) Login(ctxt context.Context, login dto.LoginInput) (*dto.LoginOutput, error) {
	user, findUser, err := m.userRepository.FindUserByUsername(ctxt, login.Username)

	if err != nil {
		return nil, err
	}

	if !findUser {
		return nil, fmt.Errorf("user %s not found", login.Username)
	}

	if !helpers.CheckHashPassword(user.Password, login.Password) {
		return nil, errors.New("password incorrect")
	}

	token, err := helpers.CreateJWToken(user.Username)
	if err != nil {
		return nil, errors.New("error generating token")
	}

	loginOutput := dto.LoginOutput{
		Username: user.Username,
		Message:  "you are logged..!",
		Token:    token,
	}

	return &loginOutput, nil
}

// FindMerchantByCode: retrieve merchant data from repository
func (m *merchantServiceImpl) FindMerchantByCode(ctxt context.Context, merchantCode string) (*dto.Merchant, error) {

	merchant, findMerchant, err := m.merchantRepository.FindMerchantByCode(ctxt, merchantCode)

	if err != nil {
		return nil, err
	}

	if !findMerchant {
		return nil, errors.New("merchant not found")
	}

	merchantDTO := dto.Merchant{
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	return &merchantDTO, nil
}
