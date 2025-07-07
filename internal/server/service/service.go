package service

import (
	"payment/internal/server/models"
	"payment/pkg/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentService struct {
	mysqlClient mysql.MysqlClientService
	logger      *zap.SugaredLogger
}

type Service interface {
	CreateAccount(ctx *gin.Context, accountData *models.AccountRequest) (*models.AccountResponse, error)
	GetAccountDetailsById(ctx *gin.Context, accountId int) (*models.AccountResponse, error)
	CreateTransaction(ctx *gin.Context, txnReq *models.TransactionRequest) (*models.TransactionResponse, error)
}

func InitService(mysqlClient mysql.MysqlClientService, logger *zap.SugaredLogger) Service {
	return &PaymentService{
		mysqlClient: mysqlClient,
		logger:      logger,
	}
}
