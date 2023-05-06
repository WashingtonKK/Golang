package main

import (
	"time"

	"github.com/Golang/logging"
)

func main() {
	logger := logging.New(time.RFC3339, true)
	logger.Log("Error", "This is an error")
	logger.Log("Info", "An informative log")
	logger.Log("Warning", "Harsh warnings")
}
