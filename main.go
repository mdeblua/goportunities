package main

import (
	"github.com/mdeblua/goportunities/config"
	"github.com/mdeblua/goportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error loading config: %s", err)
		return
	}

	logger.Info("Config initialized successfully")

	router.Initialize()

	logger.Info("Api running with real data")
}
