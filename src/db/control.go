package db

const CONTROL_TABLE_NAME = "control"

type Control struct {
	GeneralField
	SatelliteId          string `gorm:"index"`
	SatelliteName        string `gorm:"index"`
	SatelliteAttitude    string
	SatelliteTemperature float64
	SatellitePower       string
}

func (c *Control) TableName() string {
	return CONTROL_TABLE_NAME
}

func init() {
	control := new(Control)
	TableSlice = append(TableSlice, &control)
}
