package router

import (
	"payment/internal/server/handler"
	"payment/internal/server/models"
	"payment/internal/server/service"

	"payment/pkg/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	ginHandler  *gin.Engine
	handler     handler.Controller
	mysqlClient mysql.MysqlClientService
	logger      *zap.SugaredLogger
}

func InitRouter(ginHandler *gin.Engine, logger *zap.SugaredLogger) *Router {
	mysqlClient := mysql.InitMysqlClient()
	service := service.InitService(mysqlClient, logger)
	controller := handler.InitController(service, logger)
	models.AccountDataDetails = make(map[string]*models.AccountResponse)
	return &Router{
		ginHandler:  ginHandler,
		handler:     controller,
		mysqlClient: mysqlClient,
	}
}

func (r *Router) CreateRouterGroups() {
	v1Group := r.ginHandler.Group("/api/v1/pismo")
	v1Group.POST("/account", r.handler.CreateAccount)
	v1Group.GET("/account/:accountId", r.handler.GetAccountDetails)

	v1Group.POST("/transactions", r.handler.CreateTransaction)

}
