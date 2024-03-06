package controller

import (
	"errors"
	"github.com/ffalconesmera/payments-platform/merchants/helpers"
	"github.com/ffalconesmera/payments-platform/merchants/model/dto"
	"github.com/ffalconesmera/payments-platform/merchants/service"
	"github.com/gin-gonic/gin"
)

// MerchantController is an interface to comunicate with service layer and define context for identify each request
// Login: execute Login defined in service
// SingUp: execute SingUp defined in service
// FindMerchantByCode: execute FindMerchantByCode defined in service
// FindMerchantById: execute FindMerchantById defined in service
type MerchantController interface {
	Login(c *gin.Context)
	SingUp(c *gin.Context)
	FindMerchantByCode(c *gin.Context)
}

type merchantControllerImpl struct {
	merchantService service.MerchantService
}

func NewMerchantController(merchantService *service.MerchantService) MerchantController {
	if merchantService == nil {
		return nil
	}

	return &merchantControllerImpl{
		merchantService: *merchantService,
	}
}

func (mc *merchantControllerImpl) SingUp(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())

	var merchant dto.Merchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		helpers.ResponseJson(c, nil, err)
		return
	}

	if helpers.EmptyString(merchant.Name) {
		helpers.ResponseJson(c, nil, errors.New("merchant name could not be empty"))
		return
	}

	if helpers.EmptyString(merchant.User.Username) {
		helpers.ResponseJson(c, nil, errors.New("username could not be empty"))
		return
	}

	if helpers.EmptyString(merchant.User.Email) {
		helpers.ResponseJson(c, nil, errors.New("email user could not be empty"))
		return
	}

	if helpers.PasswordInvalid(merchant.User.Password) {
		helpers.ResponseJson(c, nil, errors.New(helpers.PasswordInvalidMessage()))
		return
	}

	merch, err := mc.merchantService.SingUp(c, merchant)
	helpers.ResponseJson(c, merch, err)
}

func (mc *merchantControllerImpl) Login(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())
	var login dto.LoginInput

	if err := c.ShouldBindJSON(&login); err != nil {
		helpers.ResponseJson(c, nil, errors.New("username could not be zero"))
	}

	if helpers.EmptyString(login.Username) {
		helpers.ResponseJson(c, nil, errors.New("username could not be zero"))
		return
	}

	loginResp, err := mc.merchantService.Login(c, login)

	helpers.ResponseJson(c, loginResp, err)

}

func (mc *merchantControllerImpl) FindMerchantByCode(c *gin.Context) {
	c.AddParam("REQUEST_ID", helpers.NewUUIDString())
	merchantCode := c.Params.ByName("merchant_code")
	merchant, err := mc.merchantService.FindMerchantByCode(c, merchantCode)
	helpers.ResponseJson(c, merchant, err)
}
