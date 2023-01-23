package domain

import "fmt"
import "geoservice/internal/logger"
func Log(l logger) {
	fmt.Println("Domain run")
	l.D
}
