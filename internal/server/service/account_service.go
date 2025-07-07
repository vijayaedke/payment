package service

import (
	"payment/internal/server/models"
	"payment/internal/server/models/entities"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *PaymentService) CreateAccount(ctx *gin.Context, accountData *models.AccountRequest) (*models.AccountResponse, error) {
	resp, err := s.mysqlClient.Create(&entities.Account{
		DocumentNumber: accountData.DocumentNumber,
	})
	if err != nil {
		s.logger.Error("Failed mysqlClient.Create", zap.Error(err), zap.String("document_number", accountData.DocumentNumber))
		return nil, err
	}
	accountEntity := resp.(*entities.Account)
	accountResponse := &models.AccountResponse{
		AccountID:      accountEntity.AccountID,
		DocumentNumber: accountEntity.DocumentNumber,
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
