package common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger logrus的再封装
type Logger struct {
	log *logrus.Logger
}

//logLevel 字典，用于确定log level
var logLevel = map[string]logrus.Level{
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

// NewLogger 构造函数，只有高于等于level的信息会记录日志
func NewLogger(filePath string, level string) *Logger {
	var log = logrus.New()
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Infoln("Failed to log to file, use default stderr")
		log.Out = os.Stdout
	}
	log.SetLevel(logLevel[strings.ToLower(level)])
	return &Logger{log: log}
}

// LogInfo InfoLevel Log
func (l *Logger) LogInfo(event string, state string, msg string) {
	l.log.WithFields(
		logrus.Fields{
			"event": event,
			"state": state,
		},
	).Info(msg)
	fmt.Printf("%s [%s] EVENT:%s; STATE:%s; MSG:%v\n", time.Now().Format("2006-01-02 15:04:05"), "INFO", event, state, msg)
}

// LogWarn WarnLevel Log
func (l *Logger) LogWarn(event string, state string, msg string) {
	l.log.WithFields(
		logrus.Fields{
			"event": event,
			"state": state,
		},
	).Warn(msg)
	fmt.Printf("%s [%s] EVENT:%s; STATE:%s; MSG:%v\n", time.Now().Format("2006-01-02 15:04:05"), "WARN", event, state, msg)
}

// LogError ErrorLevel Log
func (l *Logger) LogError(event string, state string, action string, data interface{}, err error) {
	l.log.WithFields(
		logrus.Fields{
			"event":  event,
			"state":  state,
			"action": action,
			"data":   data,
		},
	).Error(err)
	fmt.Printf("%s [%s] EVENT:%s; STATE:%s; ACTION:%s; DATA:%v; ERR:%v\n", time.Now().Format("2006-01-02 15:04:05"), "ERROR", event, state, action, data, err.Error())
}
