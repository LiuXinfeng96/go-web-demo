package db

const FAULT_TABLE_NAME = "fault"

type FaultType int32

const (
	AType FaultType = iota + 1
	BType
	CType
	DType
)

const (
	AType_STR = "1"
	BType_STR = "2"
	CType_STR = "3"
	DType_STR = "4"
)

var FaultTypeName = map[FaultType]string{
	AType: AType_STR,
	BType: BType_STR,
	CType: CType_STR,
	DType: DType_STR,
}

var FaultTypeValue = map[string]FaultType{
	AType_STR: AType,
	BType_STR: BType,
	CType_STR: CType,
	DType_STR: DType,
}

type Fault struct {
	GeneralField
	SatelliteId      string `gorm:"index"`
	SatelliteName    string `gorm:"index"`
	OrbitId          string
	FaultType        FaultType
	FaultTime        int64
	FaultDescription string `gorm:"type:longtext"`
	RepairState      State
}

func (f *Fault) TableName() string {
	return FAULT_TABLE_NAME
}

func init() {
	fault := new(Fault)
	TableSlice = append(TableSlice, &fault)
}
