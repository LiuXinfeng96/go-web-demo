package models

type RegisterReq struct {
	UserName     string `json:"userName"`
	UserPwd      string `json:"userPwd"`
	UserRole     string `json:"userRole"`
	UserNickName string `json:"userNickName"`
	UserPhoneNum string `json:"userPhoneNum"`
	UserEmail    string `json:"userEmail"`
}

type LoginReq struct {
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

type AddDebirsReq struct {
	DebrisId     string  `json:"debrisId"`
	DebrisName   string  `json:"debrisName"`
	DebrisSource string  `json:"debrisSource"`
	Angle        float64 `json:"angle"`
	Speed        float64 `json:"speed"`
	Height       float64 `json:"height"`
	Volume       float64 `json:"volume"`
	Type         string  `json:"type"`
}

type AddInstructionReq struct {
	InstructionId      string `json:"instructionId"`
	DebrisId           string `json:"debrisId"`
	DebrisName         string `json:"debrisName"`
	SatelliteId        string `json:"satelliteId"`
	SatelliteName      string `json:"satelliteName"`
	InstructionContent string `json:"instructionContent"`
}

type AddOrbitReq struct {
	OrbitId                string  `json:"orbitId"`
	OrbitType              string  `json:"orbitType"`
	OrbitEccentricity      float64 `json:"orbitEccentricity"`
	OrbitSemiMajorAxis     float64 `json:"orbitSemiMajorAxis"`
	OrbitAngle             float64 `json:"orbitAngle"`
	AscendingNodeLongitude float64 `json:"ascendingNodeLongitude"`
	Perigee                float64 `json:"perigee"`
}

type AddConstellationReq struct {
	ConstellationId    string `json:"constellationId"`
	ConstellationName  string `json:"constellationName"`
	SatelliteTotalNum  int32  `json:"satelliteTotalNum"`
	SatelliteUpNum     int32  `json:"satelliteUpNum"`
	SatelliteDownNum   int32  `json:"satelliteDownNum"`
	SatelliteLinkState string `json:"satelliteLinkState"`
}

type AddSatelliteState struct {
	SatelliteId   string  `json:"satelliteId"`
	SatelliteName string  `json:"satelliteName"`
	OrbitId       string  `json:"orbitId"`
	MeanAnomaly   float64 `json:"meanAnomaly"`
	Speed         float64 `json:"speed"`
	RunState      string  `json:"runState"`
}

type AddControlsReq struct {
	SatelliteId          string  `json:"satelliteId"`
	SatelliteName        string  `json:"satelliteName"`
	SatelliteAttitude    string  `json:"satelliteAttitude"`
	SatelliteTemperature float64 `json:"satelliteTemperature"`
	SatellitePower       string  `json:"satellitePower"`
}

type AddFaultReq struct {
	SatelliteId      string `json:"satelliteId"`
	SatelliteName    string `json:"satelliteName"`
	OrbitId          string `json:"orbitId"`
	FaultType        string `json:"faultType"`
	FaultDescription string `json:"faultDescription"`
	RepairState      string `json:"repairState"`
}

type AddNetStateReq struct {
	SatelliteId      string `json:"satelliteId"`
	SatelliteName    string `json:"satelliteName"`
	OrbitId          string `json:"orbitId"`
	NetworkSegment   string `json:"networkSegment"`
	NetworkState     string `json:"networkState"`
	NetworkBandwidth string `json:"networkBandwidth"`
}

type AddCommStateReq struct {
	SatelliteId   string `json:"satelliteId"`
	SatelliteName string `json:"satelliteName"`
	OrbitId       string `json:"orbitId"`
	CommState     string `json:"commState"`
	CommPort      string `json:"commPort"`
	CommBandwidth string `json:"commBandwidth"`
	CommDelay     string `json:"commDelay"`
	LinkLoad      string `json:"linkLoad"`
}
