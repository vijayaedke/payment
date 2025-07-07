package service

import (
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
		accountIDExists, operationIDExists bool
		wg                                 sync.WaitGroup
		accountData                        entities.Account
		operationData                      entities.OperationsTypes
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		accountIDExists = s.mysqlClient.Exists(&accountData, utils.ACCOUNT_ID_PARAM, txnReq.AccountId)
	}()

	go func() {
		defer wg.Done()
		operationIDExists = s.mysqlClient.Exists(&operationData, utils.OPERTAION_ID_PARAM, txnReq.OperationTypeID)
	}()

	wg.Wait()

	if !accountIDExists {
		s.logger.Error("Failed mysqlClient.Exists: Provided account id doesn't exists", zap.Int("account_id", txnReq.AccountId))
		return nil, utils.AccIdNotExists
	}

	if !operationIDExists {
		s.logger.Error("Failed mysqlClient.Exists: Provided account id doesn't exists", zap.Int("operation_id", int(txnReq.OperationTypeID)))
		return nil, utils.OpIdNotExists
	}

	amount := txnReq.Amount
	if txnReq.OperationTypeID != models.PAYMENT {
		amount = -txnReq.Amount
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
