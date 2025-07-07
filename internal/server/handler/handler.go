package handler

import (
	"payment/internal/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service service.Service
	logger  *zap.SugaredLogger
}

type Controller interface {
	GetAccountDetails(ctx *gin.Context)
	CreateAccount(ctx *gin.Context)
	CreateTransaction(ctx *gin.Context)
}

func InitController(service service.Service, logger *zap.SugaredLogger) Controller {
	return &Handler{
		service: service,
		logger:  logger,
	}
}
