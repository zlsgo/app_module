package model

import (
	"github.com/sohaha/zlsgo/zlog"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type defaultLogger struct {
	log *zlog.Logger
}

var modelLogger Logger = &defaultLogger{log: zlog.New()}

func SetLogger(l Logger) {
	if l != nil {
		modelLogger = l
	}
}

func GetLogger() Logger {
	return modelLogger
}

func (l *defaultLogger) Debug(v ...interface{}) {
	l.log.Debug(v...)
}

func (l *defaultLogger) Info(v ...interface{}) {
	l.log.Info(v...)
}

func (l *defaultLogger) Warn(v ...interface{}) {
	l.log.Warn(v...)
}

func (l *defaultLogger) Error(v ...interface{}) {
	l.log.Error(v...)
}

func (l *defaultLogger) Debugf(format string, v ...interface{}) {
	l.log.Debugf(format, v...)
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	l.log.Infof(format, v...)
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.log.Warnf(format, v...)
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.log.Errorf(format, v...)
}
