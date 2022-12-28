package db

const LOGINLOG_TABLE_NAME = "login_log"

type LoginLog struct {
	GeneralField
	UserName  string `gorm:"index"`
	LoginTime int64
	LoginIp   string
}

func (l *LoginLog) TableName() string {
	return LOGINLOG_TABLE_NAME
}

func init() {
	loginLog := new(LoginLog)
	TableSlice = append(TableSlice, &loginLog)
}
