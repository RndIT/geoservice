package cmd

import (
	"fmt"
	"geoservice/internal/domain"
)

func Execute() {
	fmt.Println("Execute")
	domain.Print()
}
