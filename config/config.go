package config

import (
	"strconv"

	"github.com/boqier/gin-scaffold/utils/logs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type ReturnDate struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewReturnDate() ReturnDate {
	return ReturnDate{
		Status:  200,
		Data:    make(map[string]interface{}),
		Message: "success",
	}
}

var (
	Port            string
	JWTSecret       string
	JWT_EXPIRE_TIME int64  //TOKEN 过期时间（分钟）
	Username        string //默认admin
	Password        string //默认123456
)

func initLogConfig(logLevel string) {
	switch logLevel {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
	//设置文件名和行号
	logrus.SetReportCaller(true)
	//设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})

}
func init() {
	logs.Debug(map[string]interface{}{"module": "config"}, "开始加载程序配置")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	initLogConfig(logLevel)
	//处理端口号：获取程序启动端口号的配置
	viper.SetDefault("PORT", "8080")
	Port = viper.GetString("PORT")
	logs.Info(map[string]interface{}{"module": "config"}, "程序启动端口号："+Port)
	//获取Jwt加密的serect
	viper.SetDefault("JWT_SECRET", "liuneng")
	JWTSecret = viper.GetString("JWT_SECRET")
	logs.Info(map[string]interface{}{"module": "config"}, "JWT_SECRET："+JWTSecret)
	//获取JWT过期时间
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	JWT_EXPIRE_TIME = viper.GetInt64("JWT_EXPIRE_TIME")
	logs.Info(map[string]interface{}{"module": "config"}, "JWT_EXPIRE_TIME："+strconv.FormatInt(JWT_EXPIRE_TIME, 10))
	//获取用户名与密码
	viper.SetDefault("USERNAME", "21232F297A57A5A743894A0E4A801FC3")
	Username = viper.GetString("USERNAME")
	logs.Info(map[string]interface{}{"module": "config"}, "USERNAME："+Username)
	viper.SetDefault("PASSWORD", "E10ADC3949BA59ABBE56E057F20F883E")
	Password = viper.GetString("PASSWORD")
	logs.Info(map[string]interface{}{"module": "config"}, "PASSWORD："+Password)
}
