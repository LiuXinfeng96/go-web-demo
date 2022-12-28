package configs

import (
	"errors"
	"flag"
	"fmt"
	"go-web-demo/src/db"
	"go-web-demo/src/loggers"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string             `mapstructure:"server_port"`
	LogConfig  *loggers.LogConfig `mapstructure:"log_config"`
	DBConfig   *db.DBConfig       `mapstructure:"db_config"`
}

const DEFAULT_SERVER_PORT = "8096"

// GetConfigEnv --Specify the path and name of the configuration file (Env)
func GetConfigEnv() string {
	var env string
	n := len(os.Args)
	for i := 1; i < n-1; i++ {
		if os.Args[i] == "-e" || os.Args[i] == "--env" {
			env = os.Args[i+1]
			break
		}
	}
	fmt.Println("[env]:", env)
	if env == "" {
		fmt.Println("env is empty, set default")
		env = ""
	}
	return env
}

// GetFlagPath --Specify the path and name of the configuration file (flag)
func GetFlagPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "../conf/system_config.yaml", "please input the system config file path")
	flag.Parse()
	return configPath
}

// InitConfig --Set config path and file name
func InitConfig(configPath string) (*Config, error) {
	var err error
	if len(configPath) == 0 {
		configPath = GetFlagPath()
	}

	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	if conf.LogConfig == nil {
		conf.LogConfig = new(loggers.LogConfig)
	}

	if conf.DBConfig == nil {
		return nil, errors.New("not found the db config")
	}

	if len(conf.ServerPort) == 0 {
		conf.ServerPort = DEFAULT_SERVER_PORT
	}

	return &conf, nil
}
