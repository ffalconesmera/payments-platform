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
	SingUp(ctxt context.Context, c *gin.Context)
	Login(ctxt context.Context, c *gin.Context)
	FindMerchantByCode(ctxt context.Context, c *gin.Context)
}

type merchantServiceImpl struct {
	userRepository     repository.UserRepository
	merchantRepository repository.MerchantRepository
}

func NewMerchantService(userRepository *repository.UserRepository, merchantRepository *repository.MerchantRepository) MerchantService {
	return &merchantServiceImpl{
		userRepository:     *userRepository,
		merchantRepository: *merchantRepository,
	}
}

// SingUp: register a new merchant
func (m *merchantServiceImpl) SingUp(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute sing up method..")

	helpers.PrintInfo(ctxt, c, "mapping merchant data..")
	var merchant dto.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error mapping merchant information. Error: %s", err), false)
		helpers.JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	helpers.PrintInfo(ctxt, c, "validating merchant data..")
	if helpers.EmptyString(merchant.Name) {
		helpers.PrintError(ctxt, c, "merchant name could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "merchant name could not be empty")
		return
	}

	if helpers.EmptyString(merchant.User.Username) {
		helpers.PrintError(ctxt, c, "username could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "username could not be empty")
		return
	}

	if helpers.EmptyString(merchant.User.Email) {
		helpers.PrintError(ctxt, c, "email could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "email could not be empty")
		return
	}

	if helpers.PasswordInvalid(merchant.User.Password) {
		helpers.PrintError(ctxt, c, helpers.PasswordInvalidMessage(), false)
		helpers.JsonFail(c, http.StatusBadRequest, helpers.PasswordInvalidMessage())
		return
	}

	_, findUser, err := m.userRepository.FindUserByUsername(ctxt, merchant.User.Username)

	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error finding user %s. Error: %s", merchant.User.Username, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding user %s.", merchant.User.Username))
		return
	}

	if findUser {
		helpers.PrintError(ctxt, c, fmt.Sprintf("username %s is alreay exits.", merchant.User.Username), false)
		helpers.JsonFail(c, http.StatusConflict, fmt.Sprintf("username %s is alreay exits", merchant.User.Username))
		return
	}

	helpers.PrintInfo(ctxt, c, "generating password hash..")
	hash := helpers.GenerateHashPassword(merchant.User.Password)
	password := hash

	helpers.PrintInfo(ctxt, c, "parsing dto data to merchant model..")
	merchant.MerchantCode = helpers.NewUUIDString()
	merchantModel := model.PayMerchant{
		UUID:         helpers.NewUUIDString(),
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	helpers.PrintInfo(ctxt, c, "parsing dto data to user model..")
	userModel := model.PayUser{
		UUID:         helpers.NewUUIDString(),
		Username:     merchant.User.Username,
		Email:        merchant.User.Email,
		Password:     password,
		MerchantUUID: merchantModel.UUID,
	}

	helpers.PrintInfo(ctxt, c, "saving merchant data..")
	errInsertMerchant := m.merchantRepository.CreateMerchant(ctxt, &merchantModel)

	if errInsertMerchant != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error saving merchant information. Error: %s", errInsertMerchant.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error saving merchant information")
		return
	}

	helpers.PrintInfo(ctxt, c, "saving user data..")
	errUserMerchant := m.userRepository.CreateUser(ctxt, &userModel)

	if errUserMerchant != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error saving merchant information. Error: %s", errUserMerchant.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error saving merchant information")
		return
	}

	merchant.User.Password = ""

	helpers.JsonSuccess(c, merchant)
	helpers.PrintInfo(ctxt, c, "execution of sing up finished.")
}

// Login: generate a authorization token for access as merchant
func (m *merchantServiceImpl) Login(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute login method..")

	helpers.PrintInfo(ctxt, c, "mapping login data..")
	var loginDTO dto.LoginInput
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error mapping merchant information. Error: %s", err), false)
		helpers.JsonFail(c, http.StatusBadRequest, "data sent is invalid")
		return
	}

	helpers.PrintInfo(ctxt, c, "validating login data..")
	if helpers.EmptyString(loginDTO.Username) {
		helpers.PrintError(ctxt, c, "username could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "username could not be empty")
		return
	}

	if helpers.EmptyString(loginDTO.Password) {
		helpers.PrintError(ctxt, c, "password could not be empty", false)
		helpers.JsonFail(c, http.StatusBadRequest, "password could not be empty")
		return
	}

	helpers.PrintInfo(ctxt, c, "finding user by username..")
	user, findUser, err := m.userRepository.FindUserByUsername(ctxt, loginDTO.Username)

	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error finding user %s. Error: %s", loginDTO.Username, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding user: %s", loginDTO.Username))
		return
	}

	if !findUser {
		helpers.PrintError(ctxt, c, fmt.Sprintf("user %s not found.", loginDTO.Username), false)
		helpers.JsonFail(c, http.StatusNotFound, fmt.Sprintf("user %s not found", loginDTO.Username))
		return
	}

	if !helpers.CheckHashPassword(user.Password, loginDTO.Password) {
		helpers.PrintError(ctxt, c, fmt.Sprintf("password incorrect:  %s", loginDTO.Username), false)
		helpers.JsonFail(c, http.StatusUnauthorized, "password incorrect")
		return
	}

	helpers.PrintInfo(ctxt, c, "creating jwt")

	token, err := helpers.CreateJWToken(user.Username)
	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error generating token:  %s. Error: %s", user.Username, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, "error generating token")
		return
	}

	loginOutput := dto.LoginOutput{
		Username: user.Username,
		Message:  "you are logged..!",
		Token:    token,
	}

	helpers.JsonSuccess(c, loginOutput)
	helpers.PrintInfo(ctxt, c, fmt.Sprintf("user logged successfully: %s", loginDTO.Username))
}

// FindMerchantByCode: retrieve merchant data from repository
func (m *merchantServiceImpl) FindMerchantByCode(ctxt context.Context, c *gin.Context) {
	helpers.PrintInfo(ctxt, c, "start to execute find merchant by code method..")

	helpers.PrintInfo(ctxt, c, "finding merchant by code..")
	merchantCode := c.Params.ByName("merchant_code")
	merchant, findMerchant, err := m.merchantRepository.FindMerchantByCode(ctxt, merchantCode)

	if err != nil {
		helpers.PrintError(ctxt, c, fmt.Sprintf("error finding merchant %s. Error: %s", merchantCode, err.Error()), false)
		helpers.JsonFail(c, http.StatusInternalServerError, fmt.Sprintf("error finding merchant %s", merchantCode))
		return
	}

	if !findMerchant {
		helpers.PrintError(ctxt, c, fmt.Sprintf("merchant not found %s.", merchantCode), false)
		helpers.JsonFail(c, http.StatusNotFound, "merchant not found")
		return
	}

	merchantDTO := dto.Merchant{
		MerchantCode: merchant.MerchantCode,
		Name:         merchant.Name,
	}

	helpers.JsonSuccess(c, merchantDTO)
	helpers.PrintInfo(ctxt, c, "execution of find merchant by code finished.")
}
