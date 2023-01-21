package logger

import (
	"sync"
)

type Logger interface {
	Debug(v ...any)
}
type logger struct {
	Log *DefPrintLog
}

var myLog *logger
var onceLocal sync.Once

// enum
type LogType int

const (
	PrintLog LogType = iota + 1
	StdLog
)

func GetLog(logType LogType) Logger {
	onceLocal.Do(func() {
		myLog = &logger{}
		myLog.Log = BuildDefault()
	})
	return *myLog.Log
}
