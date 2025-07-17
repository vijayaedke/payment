package service

import (
	"fmt"
	"sync"
	"time"

	"payment/internal/server/models"
	"payment/internal/server/models/entities"
	"payment/internal/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *PaymentService) CreateTransaction(ctx *gin.Context, txnReq *models.TransactionRequest) (*models.TransactionResponse, error) {
	var (
		operationIDExists bool
		wg                sync.WaitGroup
		accountData       entities.Account
		operationData     entities.OperationsTypes
		response          interface{}
		err               error
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		response, err = s.mysqlClient.FindOne(&accountData, txnReq.AccountId)
	}()

	go func() {
		defer wg.Done()
		operationIDExists = s.mysqlClient.Exists(&operationData, utils.OPERTAION_ID_PARAM, txnReq.OperationTypeID)
	}()

	wg.Wait()

	accountDetails := response.(*entities.Account)
	if response == nil || accountDetails.AccountID != txnReq.AccountId {
		s.logger.Error("Failed mysqlClient.Exists: Provided account id doesn't exists", zap.Int("account_id", txnReq.AccountId))
		return nil, utils.AccIdNotExists
	}

	if !operationIDExists {
		s.logger.Error("Failed mysqlClient.Exists: Provided account id doesn't exists", zap.Int("operation_id", int(txnReq.OperationTypeID)))
		return nil, utils.OpIdNotExists
	}

	amount := txnReq.Amount
	if amount < 0 {
		s.logger.Error("Failed operation type id with incorrect amount format ", zap.Float64("amount", amount), zap.Int("account_id", txnReq.AccountId))
		return nil, fmt.Errorf("incorrect amount format")
	}

	if txnReq.OperationTypeID != models.PAYMENT {
		amount = -txnReq.Amount
	}

	fmt.Println("balance = ", accountDetails.AvailableCreditLimit)
	if txnReq != nil && accountDetails.AvailableCreditLimit < txnReq.Amount {
		s.logger.Error("Insufficient available credit limit", zap.Float64("amount", amount), zap.Int("account_id", txnReq.AccountId),
			zap.Float64("available_credit_limit", accountDetails.AvailableCreditLimit))
		return &models.TransactionResponse{
			AccountID:            txnReq.AccountId,
			AvailableCreditLimit: &accountDetails.AvailableCreditLimit,
		}, utils.InsufficientCredit
	}

	var updatedAccountDetails = entities.Account{
		AvailableCreditLimit: accountDetails.AvailableCreditLimit - txnReq.Amount,
	}
	updateCredit := s.mysqlClient.Update(ctx, accountData, "account_id", txnReq.AccountId, updatedAccountDetails)
	if updateCredit != nil {
		s.logger.Error("Failed to Update the available credit limit", zap.Error(updateCredit), zap.Float64("amount", amount), zap.Int("account_id", txnReq.AccountId),
			zap.Float64("available_credit_limit", accountDetails.AvailableCreditLimit))
		return &models.TransactionResponse{
			AccountID:            txnReq.AccountId,
			AvailableCreditLimit: &accountDetails.AvailableCreditLimit,
		}, utils.InsufficientCredit
	}

	resp, err := s.mysqlClient.Create(&entities.Transaction{
		AccountID:       txnReq.AccountId,
		OperationTypeID: txnReq.OperationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	})
	if err != nil {
		s.logger.Error("Failed mysqlClient.Create", zap.Error(err), zap.Any("txn_req", txnReq))
		return nil, err
	}
	txnEntity := resp.(*entities.Transaction)
	txnResponse := &models.TransactionResponse{
		TxnId:           txnEntity.TxnID,
		AccountID:       txnEntity.AccountID,
		OperationTypeId: txnEntity.OperationTypeID,
		Amount:          txnEntity.Amount,
	}

	return txnResponse, nil
}
