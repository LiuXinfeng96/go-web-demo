package db

const COMMSTATE_TABLE_NAME = "comm_state"

type CommState struct {
	GeneralField
	SatelliteId   string `gorm:"index"`
	SatelliteName string `gorm:"index"`
	OrbitId       string
	CommState     State
	CommBandwidth string
	CommDelay     string
	CommPort      string
	LinkLoad      string
}

func (c *CommState) TableName() string {
	return COMMSTATE_TABLE_NAME
}

func init() {
	commState := new(CommState)
	TableSlice = append(TableSlice, &commState)
}
