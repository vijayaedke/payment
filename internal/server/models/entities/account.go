package entities

type Account struct {
	AccountID            int     `gorm:"column:account_id;primaryKey;autoIncrement"`
	DocumentNumber       string  `gorm:"column:document_number"`
	AvailableCreditLimit float64 `gorm:"column:available_credit_limit"`
}

func (a *Account) TableName() string {
	return "accounts"
}
