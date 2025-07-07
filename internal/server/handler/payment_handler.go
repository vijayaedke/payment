package handler

import (
	"net/http"

	"payment/internal/server/models"
	"payment/internal/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) CreateTransaction(ctx *gin.Context) {
	var txnReq models.TransactionRequest

	if err := ctx.BindJSON(&txnReq); err != nil {
		h.logger.Error("Invalid request payload", zap.Error(err), zap.Any("txn_request", txnReq))
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorCode: utils.INVALID_TXN_REQUEST_PAYLOAD,
			ErrorMsg:  utils.INVALID_TXN_REQUEST_PAYLOAD_MSG,
		})
		return
	}

	resp, err := h.service.CreateTransaction(ctx, &txnReq)
	if err == utils.AccIdNotExists {
		h.logger.Error("Provided account id not found", zap.Error(err), zap.Any("txn_request", txnReq))
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrorCode: utils.ACC_NOT_FOUND,
			ErrorMsg:  utils.ACC_NOT_FOUND_MSG,
		})
		return
	}

	if err == utils.OpIdNotExists {
		h.logger.Error("Provided operation id not found", zap.Error(err), zap.Any("txn_request", txnReq))
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrorCode: utils.OP_ID_NOT_FOUND,
			ErrorMsg:  utils.OP_ID_NOT_FOUND_MSG,
		})
		return
	}

	if resp == nil || err != nil {
		h.logger.Error("Failed CreateTransaction ", zap.Error(err), zap.Any("txn_request", txnReq))
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorCode: utils.FAILED_TXN_CREATE,
			ErrorMsg:  utils.FAILED_TXN_CREATE_MSG,
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, resp)
}
