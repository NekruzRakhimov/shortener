package models

type Url struct {
	ID       int    `json:"id"`
	ShortUrl string `json:"short_url"`
	FullUrl  string `json:"full_url"`
}
