package handler

import (
	"net/http"
	"payment/internal/server/models"
	"payment/internal/server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var accountData models.AccountRequest

	if err := ctx.BindJSON(&accountData); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorCode: utils.INVALID_ACC_REQUEST_PAYLOAD,
			ErrorMsg:  utils.INVALID_ACC_REQUEST_PAYLOAD_MSG,
		})
		return
	}

	resp, err := h.service.CreateAccount(ctx, &accountData)
	if resp == nil || err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorCode: utils.CREATE_ACC_FAILED,
			ErrorMsg:  utils.CREATE_ACC_FAILED_MSG,
		})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, resp)
}

func (h *Handler) GetAccountDetails(ctx *gin.Context) {
	pathParam := ctx.Param("accountId")
	if len(pathParam) == 0 {
		h.logger.Error("Invalid path param format", zap.String("path_param", pathParam))
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorCode: utils.INVALID_ACC_ID,
			ErrorMsg:  utils.INVALID_ACC_ID_MSG,
		})
		return
	}

	accountID, err := strconv.Atoi(pathParam)
	if err != nil {
		h.logger.Error("Failed strconv.Atoi ", zap.Error(err), zap.String("account_id_path_param", pathParam))
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrorCode: utils.INVALID_ACC_ID,
			ErrorMsg:  utils.INVALID_ACC_ID_MSG,
		})
		return
	}

	resp, err := h.service.GetAccountDetailsById(ctx, accountID)
	h.logger.Error("Account details not found ", zap.Error(err), zap.Int("account_id", accountID))
	if resp == nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{
			ErrorCode: utils.ACC_NOT_FOUND,
			ErrorMsg:  utils.ACC_NOT_FOUND_MSG,
		})
		return
	}

	if err != nil {
		h.logger.Error("Failed GetAccountDetailsById ", zap.Error(err), zap.String("account_id", pathParam))
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrorCode: utils.FAILED_ACC_ID_DETAIL,
			ErrorMsg:  utils.FAILED_ACC_ID_DETAIL_MSG,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, resp)
}
