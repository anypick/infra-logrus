/**
TODO
自定义一些hook
 1. 将日志输出给kafka
 2. 将错误日志发送邮件给用户
 ...
 */
package helper

import (
	"github.com/sirupsen/logrus"
)

// error级别日志发送给对应的邮箱，这个邮箱应该来自数据库
type MailHook struct {
	MailAddr []string
	SendMail func(message string) error
}

func (m MailHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (m MailHook) Fire(entry *logrus.Entry) error {
	return m.SendMail(entry.Message)
}
type KafkaHook struct {}

func (m KafkaHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (m KafkaHook) Fire(entry *logrus.Entry) error {
	return nil
}

