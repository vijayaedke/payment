package service

import (
	"payment/internal/server/models"
	"payment/internal/server/models/entities"
	"payment/internal/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *PaymentService) CreateAccount(ctx *gin.Context, accountData *models.AccountRequest) (*models.AccountResponse, error) {
	var availableCreditLimit = utils.DEFAULT_AVAILABLE_CREDIT_LIMIT
	if accountData.AvailableCreditLimit != nil {
		availableCreditLimit = *accountData.AvailableCreditLimit
	}
	resp, err := s.mysqlClient.Create(&entities.Account{
		DocumentNumber:       accountData.DocumentNumber,
		AvailableCreditLimit: availableCreditLimit,
	})
	if err != nil {
		s.logger.Error("Failed mysqlClient.Create", zap.Error(err), zap.String("document_number", accountData.DocumentNumber))
		return nil, err
	}
	accountEntity := resp.(*entities.Account)
	accountResponse := &models.AccountResponse{
		AccountID:            accountEntity.AccountID,
		DocumentNumber:       accountEntity.DocumentNumber,
		AvailableCreditLimit: accountEntity.AvailableCreditLimit,
	}

	return accountResponse, nil
}

func (s *PaymentService) GetAccountDetailsById(ctx *gin.Context, accountID int) (*models.AccountResponse, error) {
	var accountInfo entities.Account
	resp, err := s.mysqlClient.FindOne(&accountInfo, accountID)
	if err != nil {
		s.logger.Error("Failed mysqlClient.FindOne", zap.Error(err), zap.Int("account_id", accountID))
		return nil, err
	}

	accountEntity := resp.(*entities.Account)
	accountResponse := &models.AccountResponse{
		AccountID:      accountEntity.AccountID,
		DocumentNumber: accountEntity.DocumentNumber,
	}

	return accountResponse, nil
}
