package cmd

import (
	"geoservice/internal/domain"
	"geoservice/internal/logger"
)

var (
	logPrint    logger.Logger
	logStdPrint logger.Logger
)

func init() {
	// init logs
	logPrint = logger.GetLog(logger.EPrintLog)
	logPrint.Debug("init: Use EPrintLog DEBUG")
	logStdPrint = logger.GetLog(logger.EStdLog)
	logStdPrint.Debug("init: Use EStdLog DEBUG")
}

func Run() {

	logPrint.Debug("run: Use EPrintLog DEBUG")
	logStdPrint.Debug("run: Use EStdLog DEBUG")

	domain.Print()
}
