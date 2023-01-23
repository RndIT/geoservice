package logger

import (
	"sync"
)

type Logger interface {
	Debug(v ...any)
}
type logger struct {
	Log *SimpleLogger
}

var myLog *logger
var once sync.Once

// enum
type LogType int

const (
	EPrintLog LogType = iota + 1
	EStdLog
)

func GetLog(logType LogType) Logger {
	once.Do(func() {
		myLog = &logger{}
		myLog.Log = BuildSimpleLogger()

	})
	return *myLog.Log
}
