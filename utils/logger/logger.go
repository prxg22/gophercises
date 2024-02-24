package logger

import (
	"log"
	"os"
)

const flags = log.LstdFlags | log.Lshortfile

var infoLogger, warnLogger, errorLogger = log.New(os.Stdout, "INFO: ", flags),
	log.New(os.Stdout, "WARN: ", flags),
	log.New(os.Stdout, "ERROR: ", flags)

func Info(str string, v ...interface{}) {
	infoLogger.Printf(str, v...)
}

func Warn(str string, v ...interface{}) {
	warnLogger.Printf(str, v...)
}

func Error(err error) {
	errorLogger.Println(err)
}
