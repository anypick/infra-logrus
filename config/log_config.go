package config

import (
	"fmt"
	"github.com/anypick/infra/base/props/container"
)

const (
	YamlPrefix = "logrus"
)

func init() {
	container.RegisterYamContainer(&LogConfig{Prefix: YamlPrefix})
}

// 日志配置
type LogConfig struct {
	Prefix       string
	Level        string `yaml:"level"`
	LogFileName  string `yaml:"logFileName"`
	FilePath     string `yaml:"filePath"`
	MaxAge       int    `yaml:"maxAge"`
	RotationTime int    `yaml:"rotationTime"`
}

func (l *LogConfig) ConfigAdd(config map[interface{}]interface{}) {
	l.Level = fmt.Sprintf("%v", config["level"])
	l.LogFileName = fmt.Sprintf("%v", config["logFileName"])
	l.FilePath = fmt.Sprintf("%v", config["filePath"])
	l.MaxAge = config["maxAge"].(int)
	l.RotationTime = config["rotationTime"].(int)
}
