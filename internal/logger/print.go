package logger

import (
	"fmt"
	"sync"
)

type PrintLogger struct {
}

var instancePrintLogger *PrintLogger = nil
var oncePL sync.Once
func BuildPrintLogger() Logger {
	oncePL.Do(func() {
		instancePrintLogger = &PrintLogger{}
		instancePrintLogger.Debug("Instance Print logger done")
	})
	return instancePrintLogger
}

func (l PrintLogger) Debug(v ...any) {
	fmt.Println(v...)
	//instanceSimpleLogger.log.Print(v...)
}
