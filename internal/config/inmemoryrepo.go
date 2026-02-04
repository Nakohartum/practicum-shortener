package config

import "github.com/Nakohartum/practicum-shortener/internal/model"

type InMemoryRepo struct {
	urls map[string]model.ShortenURL
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		urls: make(map[string]model.ShortenURL),
	}
}

func (im *InMemoryRepo) GetData(key string) model.ShortenURL {
	if val, exists := im.urls[key]; exists {
		return val
	}
	return model.ShortenURL("")
}

func (im *InMemoryRepo) SetData(key, value string) {
	im.urls[key] = model.ShortenURL(value)
}
