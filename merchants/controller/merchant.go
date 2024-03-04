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
	Login(ctx context.Context, c *gin.Context)
	SingUp(ctx context.Context, c *gin.Context)
	FindMerchantByCode(ctx context.Context, c *gin.Context)
	FindMerchantById(ctx context.Context, c *gin.Context)
}

type merchantControllerImpl struct {
	merchantService service.MerchantService
	hash            helpers.CustomHash
}

func NewMerchantController(ctx context.Context, merchantService service.MerchantService, hash helpers.CustomHash) *merchantControllerImpl {
	return &merchantControllerImpl{
		merchantService: merchantService,
		hash:            hash,
	}
}

func (m *merchantControllerImpl) Login(ctx context.Context, c *gin.Context) {
	loginCtx := context.WithValue(ctx, "REQUEST_ID", m.hash.NewUUIDString())
	m.merchantService.Login(loginCtx, c)
}

func (m *merchantControllerImpl) SingUp(ctx context.Context, c *gin.Context) {
	singUpCtx := context.WithValue(ctx, "REQUEST_ID", m.hash.NewUUIDString())
	m.merchantService.SingUp(singUpCtx, c)
}

func (m *merchantControllerImpl) FindMerchantByCode(ctx context.Context, c *gin.Context) {
	findMerchantCtx := context.WithValue(ctx, "REQUEST_ID", m.hash.NewUUIDString())
	m.merchantService.FindMerchantByCode(findMerchantCtx, c)
}

func (m *merchantControllerImpl) FindMerchantById(ctx context.Context, c *gin.Context) {
	findMerchantCtx := context.WithValue(ctx, "REQUEST_ID", m.hash.NewUUIDString())
	m.merchantService.FindMerchantByCode(findMerchantCtx, c)
}
