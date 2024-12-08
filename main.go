
package main

import (
	"Demo/src/config"
	"Demo/src/logger"
	"Demo/src/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Logger Initialized successfully")
	logger.InfoLn("Database Initialize successfully")
	router.Init()
	logger.InfoLn("Router Initialized successfully")
}
