package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"shortener/models"
	"shortener/pkg/repository"
	"shortener/utils"
)

var maxLength = 10

func ShortenUrl(url *models.Url) error {
	err := repository.GetUrlByFullUrl(url)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	for {
		url.ShortUrl = fmt.Sprintf("%s/%s", utils.AppSettings.AppParams.ServerURL, utils.RandomString(maxLength))
		err = repository.GetUrlByShortUrl(url)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := repository.SaveShortenUrl(url); err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func GetUrlByShortUrl(url *models.Url) error {
	return repository.GetUrlByShortUrl(url)
}
