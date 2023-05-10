package tools

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	// 设置日志格式为json格式
	Log.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true, //键值对加引号
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	Log.SetOutput(os.Stdout)
	// 设置日志级别为warn以上
	Log.SetLevel(logrus.InfoLevel)

}
