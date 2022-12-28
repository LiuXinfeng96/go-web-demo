package db

const NETSTATE_TABLE_NAME = "net_state"

type NetState struct {
	GeneralField
	SatelliteId      string `gorm:"index"`
	SatelliteName    string `gorm:"index"`
	OrbitId          string
	NetworkSegment   string
	NetworkState     State
	NetworkBandwidth string
}

func (n *NetState) TableName() string {
	return NETSTATE_TABLE_NAME
}

func init() {
	netState := new(NetState)
	TableSlice = append(TableSlice, &netState)
}
