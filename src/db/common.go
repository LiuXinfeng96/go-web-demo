package db

var TableSlice = make([]interface{}, 0)

type GeneralField struct {
	Id       int32 `gorm:"primaryKey;autoIncrement"`
	LastTime int64 `gorm:"autoCreateTime:milli;index"`
}

type ModelStruct interface {
	TableName() string
}

type State int32

const (
	NORMAL State = iota + 1
	WRONG
)

const (
	NORMAL_STR = "正常"
	WRONG_STR  = "异常"
)

var StateName = map[State]string{
	NORMAL: NORMAL_STR,
	WRONG:  WRONG_STR,
}

var StateValue = map[string]State{
	NORMAL_STR: NORMAL,
	WRONG_STR:  WRONG,
}
