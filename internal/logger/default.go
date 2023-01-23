package logger

import (
	"log"
	"sync"
)

type SimpleLogger struct {
	log *log.Logger
}

var instanceSimpleLogger *SimpleLogger = nil

func BuildSimpleLogger() *SimpleLogger {
	var once sync.Once
	once.Do(func() {
		instanceSimpleLogger = &SimpleLogger{}
		instanceSimpleLogger.log = log.Default()
	})

	return instanceSimpleLogger
}

func (l SimpleLogger) Debug(v ...any) {
	instanceSimpleLogger.log.Print(v...)
}
