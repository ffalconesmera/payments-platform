package service

import (
	"context"
	"errors"
	mock_repository "github.com/ffalconesmera/payments-platform/merchants/mocks/repository"
	"github.com/ffalconesmera/payments-platform/merchants/model"
	"github.com/ffalconesmera/payments-platform/merchants/model/dto"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSingUp(t *testing.T) {
	t.Run("failed in method FindMerchantByCode", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(&model.PayUser{}, false, errors.New("merchant not found"))
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.SingUp(context.TODO(), dto.Merchant{
			User: &dto.User{
				Username: "",
			},
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindMerchantByCode is finded", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(nil, true, nil)
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.SingUp(context.TODO(), dto.Merchant{
			User: &dto.User{
				Username: "",
			},
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method CreateMerchant", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(&model.PayUser{}, false, nil)
		merchantRepository.EXPECT().CreateMerchant(context.TODO(), gomock.Any()).Return(errors.New("error saving merchant"))
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.SingUp(context.TODO(), dto.Merchant{
			User: &dto.User{
				Username: "",
			},
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method CreateMerchant", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(&model.PayUser{}, false, nil)
		merchantRepository.EXPECT().CreateMerchant(context.TODO(), gomock.Any()).Return(errors.New("error saving merchant"))
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.SingUp(context.TODO(), dto.Merchant{
			User: &dto.User{
				Username: "",
			},
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method CreateUser", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(&model.PayUser{}, false, nil)
		merchantRepository.EXPECT().CreateMerchant(context.TODO(), gomock.Any()).Return(nil)
		userRepository.EXPECT().CreateUser(context.TODO(), gomock.Any()).Return(errors.New("error saving user"))
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.SingUp(context.TODO(), dto.Merchant{
			User: &dto.User{
				Username: "",
			},
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("successful", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(&model.PayUser{}, false, nil)
		merchantRepository.EXPECT().CreateMerchant(context.TODO(), gomock.Any()).Return(nil)
		userRepository.EXPECT().CreateUser(context.TODO(), gomock.Any()).Return(nil)
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.SingUp(context.TODO(), dto.Merchant{
			User: &dto.User{
				Username: "",
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}

func TestLogin(t *testing.T) {
	t.Run("failed in method FindUserByUsername", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(nil, false, errors.New("merchant not found"))
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.Login(context.TODO(), dto.LoginInput{
			Username: "",
			Password: "",
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindUserByUsername is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(nil, false, nil)
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.Login(context.TODO(), dto.LoginInput{
			Username: "",
			Password: "",
		})

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in test by secret key not defined", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		userRepository.EXPECT().FindUserByUsername(context.TODO(), "").Return(&model.PayUser{
			Username: "",
			Password: "",
		}, true, nil)
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.Login(context.TODO(), dto.LoginInput{
			Username: "",
			Password: "",
		})

		assert.NotNil(t, err)
		assert.Nil(t, response)
	})
}

func TestFindMerchantByCode(t *testing.T) {
	t.Run("failed in method FindMerchantByCode", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode(context.TODO(), "").Return(nil, false, errors.New("merchant not found"))
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.FindMerchantByCode(context.TODO(), "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("failed in method FindMerchantByCode is false", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode(context.TODO(), "").Return(nil, false, nil)
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.FindMerchantByCode(context.TODO(), "")

		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("successful", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		merchantRepository := mock_repository.NewMockMerchantRepository(ctrl)
		userRepository := mock_repository.NewMockUserRepository(ctrl)
		merchantRepository.EXPECT().FindMerchantByCode(context.TODO(), "").Return(&model.PayMerchant{MerchantCode: "", Name: ""}, true, nil)
		merchantService := NewMerchantService(userRepository, merchantRepository)
		response, err := merchantService.FindMerchantByCode(context.TODO(), "")

		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
