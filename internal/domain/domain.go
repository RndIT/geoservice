package domain

import (
	"fmt"
	"geoservice/internal/logger"
)

func Log(l logger.Logger) {
	fmt.Println("Domain run")
	l.Debug("domain logger message")
}
