package db

const USER_TABLE_NAME = "user"

type UserRoleType int32

const (
	CONTROL UserRoleType = iota + 1
	EXEC
	TRACE
	MONITOR
)

const (
	CONTROL_STR = "控制系统"

	EXEC_STR = "执行系统"

	TRACE_STR = "溯源系统"

	MONITOR_STR = "监控系统"
)

var UserRoleTypeName = map[UserRoleType]string{
	CONTROL: CONTROL_STR,
	EXEC:    EXEC_STR,
	TRACE:   TRACE_STR,
	MONITOR: MONITOR_STR,
}

var UserRoleTypeValue = map[string]UserRoleType{
	CONTROL_STR: CONTROL,
	EXEC_STR:    EXEC,
	TRACE_STR:   TRACE,
	MONITOR_STR: MONITOR,
}

type User struct {
	GeneralField
	UserName     string `gorm:"uniqueIndex"`
	UserRole     UserRoleType
	UserPwd      string
	UserNickName string
	UserPhoneNum string
	UserEmail    string
}

func (u *User) TableName() string {
	return USER_TABLE_NAME
}

func init() {
	user := new(User)
	TableSlice = append(TableSlice, &user)
}
