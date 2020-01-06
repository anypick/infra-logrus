package config

import (
	"fmt"
)

const (
	YamlPrefix = "logrus"
)

// 日志配置
type LogConfig struct {
	Prefix       string
	Level        string `yaml:"level"`
	LogFileName  string `yaml:"logFileName"`
	FilePath     string `yaml:"filePath"`
	MaxAge       int    `yaml:"maxAge"`
	RotationTime int    `yaml:"rotationTime"`
	Formatter    string `yaml:"formatter"`
	KvCom        string `yaml:"kv_com"`
	FieldMapCom  string `yaml:"field_map_com"`
}

func (l *LogConfig) ConfigAdd(config map[interface{}]interface{}) {
	l.Level = fmt.Sprintf("%v", config["level"])
	l.LogFileName = fmt.Sprintf("%v", config["logFileName"])
	l.FilePath = fmt.Sprintf("%v", config["filePath"])
	l.MaxAge = config["maxAge"].(int)
	l.RotationTime = config["rotationTime"].(int)
	l.Formatter = config["formatter"].(string)
	l.KvCom = config["kv_com"].(string)
	l.FieldMapCom = config["field_map_com"].(string)
}
