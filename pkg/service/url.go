package service

import (
	"fmt"
	"shortener/models"
	"shortener/pkg/repository"
	"shortener/utils"
)

var maxLength int = 10

func ShortenUrl(url *models.Url) error {
	url.ShortUrl = fmt.Sprintf("%s/%s", utils.AppSettings.AppParams.ServerURL, utils.RandomString(maxLength))
	if err := repository.SaveShortenUrl(url); err != nil {
		return err
	}

	return nil
}

func ExpandUrl(url *models.Url) error {
	return repository.ExpandUrl(url)
}
