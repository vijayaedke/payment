package entities

type OperationsTypes struct {
	OperationTypeID int    `gorm:"column:operation_type_id"`
	Description     string `gorm:"column:description"`
}

func (a *OperationsTypes) TableName() string {
	return "operation_types"
}
