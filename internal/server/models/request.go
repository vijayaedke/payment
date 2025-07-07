package models

type AccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type OperationType int

const (
	PURCHASE OperationType = 1
	INSTALLMENT_PURCHASE OperationType = 2
	WITHDRAWAL OperationType = 3
	PAYMENT OperationType = 4
)

type TransactionRequest struct {
	AccountId       int           `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
}

var AccountDataDetails map[string]*AccountResponse
