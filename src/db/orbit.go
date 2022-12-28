package db

const ORBIT_TABLE_NAME = "orbit"

type Orbit struct {
	GeneralField
	OrbitId                string `gorm:"index"`
	OrbitType              string
	OrbitSemiMajorAxis     float64
	OrbitEccentricity      float64
	OrbitAngle             float64
	AscendingNodeLongitude float64
	Perigee                float64
}

func (o *Orbit) TableName() string {
	return ORBIT_TABLE_NAME
}

func init() {
	orbit := new(Orbit)
	TableSlice = append(TableSlice, &orbit)
}
