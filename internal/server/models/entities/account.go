package entities

type Account struct {
	AccountID      int    `gorm:"column:account_id;primaryKey;autoIncrement"`
	DocumentNumber string `gorm:"column:document_number"`
}

func (a *Account) TableName() string {
	return "accounts"
}
