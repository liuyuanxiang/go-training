package logger

import baseLogger "git.mysre.cn/yunlian-golang/go-hulk/logger"

var l baseLogger.LoggerInterface

func InitLogger(bl baseLogger.LoggerInterface) {
	l = bl
}

func Debug(v ...interface{}) { l.Debug(v...) }
func Info(v ...interface{})  { l.Info(v...) }
func Warn(v ...interface{})  { l.Warn(v...) }
func Error(v ...interface{}) { l.Error(v...) }
