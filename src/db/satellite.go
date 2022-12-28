package db

const SATELLITE_TABLE_NAME = "satellite"

type Satellite struct {
	GeneralField
	SatelliteId   string `gorm:"index"`
	SatelliteName string `gorm:"index"`
	OrbitId       string
	MeanAnomaly   float64
	Speed         float64
	RunState      State
}

func (s *Satellite) TableName() string {
	return SATELLITE_TABLE_NAME
}

func init() {
	satellite := new(Satellite)
	TableSlice = append(TableSlice, &satellite)
}
