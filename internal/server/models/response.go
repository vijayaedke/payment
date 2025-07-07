package models

import "payment/internal/server/utils"

type AccountResponse struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

type TransactionResponse struct {
	TxnId           int           `json:"transaction_id"`
	AccountID       int           `json:"account_id"`
	OperationTypeId OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
}

type ErrorResponse struct {
	ErrorCode utils.ErrorCodes `json:"error_code"`
	ErrorMsg  utils.ErrorMsg   `json:"error_msg"`
}
