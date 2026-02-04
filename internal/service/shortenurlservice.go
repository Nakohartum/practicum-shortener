package service

import "github.com/Nakohartum/practicum-shortener/internal/model"

type DataHandler interface {
	GetData(key string) model.ShortenURL
	SetData(key, value string)
}

type ShortenURLService struct {
	repo DataHandler
}

func NewShortenURLService(dataHandler DataHandler) *ShortenURLService {
	return &ShortenURLService{
		repo: dataHandler,
	}
}

func (sr *ShortenURLService) GetData(key string) model.ShortenURL {
	return sr.repo.GetData(key)
}

func (sr *ShortenURLService) SetData(key, value string) {
	sr.repo.SetData(key, value)
}
