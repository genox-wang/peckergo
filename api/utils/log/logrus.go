package log

import (
	"github.com/sirupsen/logrus"
)

// Logrus 基于 Logrus 的 LogBase 实现
type Logrus struct {
}

// Debug 字面意
func (*Logrus) Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Info 字面意
func (*Logrus) Info(args ...interface{}) {
	logrus.Info(args...)
}

// Warn 字面意
func (*Logrus) Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Error 字面意
func (*Logrus) Error(args ...interface{}) {
	logrus.Error(args...)
}

// Fatal 字面意
func (*Logrus) Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Panic 字面意
func (*Logrus) Panic(args ...interface{}) {
	logrus.Panic(args...)
}

// Debugf 字面意
func (*Logrus) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Infof 字面意
func (*Logrus) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Warnf 字面意
func (*Logrus) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Errorf 字面意
func (*Logrus) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Fatalf 字面意
func (*Logrus) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Panicf 字面意
func (*Logrus) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// SetLevel 字面意
func (*Logrus) SetLevel(level int) {
	switch level {
	case LogLevelPanic:
		logrus.SetLevel(logrus.PanicLevel)
	case LogLevelFatal:
		logrus.SetLevel(logrus.FatalLevel)
	case LogLevelError:
		logrus.SetLevel(logrus.ErrorLevel)
	case LogLevelWarn:
		logrus.SetLevel(logrus.WarnLevel)
	case LogLevelInfo:
		logrus.SetLevel(logrus.InfoLevel)
	case LogLevelDebug:
		logrus.SetLevel(logrus.DebugLevel)
	}
}
