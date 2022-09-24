package main

import (
	"shortener/db"
	"shortener/logger"
	"shortener/routes"
	"shortener/utils"
)

// @title Gin Swagger UrlShortener Api
// @version 1.0
// @description Укоротитель ссылок.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email nekruzrakhimov@icloud.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://shortenertj.herokuapp.com
// @BasePath /
// @schemes http
func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()
	logger.Init()
	db.StartDbConnection()
	routes.RunAllRoutes()
}
