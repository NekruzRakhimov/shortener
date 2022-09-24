package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"shortener/logger"
	"shortener/models"
	"shortener/pkg/service"
	"shortener/utils"
)

//ShortenUrl shortener godoc
// @Summary Shortening url
// @Description Роут для сокращения ссылок. Нужно заполнять только поле full_url. И в ответе смотреть на short_url
// @Accept  json
// @Produce  json
// @Tags url
// @Param  url  body models.Url true "full_url"
// @Success 200 {object} models.Url
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /shorten [post]
func ShortenUrl(c *gin.Context) {
	var myUrl models.Url
	if err := c.BindJSON(&myUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время парсинга тело"})
		logger.Error.Printf("[%s] Error is: %s", utils.FuncName(), err.Error())
		return
	}

	if err := service.ShortenUrl(&myUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "что-то пошло не так"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": myUrl.ShortUrl})
}

//GetUrlByShortUrl shortener godoc
// @Summary Shortening url
// @Description Роут для получения исходной ссылки из сокращенной. В ответе нужно смотреть на поле full_url
// @Accept  json
// @Produce  json
// @Tags url
// @Param  url  query string true "short_url"
// @Success 200 {object} models.Url
// @Failure 400,404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /expand [get]
func GetUrlByShortUrl(c *gin.Context) {
	var myUrl models.Url
	unescape, err := url.QueryUnescape(c.Query("url"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "ошибка во время чтения query параметра"})
		logger.Error.Printf("[%s] Error is: %s", utils.FuncName(), err.Error())
		return
	}
	logger.Debug.Println(unescape)
	myUrl.ShortUrl = unescape

	if err := service.GetUrlByShortUrl(&myUrl); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": "ваша ссылка не найдена"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "что-то пошло не так"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"full_url": myUrl.FullUrl})
}
