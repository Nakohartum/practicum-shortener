package service

import "github.com/Nakohartum/practicum-shortener/internal/model"

type DataHandler interface {
	GetData(key string) model.ShortenUrl
	SetData(key, value string)
}

type ShortenUrlService struct {
	repo DataHandler
}

func NewShortenUrlService(dataHandler DataHandler) *ShortenUrlService {
	return &ShortenUrlService{
		repo: dataHandler,
	}
}

func (sr *ShortenUrlService) GetData(key string) model.ShortenUrl {
	return sr.repo.GetData(key)
}

func (sr *ShortenUrlService) SetData(key, value string) {
	sr.repo.SetData(key, value)
}