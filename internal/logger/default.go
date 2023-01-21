package logger

import (
	"log"
	"sync"
)

type DefPrintLog struct {
	log *log.Logger
}

var defLogInstance *DefPrintLog = nil
var onceDef sync.Once

func BuildDefault() *DefPrintLog {
	onceDef.Do(func() {
		defLogInstance = &DefPrintLog{}
		defLogInstance.log = log.Default()
	})
	return defLogInstance
}

func (l DefPrintLog) Debug(v ...any) {
	defLogInstance.log.Print(v...)
}
