package entities

import (
	"time"

	"payment/internal/server/models"
)

type Transaction struct {
	TxnID           int                  `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	AccountID       int                  `gorm:"column:account_id"`
	OperationTypeID models.OperationType `gorm:"column:operation_type_id"`
	Amount          float64              `gorm:"column:amount"`
	EventDate       time.Time            `gorm:"column:event_date"`
}

func (a *Transaction) TableName() string {
	return "transactions"
}
