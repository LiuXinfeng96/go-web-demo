package models

type StandardResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type StandardRespWithPage struct {
	StandardResp
	Total int64 `json:"total"`
}

const (
	RESP_CODE_SUCCESS = 200

	RESP_CODE_SERVER_ERROR = 500

	RESP_CODE_UNIQUE_INDEX = 501

	RESP_CODE_NOT_EXIST = 502

	RESP_CODE_PWD_ERROR = 503

	RESP_CODE_PARAMS_TYPE_ERROR = 101

	RESP_CODE_PARAMS_MISSING = 102

	RESP_CODE_PARAMS_FORMAT_ERROR = 103

	RESP_CODE_PARAMS_VALUE_ERROR = 104

	RESP_CODE_WITHOUT_PERMISSION = 301

	RESP_CODE_NOT_IN_CHAIN = 302

	RESP_CODE_TOKEN_ERROR = 303

	RESP_CODE_CHAIN_PUT_FAIL = 600
)

const (
	RESP_MSG_SERVER_ERROR = "服务出错"

	RESP_MSG_PARAMS_TYPE_ERROR = "错误的参数类型"

	RESP_MSG_PARAMS_MISSING = "缺少必要参数"

	RESP_MSG_PARAMS_FORMAT_ERROR = "参数不符合规范"

	RESP_MSG_PARAMS_VALUE_ERROR = "参数值域不符合预期"

	RESP_MSG_WITHOUT_PERMISSION = "无权限访问"

	RESP_MSG_UNIQUE_INDEX = "唯一索引已经存在"

	RESP_MSG_NOT_EXIST = "唯一索引不存在"

	RESP_MSG_PWD_ERROR = "密码错误"

	RESP_MSG_NOT_IN_CHAIN = "不属于链上用户，无法操作链"

	RESP_MSG_TOKEN_ERROR = "请重新登录"

	RESP_MSG_CHAIN_PUT_FAIL = "上链失败"
)

type UserInfo struct {
	Id           int32  `json:"id"`
	UserName     string `json:"userName"`
	UserNickName string `json:"userNickName"`
	UserRole     string `json:"userRole"`
	UserPhoneNum string `json:"userPhoneNum"`
	UserEmail    string `json:"userEmail"`
	Expires      int64  `json:"expires"`
}

type LoginInfo struct {
	UserNickName string `json:"userNickName"`
	UserRole     string `json:"userRole"`
	Expires      int64  `json:"expires"`
	Token        string `json:"token"`
}

type BaseRespInfo struct {
	Id       int32 `json:"id"`
	LastTime int64 `json:"lastTime"`
}

type DebrisInfo struct {
	DebrisId     string  `json:"debrisId"`
	DebrisName   string  `json:"debrisName"`
	DebrisSource string  `json:"debrisSource"`
	Angle        float64 `json:"angle"`
	Speed        float64 `json:"speed"`
	Height       float64 `json:"height"`
	Volume       float64 `json:"volume"`
	Type         string  `json:"type"`
	BaseRespInfo
}

type InstructionInfo struct {
	BaseRespInfo
	InstructionId       string `json:"instructionId"`
	InstructionType     string `json:"instructionType"`
	ExecInstructionTime int64  `json:"execInstructionTime"`
	InstructionSource   string `json:"instructionSource"`
	DebrisId            string `json:"debrisId"`
	DebrisName          string `json:"debrisName"`
	SatelliteId         string `json:"satelliteId"`
	SatelliteName       string `json:"satelliteName"`
	InstructionContent  string `json:"instructionContent"`
	GenInstructionTime  int64  `json:"genInstructionTime"`
}

type InstructionDetails struct {
	InstructionInfo
	ExecState string `json:"execState"`
}

type OrbitInfo struct {
	BaseRespInfo
	OrbitId                string  `json:"orbitId"`
	OrbitType              string  `json:"orbitType"`
	OrbitEccentricity      float64 `json:"orbitEccentricity"`
	OrbitSemiMajorAxis     float64 `json:"orbitSemiMajorAxis"`
	OrbitAngle             float64 `json:"orbitAngle"`
	AscendingNodeLongitude float64 `json:"ascendingNodeLongitude"`
	Perigee                float64 `json:"perigee"`
}

type ConstellationInfo struct {
	BaseRespInfo
	ConstellationId    string `json:"constellationId"`
	ConstellationName  string `json:"constellationName"`
	SatelliteTotalNum  int32  `json:"satelliteTotalNum"`
	SatelliteUpNum     int32  `json:"satelliteUpNum"`
	SatelliteDownNum   int32  `json:"satelliteDownNum"`
	SatelliteLinkState string `json:"satelliteLinkState"`
}

type OperationInfo struct {
	BaseRespInfo
	Operator        string `json:"operator"`
	OperationTime   int64  `json:"operationTime"`
	OperationIp     string `json:"operationIp"`
	OperationRecord string `json:"operationRecord"`
	SatelliteId     string `json:"satelliteId"`
	SatelliteName   string `json:"satelliteName"`
}

type LoginLogInfo struct {
	BaseRespInfo
	UserName  string `json:"userName"`
	LoginTime int64  `json:"loginTime"`
	LoginIp   string `json:"loginIp"`
}

type SatelliteStateInfo struct {
	BaseRespInfo
	SatelliteId   string  `json:"satelliteId"`
	SatelliteName string  `json:"satelliteName"`
	OrbitId       string  `json:"orbitId"`
	MeanAnomaly   float64 `json:"meanAnomaly"`
	Speed         float64 `json:"speed"`
	RunState      string  `json:"runState"`
}

type ControlsInfo struct {
	BaseRespInfo
	SatelliteId          string  `json:"satelliteId"`
	SatelliteName        string  `json:"satelliteName"`
	SatelliteAttitude    string  `json:"satelliteAttitude"`
	SatelliteTemperature float64 `json:"satelliteTemperature"`
	SatellitePower       string  `json:"satellitePower"`
}

type FaultInfo struct {
	BaseRespInfo
	SatelliteId      string `json:"satelliteId"`
	SatelliteName    string `json:"satelliteName"`
	OrbitId          string `json:"orbitId"`
	FaultType        string `json:"faultType"`
	FaultDescription string `json:"faultDescription"`
	FaultTime        int64  `json:"faultTime"`
	RepairState      string `json:"repairState"`
}

type NetStateInfo struct {
	BaseRespInfo
	SatelliteId      string `json:"satelliteId"`
	SatelliteName    string `json:"satelliteName"`
	OrbitId          string `json:"orbitId"`
	NetworkSegment   string `json:"networkSegment"`
	NetworkState     string `json:"networkState"`
	NetworkBandwidth string `json:"networkBandwidth"`
}

type CommStateInfo struct {
	BaseRespInfo
	SatelliteId   string `json:"satelliteId"`
	SatelliteName string `json:"satelliteName"`
	OrbitId       string `json:"orbitId"`
	CommState     string `json:"commState"`
	CommPort      string `json:"commPort"`
	CommBandwidth string `json:"commBandwidth"`
	CommDelay     string `json:"commDelay"`
	LinkLoad      string `json:"linkLoad"`
}

type AllState struct {
	FaultState         bool `json:"faultState"`
	SatelliteState     bool `json:"satelliteState"`
	ConstellationState bool `json:"constellationState"`
	NetState           bool `json:"netState"`
	CommState          bool `json:"commState"`
}

type ChainDataNum struct {
	DebrisNum         int64 `json:"debrisNum"`
	InstructionNum    int64 `json:"instructionNum"`
	SatelliteStateNum int64 `json:"satelliteStateNum"`
	OrbitNum          int64 `json:"orbitNum"`
	ConstellationNum  int64 `json:"constellationNum"`
	ControlNum        int64 `json:"controlNum"`
	FaultNum          int64 `json:"faultNum"`
	NetStateNum       int64 `json:"netStateNum"`
	CommStateNum      int64 `json:"commStateNum"`
	OperationNum      int64 `json:"operationNum"`
	LoginLogNum       int64 `json:"loginLogNum"`
}

type FaultTypeInfo struct {
	AType int64 `json:"atype"`
	BType int64 `json:"btype"`
	CType int64 `json:"ctype"`
	DType int64 `json:"dtype"`
}

type EarlyWarningInfo struct {
	DebrisId      string  `json:"debrisId"`
	DebrisName    string  `json:"debrisName"`
	Speed         float64 `json:"speed"`
	Height        float64 `json:"height"`
	Treaten       int32   `json:"threaten"`
	SatelliteName string  `json:"satelliteName"`
	SatelliteId   string  `json:"satelliteId"`
}

type ChainInfo struct {
	NodeNum     int    `json:"nodeNum"`
	BlockHeight uint64 `json:"blockHeight"`
	ChainName   string `json:"chainName"`
}

type LatestInfo struct {
	IsLatest bool `json:"isLatest"`
}

type LastTime struct {
	LastTime int64 `json:"serverLastTime"`
}
