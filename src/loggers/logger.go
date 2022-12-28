package loggers

import (
	"errors"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志级别，配置文件定义的常量
const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

const (
	DEFAULT_LOG_PATH = "./log/default.log"

	// 默认日志文件切割大小
	DEFAULT_LOG_MAXSIZE = 10

	// 默认最大保留天数
	DEFAULT_LOG_MAXAGE = 30

	// 默认保留日志文件的最大数量
	DEFAULT_LOG_MAXBACKUPS = 1000

	DEFAULT_LOG_LEVEL = INFO
)

// LogConfig 日志记录的配置
type LogConfig struct {
	LogPath string `mapstructure:"log_path"`

	LogLevel string `mapstructure:"log_level"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `mapstructure:"max_age"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `mapstructure:"max_size"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `mapstructure:"max_backups"`

	// jsonFormat: log file use json format
	JsonFormat bool `mapstructure:"json_format"`

	// showLine: show filename and line number
	ShowLine bool `mapstructure:"show_line"`

	// logInConsole: show logs in console at the same time
	LogInConsole bool `mapstructure:"log_in_console"`

	// if true, show color log
	ShowColor bool `mapstructure:"show_color"`

	// if true, only show log, won't print log level、caller func and line
	IsBrief bool `mapstructure:"is_brief"`

	// StackTraceLevel record a stack trace for all messages at or above a given level.
	// Empty string or invalid level will not open stack trace.
	StackTraceLevel string `mapstructure:"stack_trace_level"`
}

func checkLogConfig(logConf *LogConfig) {
	if len(logConf.LogLevel) == 0 {
		logConf.LogLevel = DEFAULT_LOG_LEVEL
	}
	if len(logConf.LogPath) == 0 {
		logConf.LogPath = DEFAULT_LOG_PATH
	}
	if logConf.MaxAge == 0 {
		logConf.MaxAge = DEFAULT_LOG_MAXAGE
	}
	if logConf.MaxSize == 0 {
		logConf.MaxSize = DEFAULT_LOG_MAXSIZE
	}
	if logConf.MaxBackups == 0 {
		logConf.MaxBackups = DEFAULT_LOG_MAXBACKUPS
	}
}

func getZapLevel(lvl string) (*zapcore.Level, error) {
	var zapLevel zapcore.Level
	switch strings.ToUpper(lvl) {
	case ERROR:
		zapLevel = zap.ErrorLevel
	case WARN:
		zapLevel = zap.WarnLevel
	case INFO:
		zapLevel = zap.InfoLevel
	case DEBUG:
		zapLevel = zap.DebugLevel
	default:
		return nil, errors.New("invalid log level")
	}
	return &zapLevel, nil
}

func getLogWriter(fileName string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func InitLogger(logConfig *LogConfig) (*zap.SugaredLogger, *zap.Logger) {

	checkLogConfig(logConfig)
	writeSyncer := getLogWriter(logConfig.LogPath, logConfig.MaxSize, logConfig.MaxBackups, logConfig.MaxAge)

	level, err := getZapLevel(logConfig.LogLevel)
	if err != nil {
		level, _ = getZapLevel(DEFAULT_LOG_LEVEL)
	}

	logger := newLogger(logConfig, level, writeSyncer)
	sugaredLogger := logger.Sugar()

	return sugaredLogger, logger
}

func newLogger(logConfig *LogConfig, level *zapcore.Level, writeSyncer zapcore.WriteSyncer) *zap.Logger {

	var encoderConfig zapcore.EncoderConfig
	if logConfig.IsBrief {
		encoderConfig = zapcore.EncoderConfig{
			TimeKey:    "time",
			MessageKey: "msg",
			EncodeTime: CustomTimeEncoder,
			LineEnding: zapcore.DefaultLineEnding,
		}
	} else {
		encoderConfig = zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "line",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    CustomLevelEncoder,
			EncodeTime:     CustomTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}
	}

	var encoder zapcore.Encoder
	if logConfig.JsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		level,
	)

	logger := zap.New(core)

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)

	if logConfig.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	if lvl, err := getZapLevel(logConfig.StackTraceLevel); err == nil {
		logger = logger.WithOptions(zap.AddStacktrace(lvl))
	}
	logger = logger.WithOptions(zap.AddCallerSkip(1))
	return logger
}

// CustomLevelEncoder 自定义日志级别的输出格式
// @param level
// @param enc
func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// CustomTimeEncoder 自定义时间转字符串的编码方法
// @param t
// @param enc
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
