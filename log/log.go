package log

import (
	"DIMSProxy/config"

	"github.com/doraemonkeys/mylog"
	"github.com/sirupsen/logrus"
)

// var Logger *logrus.Logger
// var NoConsole bool
// var LoggerLevel string
var Logger *logrus.Logger

func NewLogger() {
	cnf := mylog.LogConfig{}
	cnf.LogDir = "./log"                        //设置日志文件的存放路径
	cnf.NoConsole = config.Conf.Local.NoConsole //是否将日志打印在终端
	cnf.MaxKeepDays = 30                        //设置日志文件的最大保存时间
	cnf.LogLevel = mylog.TraceLevel             //设置日志级别为info
	cnf.ShowShortFileInConsole = true
	cnf.DisableWriterBuffer = true
	cnf.DateSplit = true
	cnf.DisableLevelTruncation = true
	cnf.PadLevelText = true
	logger, err := mylog.NewLogger(cnf)
	if err != nil {
		panic(err)
	}
	Logger = logger
}
