package main

import (
	"fmt"
	"github.com/d-kolpakov/logger"
)

const ServiceName = "device_api_shops"

var AppVersion string

func main() {
	fmt.Println("Starting service " + ServiceName)

	//Logger initialization
	logDriver := &logger.STDOUTDriver{}
	loggerConfig := logger.LoggerConfig{
		ServiceName: ServiceName,
		Level:       configo.EnvInt("logging-level", 2),
		Buffer:      configo.EnvInt("app-logger-buffer-size", 10000),
		Output:      []logger.LogDriver{logDriver},
	}
	l, err := logger.GetLogger(loggerConfig)
	if err != nil {
		panic(err)
	}
}
