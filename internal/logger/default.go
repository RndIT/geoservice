package logger

import (
	"log"
	"sync"
)

type SimpleLogger struct {
	log log.Logger
}

var instanceSimpleLogger *SimpleLogger = nil
var onceSL sync.Once
func BuildSimpleLogger() Logger {
	
	onceSL.Do(func() {
		instanceSimpleLogger = &SimpleLogger{}
		instanceSimpleLogger.log = *log.Default()
		instanceSimpleLogger.Debug("Init Simple logger done")
	})

	return instanceSimpleLogger
}

func (l SimpleLogger) Debug(v ...any) {
	instanceSimpleLogger.log.Print(v...)
}
