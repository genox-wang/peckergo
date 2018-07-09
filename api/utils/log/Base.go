package log

import (
	"log"
)

const (
	// LogLevelDebug 调试等级
	LogLevelDebug = 1
	// LogLevelInfo 信息等级
	LogLevelInfo = 2
	// LogLevelWarn 警告等级
	LogLevelWarn = 3
	// LogLevelError 错误等级
	LogLevelError = 4
	// LogLevelFatal Fatal等级
	LogLevelFatal = 5
	// LogLevelPanic Panic等级
	LogLevelPanic = 6
)

var (
	l Base
)

// Base 日志模型接口
type Base interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	SetLevel(level int)
}

// InitLog 创建并绑定 Log实例， 目前支持 &Logurs{}
func InitLog(logBase Base, level int) {
	if logBase == nil {
		log.Panic("logBase is nil")
	}
	l = logBase
	logBase.SetLevel(level)
}

// PanicWhenLogNil 字面意
func PanicWhenLogNil() {
	if l == nil {
		log.Panic("logBase is nil")
	}
}

// Debug 字面意
func Debug(args ...interface{}) {
	PanicWhenLogNil()
	l.Debug(args...)
}

// Info 字面意
func Info(args ...interface{}) {
	PanicWhenLogNil()
	l.Info(args...)
}

// Warn 字面意
func Warn(args ...interface{}) {
	PanicWhenLogNil()
	l.Warn(args...)
}

// Error 字面意
func Error(args ...interface{}) {
	PanicWhenLogNil()
	l.Error(args...)
}

// Fatal 字面意
func Fatal(args ...interface{}) {
	PanicWhenLogNil()
	l.Fatal(args...)
}

// Panic 字面意
func Panic(args ...interface{}) {
	PanicWhenLogNil()
	l.Panic(args...)
}

// Debugf 字面意
func Debugf(format string, args ...interface{}) {
	PanicWhenLogNil()
	l.Debugf(format, args...)
}

// Infof 字面意
func Infof(format string, args ...interface{}) {
	PanicWhenLogNil()
	l.Infof(format, args...)
}

// Warnf 字面意
func Warnf(format string, args ...interface{}) {
	PanicWhenLogNil()
	l.Warnf(format, args...)
}

// Errorf 字面意
func Errorf(format string, args ...interface{}) {
	PanicWhenLogNil()
	l.Errorf(format, args...)
}

// Fatalf 字面意
func Fatalf(format string, args ...interface{}) {
	PanicWhenLogNil()
	l.Fatalf(format, args...)
}

// Panicf 字面意
func Panicf(format string, args ...interface{}) {
	PanicWhenLogNil()
	l.Panicf(format, args...)
}
