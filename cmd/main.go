package main

import (
	"shortener/logger"
	"shortener/routes"
	"shortener/utils"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	//db.StartDbConnection()
	//go jobs.StartJobs()
	routes.RunAllRoutes()
}
