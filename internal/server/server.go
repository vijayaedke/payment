package server

import (
	"log"
	"net/http"
	"payment/internal/router"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ServerConfig struct {
	port       int
	ginHandler *gin.Engine
	logger     *zap.SugaredLogger
}

type ServerInterface interface {
	StartServer() error
}

func InitServer(port int, ginHandler *gin.Engine, logger *zap.SugaredLogger) ServerInterface {
	return &ServerConfig{
		port:       port,
		ginHandler: ginHandler,
		logger:     logger,
	}
}

func (s *ServerConfig) StartServer() error {
	router := router.InitRouter(s.ginHandler, s.logger)
	router.CreateRouterGroups()

	srv := &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(s.port),
		Handler: s.ginHandler,
	}
	log.Println("Server started at port", s.port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed ListenAndServe", err)
		return err
	}

	return nil
}
