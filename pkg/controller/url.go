package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shortener/logger"
	"shortener/models"
	"shortener/pkg/service"
	"shortener/utils"
)

func ShortenUrl(c *gin.Context) {
	var url models.Url
	if err := c.BindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время парсинга тело"})
		logger.Error.Printf("[%s] Error is: %s", utils.FuncName(), err.Error())
		return
	}

	if err := service.ShortenUrl(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "что-то пошло не так"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": url.ShortUrl})
}

func GetUrlByShortUrl(c *gin.Context) {
	var url models.Url
	url.FullUrl = c.Query("url")

	if err := service.GetUrlByShortUrl(&url); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": "ваша ссылка не найдена"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "что-то пошло не так"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"full_url": url.FullUrl})
}
