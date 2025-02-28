package config

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func Init() {
	// init config
	if err := LoadCfg(); err != nil {
		panic(err)
	}

	// log rotation
	logger := lumberjack.Logger{
		Filename:   viper.GetString("log.location"),
		MaxSize:    viper.GetInt("log.maxSize"),
		MaxBackups: viper.GetInt("log.maxBackups"),
		MaxAge:     viper.GetInt("log.maxAge"),
	}

	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{
		FullTimestamp: true,
	}}
	formatter.Line = true
	log.SetFormatter(&formatter)
	level := viper.GetString("log.mode")
	if level == "release" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	log.SetOutput(&logger)

	// mysql
	InitDB()
}

func LoadCfg() error {
	if os.Getenv("ATTO_MODE") == "deploy" {
		viper.SetConfigName("config.deploy")
	} else {
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("..")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func GetSecret() string {
	return viper.GetString("encrypt.jwt_secret")
}

func InitForTest() {
	// init config
	if err := LoadCfg(); err != nil {
		panic(err)
	}

	// mysql
	InitDB()
}
