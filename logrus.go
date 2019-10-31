package baselog

import (
	"github.com/anypick/infra"
	"github.com/anypick/infra-logrus/config"
	"github.com/anypick/infra-logrus/helper"
	"github.com/anypick/infra/utils/common"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"path"
	"strings"
	"time"
)

type LogrusStarter struct {
	infra.BaseStarter
}

func (l *LogrusStarter) Init(ctx infra.StarterContext) {
	initLogrus(*ctx.Yaml()[config.YamlPrefix].(*config.LogConfig))
}

// 初始化Logrus
func initLogrus(config config.LogConfig) {
	var (
		level          logrus.Level                   // 日志级别
		logPath        = config.FilePath              // 日志存放路径
		fileName       = config.LogFileName           // 日志名称
		logPathAndName = path.Join(logPath, fileName) // 日志路径+名称
		out            *rotatelogs.RotateLogs         // 日志输出
		err            error
	)

	// 使用
	formatter := &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: common.TimeFormat,
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   false,
	}
	logrus.SetFormatter(formatter)
	for _, hook := range helper.GetHook() {
		logrus.AddHook(hook)
	}
	switch config.Level {
	case "trace":
		level = logrus.TraceLevel
		break
	case "debug":
		level = logrus.DebugLevel
		break
	case "info":
		level = logrus.InfoLevel
		break
	case "warn":
		level = logrus.WarnLevel
	case "error":
		level = logrus.ErrorLevel
		break
	default:
		level = logrus.DebugLevel
		break
	}
	logrus.SetLevel(level)

	if level == logrus.TraceLevel {
		logrus.SetOutput(os.Stdout)
		return
	}
	// 创建日志目录，如果存在则忽略
	os.MkdirAll(logPath, os.ModePerm)
	// 定义日志切割
	if out, err = rotatelogs.New(
		strings.TrimSuffix(logPathAndName, ".log")+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(logPathAndName), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(config.MaxAge)),
		rotatelogs.WithRotationTime(time.Duration(config.RotationTime)),
	); err != nil {
		logrus.Errorf("new rotatelogs error %v", err)
		return
	}
	logrus.SetOutput(out)
}
