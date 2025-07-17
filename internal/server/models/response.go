package models

import "payment/internal/server/utils"

type AccountResponse struct {
	AccountID            int     `json:"account_id"`
	DocumentNumber       string  `json:"document_number"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}

type TransactionResponse struct {
	TxnId                int           `json:"transaction_id"`
	AccountID            int           `json:"account_id"`
	OperationTypeId      OperationType `json:"operation_type_id"`
	Amount               float64       `json:"amount"`
	AvailableCreditLimit *float64      `json:"available_credit_limit,omitempty"`
}

type ErrorResponse struct {
	ErrorCode utils.ErrorCodes `json:"error_code"`
	ErrorMsg  utils.ErrorMsg   `json:"error_msg"`
}

type InsufficientCreditTransactionResponse struct {
	Error                ErrorResponse `json:"error"`
	AccountId            int           `json:"account_id"`
	AvailableCreditLimit float64       `json:"available_credit_limit"`
}
