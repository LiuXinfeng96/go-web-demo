package db

const OPERATION_TABLE_NAME = "operation"

type Operation struct {
	GeneralField
	Operator        string `gorm:"index"`
	OperatorIp      string
	OperationTime   int64
	OperationRecord string `gorm:"type:longtext"`
	SatelliteId     string `gorm:"index"`
	SatelliteName   string
}

func (o *Operation) TableName() string {
	return OPERATION_TABLE_NAME
}

func init() {
	operation := new(Operation)
	TableSlice = append(TableSlice, &operation)
}
