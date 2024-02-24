package logger

import (
	"log"
	"os"
)

type logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *logger) Warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

var flags = log.LstdFlags | log.Lshortfile
var infoLogger = log.New(os.Stdout, "INFO: ", flags)
var warnLogger = log.New(os.Stdout, "WARN: ", flags)
var errorLogger = log.New(os.Stdout, "ERROR: ", flags)

var Logger *logger = &logger{infoLogger, warnLogger, errorLogger}