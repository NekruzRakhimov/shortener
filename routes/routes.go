package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"shortener/pkg/controller"
	"shortener/utils"
)

func RunAllRoutes() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = utils.AppSettings.AppParams.PortRun
	}
	_ = r.Run(fmt.Sprintf(":%s", port))
}

func initAllRoutes(r *gin.Engine) {
	r.GET("/", PingPong)

	r.POST("/shorten", controller.ShortenUrl)
	r.GET("/expand", controller.ExpandUrl)
}

// PingPong Проверка
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "server is up and running"})
}
