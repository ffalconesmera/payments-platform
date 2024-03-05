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
	userRepository     repository.UserRepository
	merchantRepository repository.MerchantRepository
}

func NewMerchantService(userRepository repository.UserRepository, merchantRepository repository.MerchantRepository) *merchantServiceImpl {
	return &merchantServiceImpl{
		userRepository:     userRepository,
		merchantRepository: merchantRepository,
	}
}

// SingUp: register a new merchant
func (m *merchantServiceImpl) SingUp(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute sing up method..")

	helpers.CustomLog().PrintInfo(ctx, c, "mapping merchant data..")
	var merchant dto.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error mapping merchant information. Error: %s", err), false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "validating merchant data..")
	if helpers.CustomValidation().EmptyString(merchant.Name) {
		helpers.CustomLog().PrintError(ctx, c, "merchant name could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "merchant name could not be empty")
		return
	}

	if helpers.CustomValidation().EmptyString(merchant.User.Username) {
		helpers.CustomLog().PrintError(ctx, c, "username could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "username could not be empty")
		return
	}

	if helpers.CustomValidation().EmptyString(merchant.User.Email) {
		helpers.CustomLog().PrintError(ctx, c, "email could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "email could not be empty")
		return
	}

	if helpers.CustomValidation().PasswordInvalid(merchant.User.Password) {
		helpers.CustomLog().PrintError(ctx, c, helpers.CustomValidation().PasswordInvalidMessage(), false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, helpers.CustomValidation().PasswordInvalidMessage())
		return
	}

	_, findUser, err := m.userRepository.FindUserByUsername(merchant.User.Username)

	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error finding user %s. Error: %s", merchant.User.Username, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding user %s.", merchant.User.Username))
		return
	}

	if findUser {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("username %s is alreay exits.", merchant.User.Username), false)
		helpers.JsonResponse().JsonFail(c, http.StatusConflict, fmt.Sprintf("username %s is alreay exits", merchant.User.Username))
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "generating password hash..")
	hash := helpers.CustomHash().GenerateHashPassword(merchant.User.Password)
	password := hash

	helpers.CustomLog().PrintInfo(ctx, c, "parsing dto data to merchant model..")
	merchant.MerchantCode = helpers.CustomHash().NewUUIDString()
	merchantModel := model.PayMerchant{
		UUID:         helpers.CustomHash().NewUUIDString(),
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	helpers.CustomLog().PrintInfo(ctx, c, "parsing dto data to user model..")
	userModel := model.PayUser{
		UUID:         helpers.CustomHash().NewUUIDString(),
		Username:     merchant.User.Username,
		Email:        merchant.User.Email,
		Password:     password,
		MerchantUUID: merchantModel.UUID,
	}

	tx := m.merchantRepository.BeginTransaction()

	helpers.CustomLog().PrintInfo(ctx, c, "saving merchant data..")
	errInsertMerchant := m.merchantRepository.CreateMerchant(&merchantModel)

	if errInsertMerchant != nil {
		m.merchantRepository.RollbackTransaction(tx)
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving merchant information. Error: %s", errInsertMerchant.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving merchant information")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "saving user data..")
	errUserMerchant := m.userRepository.CreateUser(&userModel)

	if errUserMerchant != nil {
		m.merchantRepository.RollbackTransaction(tx)
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error saving merchant information. Error: %s", errUserMerchant.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error saving merchant information")
		return
	}

	m.merchantRepository.CommitTransaction(tx)

	merchant.User.Password = ""

	helpers.JsonResponse().JsonSuccess(c, merchant)
	helpers.CustomLog().PrintInfo(ctx, c, "execution of sing up finished.")
}

// Login: generate a authorization token for access as merchant
func (m *merchantServiceImpl) Login(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute login method..")

	helpers.CustomLog().PrintInfo(ctx, c, "mapping login data..")
	var loginDTO dto.LoginInput
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error mapping merchant information. Error: %s", err), false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "validating login data..")
	if helpers.CustomValidation().EmptyString(loginDTO.Username) {
		helpers.CustomLog().PrintError(ctx, c, "username could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "username could not be empty")
		return
	}

	if helpers.CustomValidation().EmptyString(loginDTO.Password) {
		helpers.CustomLog().PrintError(ctx, c, "password could not be empty", false)
		helpers.JsonResponse().JsonFail(c, http.StatusBadRequest, "password could not be empty")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "finding user by username..")
	user, findUser, err := m.userRepository.FindUserByUsername(loginDTO.Username)

	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error finding user %s. Error: %s", loginDTO.Username, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding user: %s", loginDTO.Username))
		return
	}

	if !findUser {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("user %s not found.", loginDTO.Username), false)
		helpers.JsonResponse().JsonFail(c, http.StatusNotFound, fmt.Sprintf("user %s not found", loginDTO.Username))
		return
	}

	if !helpers.CustomHash().CheckHashPassword(user.Password, loginDTO.Password) {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("password incorrect:  %s", loginDTO.Username), false)
		helpers.JsonResponse().JsonFail(c, http.StatusUnauthorized, "password incorrect")
		return
	}

	helpers.CustomLog().PrintInfo(ctx, c, "creating jwt")

	token, err := helpers.CustomHash().CreateJWToken(user.Username)
	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error generating token:  %s. Error: %s", user.Username, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, "error generating token")
		return
	}

	loginOutput := dto.LoginOutput{
		Username: user.Username,
		Message:  "you are logged..!",
		Token:    token,
	}

	helpers.JsonResponse().JsonSuccess(c, loginOutput)
	helpers.CustomLog().PrintInfo(ctx, c, fmt.Sprintf("user logged successfully: %s", loginDTO.Username))
}

// FindMerchantByCode: retrieve merchant data from repository
func (m *merchantServiceImpl) FindMerchantByCode(ctx context.Context, c *gin.Context) {
	helpers.CustomLog().PrintInfo(ctx, c, "start to execute find merchant by code method..")

	helpers.CustomLog().PrintInfo(ctx, c, "finding merchant by code..")
	merchantCode := c.Params.ByName("merchant_code")
	merchant, findMerchant, err := m.merchantRepository.FindMerchantByCode(merchantCode)

	if err != nil {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("error finding merchant %s. Error: %s", merchantCode, err.Error()), false)
		helpers.JsonResponse().JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding merchant %s", merchantCode))
		return
	}

	if !findMerchant {
		helpers.CustomLog().PrintError(ctx, c, fmt.Sprintf("merchant not found %s.", merchantCode), false)
		helpers.JsonResponse().JsonFail(c, http.StatusNotFound, "merchant not found")
		return
	}

	merchantDTO := dto.Merchant{
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	helpers.JsonResponse().JsonSuccess(c, merchantDTO)
	helpers.CustomLog().PrintInfo(ctx, c, "execution of find merchant by code finished.")
}
