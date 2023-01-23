package cmd

import (
	"geoservice/internal/logger"
)

var (
	log logger.Logger
)

func init() {
	// init logs
	logger.NewLogger(logger.EStdLog)
	logger.NewLogger(logger.EPrintLog)

	logger.Debug("init: Use EPrintLog DEBUG")
	//logStdPrint = logger.GetLog(logger.EStdLog)
	//logStdPrint.Debug("init: Use EStdLog DEBUG")
}

func Run() {

	//logPrint.Debug("run: Use EPrintLog DEBUG")
	//logStdPrint.Debug("run: Use EStdLog DEBUG")

	//domain.Log(logPrint)
}
