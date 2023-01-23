package logger

import (
	"fmt"
	"sync"
)

type PrintLogger struct {
	//log *log.Logger
}

var instancePrintLogger *PrintLogger = nil

func BuildPrintLogger() *PrintLogger {
	var once sync.Once
	once.Do(func() {
		instancePrintLogger = &PrintLogger{}
		//instanceLogger.log = log.Default()
	})

	return instancePrintLogger
}

func (l PrintLogger) Debug(v ...any) {
	fmt.Println(v...)
	//instanceSimpleLogger.log.Print(v...)
}
