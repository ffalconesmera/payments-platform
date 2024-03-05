package controller

import (
	"context"

	"github.com/ffalconesmera/payments-platform/merchants/helpers"
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
	FindMerchantById(c *gin.Context)
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

func (mc *merchantControllerImpl) Login(c *gin.Context) {
	loginCtxt := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	mc.merchantService.Login(loginCtxt, c)
}

func (mc *merchantControllerImpl) SingUp(c *gin.Context) {
	singUpCtxt := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	mc.merchantService.SingUp(singUpCtxt, c)
}

func (mc *merchantControllerImpl) FindMerchantByCode(c *gin.Context) {
	findMerchantCtxt := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	mc.merchantService.FindMerchantByCode(findMerchantCtxt, c)
}

func (mc *merchantControllerImpl) FindMerchantById(c *gin.Context) {
	findMerchantCtxt := context.WithValue(c, "REQUEST_ID", helpers.NewUUIDString())
	mc.merchantService.FindMerchantByCode(findMerchantCtxt, c)
}
