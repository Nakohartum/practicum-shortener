package config

import "github.com/Nakohartum/practicum-shortener/internal/model"

type InMemoryRepo struct {
	urls map[string]model.ShortenUrl
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		urls: make(map[string]model.ShortenUrl),
	}
}

func (im *InMemoryRepo) GetData(key string) model.ShortenUrl {
	if val, exists := im.urls[key]; exists{
		return val
	}
	return model.ShortenUrl("")
}

func (im *InMemoryRepo) SetData(key, value string) {
	im.urls[key] = model.ShortenUrl(value)
}