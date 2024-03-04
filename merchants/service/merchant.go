package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ffalconesmera/payments-platform/merchants/helpers"
	"github.com/ffalconesmera/payments-platform/merchants/model"
	"github.com/ffalconesmera/payments-platform/merchants/model/dto"
	"github.com/ffalconesmera/payments-platform/merchants/repository"
	"github.com/gin-gonic/gin"
)

// MerchantService is an inteface sing up and login as merchant
type MerchantService interface {
	SingUp(ctx context.Context, c *gin.Context)
	Login(ctx context.Context, c *gin.Context)
	FindMerchantByCode(ctx context.Context, c *gin.Context)
}

type merchantServiceImpl struct {
	log                helpers.CustomLog
	validation         helpers.CustomValidation
	hash               helpers.CustomHash
	userRepository     repository.UserRepository
	merchantRepository repository.MerchantRepository
}

func NewMerchantService(log helpers.CustomLog, validation helpers.CustomValidation, hash helpers.CustomHash, userRepository repository.UserRepository, merchantRepository repository.MerchantRepository) *merchantServiceImpl {
	return &merchantServiceImpl{
		log:                log,
		validation:         validation,
		hash:               hash,
		userRepository:     userRepository,
		merchantRepository: merchantRepository,
	}
}

// SingUp: register a new merchant
func (m *merchantServiceImpl) SingUp(ctx context.Context, c *gin.Context) {
	m.log.PrintInfo(ctx, c, "start to execute sing up method..")

	response := helpers.NewJsonResponse()

	m.log.PrintInfo(ctx, c, "mapping merchant data..")
	var merchant dto.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		m.log.PrintError(ctx, c, fmt.Sprintf("error mapping merchant information. Error: %s", err), false)
		response.JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	m.log.PrintInfo(ctx, c, "validating merchant data..")
	if m.validation.EmptyString(merchant.Name) {
		m.log.PrintError(ctx, c, "merchant name could not be empty", false)
		response.JsonFail(c, http.StatusBadRequest, "merchant name could not be empty")
		return
	}

	if m.validation.EmptyString(merchant.User.Username) {
		m.log.PrintError(ctx, c, "username could not be empty", false)
		response.JsonFail(c, http.StatusBadRequest, "username could not be empty")
		return
	}

	if m.validation.EmptyString(merchant.User.Email) {
		m.log.PrintError(ctx, c, "email could not be empty", false)
		response.JsonFail(c, http.StatusBadRequest, "email could not be empty")
		return
	}

	if m.validation.PasswordInvalid(merchant.User.Password) {
		m.log.PrintError(ctx, c, m.validation.PasswordInvalidMessage(), false)
		response.JsonFail(c, http.StatusBadRequest, m.validation.PasswordInvalidMessage())
		return
	}

	_, findUser, err := m.userRepository.FindUserByUsername(merchant.User.Username)

	if err != nil {
		m.log.PrintError(ctx, c, fmt.Sprintf("error finding user %s. Error: %s", merchant.User.Username, err.Error()), false)
		response.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding user %s.", merchant.User.Username))
		return
	}

	if findUser {
		m.log.PrintError(ctx, c, fmt.Sprintf("username %s is alreay exits.", merchant.User.Username), false)
		response.JsonFail(c, http.StatusConflict, fmt.Sprintf("username %s is alreay exits", merchant.User.Username))
		return
	}

	m.log.PrintInfo(ctx, c, "generating password hash..")
	hash := m.hash.GenerateHashPassword(merchant.User.Password)
	password := hash

	m.log.PrintInfo(ctx, c, "parsing dto data to merchant model..")
	merchant.MerchantCode = m.hash.NewUUIDString()
	merchantModel := model.PayMerchant{
		UUID:         m.hash.NewUUIDString(),
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	m.log.PrintInfo(ctx, c, "parsing dto data to user model..")
	userModel := model.PayUser{
		UUID:         m.hash.NewUUIDString(),
		Username:     merchant.User.Username,
		Email:        merchant.User.Email,
		Password:     password,
		MerchantUUID: merchantModel.UUID,
	}

	tx := m.merchantRepository.BeginTransaction()

	m.log.PrintInfo(ctx, c, "saving merchant data..")
	errInsertMerchant := m.merchantRepository.CreateMerchant(&merchantModel)

	if errInsertMerchant != nil {
		m.merchantRepository.RollbackTransaction(tx)
		m.log.PrintError(ctx, c, fmt.Sprintf("error saving merchant information. Error: %s", errInsertMerchant.Error()), false)
		response.JsonFail(c, http.StatusInternalServerError, "error saving merchant information")
		return
	}

	m.log.PrintInfo(ctx, c, "saving user data..")
	errUserMerchant := m.userRepository.CreateUser(&userModel)

	if errUserMerchant != nil {
		m.merchantRepository.RollbackTransaction(tx)
		m.log.PrintError(ctx, c, fmt.Sprintf("error saving merchant information. Error: %s", errUserMerchant.Error()), false)
		response.JsonFail(c, http.StatusInternalServerError, "error saving merchant information")
		return
	}

	m.merchantRepository.CommitTransaction(tx)

	merchant.User.Password = ""

	response.JsonSuccess(c, merchant)
	m.log.PrintInfo(ctx, c, "execution of sing up finished.")
}

// Login: generate a authorization token for access as merchant
func (m *merchantServiceImpl) Login(ctx context.Context, c *gin.Context) {
	m.log.PrintInfo(ctx, c, "start to execute login method..")

	response := helpers.NewJsonResponse()

	m.log.PrintInfo(ctx, c, "mapping login data..")
	var loginDTO dto.LoginInput
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		m.log.PrintError(ctx, c, fmt.Sprintf("error mapping merchant information. Error: %s", err), false)
		response.JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	m.log.PrintInfo(ctx, c, "validating login data..")
	if m.validation.EmptyString(loginDTO.Username) {
		m.log.PrintError(ctx, c, "username could not be empty", false)
		response.JsonFail(c, http.StatusBadRequest, "username could not be empty")
		return
	}

	if m.validation.EmptyString(loginDTO.Password) {
		m.log.PrintError(ctx, c, "password could not be empty", false)
		response.JsonFail(c, http.StatusBadRequest, "password could not be empty")
		return
	}

	m.log.PrintInfo(ctx, c, "finding user by username..")
	user, findUser, err := m.userRepository.FindUserByUsername(loginDTO.Username)

	if err != nil {
		m.log.PrintError(ctx, c, fmt.Sprintf("error finding user %s. Error: %s", loginDTO.Username, err.Error()), false)
		response.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding user: %s", loginDTO.Username))
		return
	}

	if !findUser {
		m.log.PrintError(ctx, c, fmt.Sprintf("user %s not found.", loginDTO.Username), false)
		response.JsonFail(c, http.StatusNotFound, fmt.Sprintf("user %s not found", loginDTO.Username))
		return
	}

	if !m.hash.CheckHashPassword(user.Password, loginDTO.Password) {
		m.log.PrintError(ctx, c, fmt.Sprintf("password incorrect:  %s", loginDTO.Username), false)
		response.JsonFail(c, http.StatusUnauthorized, "password incorrect")
		return
	}

	m.log.PrintInfo(ctx, c, "creating jwt")

	token, err := m.hash.CreateJWToken(user.Username)
	if err != nil {
		m.log.PrintError(ctx, c, fmt.Sprintf("error getting token:  %s. Error: %s", user.Username, err.Error()), false)
		response.JsonFail(c, http.StatusInternalServerError, "error getting token")
		return
	}

	loginOutput := dto.LoginOutput{
		Username: user.Username,
		Message:  "you are logged..!",
		Token:    token,
	}

	response.JsonSuccess(c, loginOutput)
	m.log.PrintInfo(ctx, c, fmt.Sprintf("user logged successfully: %s", loginDTO.Username))
}

// FindMerchantByCode: retrieve merchant data from repository
func (m *merchantServiceImpl) FindMerchantByCode(ctx context.Context, c *gin.Context) {
	m.log.PrintInfo(ctx, c, "start to execute find merchant by code method..")

	response := helpers.NewJsonResponse()

	m.log.PrintInfo(ctx, c, "finding merchant by code..")
	merchantCode := c.Params.ByName("merchant_code")
	merchant, findMerchant, err := m.merchantRepository.FindMerchantByCode(merchantCode)

	if err != nil {
		m.log.PrintError(ctx, c, fmt.Sprintf("error finding merchant %s. Error: %s", merchantCode, err.Error()), false)
		response.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding merchant %s", merchantCode))
		return
	}

	if !findMerchant {
		m.log.PrintError(ctx, c, fmt.Sprintf("merchant not found %s.", merchantCode), false)
		response.JsonFail(c, http.StatusNotFound, "merchant not found")
		return
	}

	merchantDTO := dto.Merchant{
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	response.JsonSuccess(c, merchantDTO)
	m.log.PrintInfo(ctx, c, "execution of find merchant by code finished.")
}
