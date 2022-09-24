package repository

import (
	"shortener/db"
	"shortener/logger"
	"shortener/models"
	"shortener/utils"
)

func SaveShortenUrl(url *models.Url) error {
	if err := db.GetDBConn().Table("urls").Create(url).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s", utils.FuncName(), err.Error())
		return err
	}

	return nil
}

func GetUrlByShortUrl(url *models.Url) error {
	if err := db.GetDBConn().Table("urls").First(&url, "short_url = ?", url.ShortUrl).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s", utils.FuncName(), err.Error())
		return err
	}
	return nil
}

func GetUrlByFullUrl(url *models.Url) error {
	if err := db.GetDBConn().Table("urls").First(&url, "full_url = ?", url.FullUrl).Error; err != nil {
		logger.Error.Printf("[%s] Error is: %s", utils.FuncName(), err.Error())
		return err
	}
	return nil
}
