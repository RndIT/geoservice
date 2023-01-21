package cmd

import (
	"fmt"
	"geoservice/internal/domain"
	"geoservice/internal/logger"
)

var (
	log logger.Logger
)

func init() {
	// init logs
	log = logger.GetLog(logger.PrintLog)
	log.Debug("DEBUG")
}

func Run() {

	fmt.Println("Execute")
	log.Debug("sdfsdfs")

	domain.Print()
}
