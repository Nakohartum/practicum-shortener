package main

import (
	"net/http"
	"github.com/Nakohartum/practicum-shortener/internal/config"
	"github.com/Nakohartum/practicum-shortener/internal/handler"
	"github.com/Nakohartum/practicum-shortener/internal/router"
	"github.com/Nakohartum/practicum-shortener/internal/service"
)

func main() {
	repo := config.NewInMemoryRepo()
	dataHandlerService := service.NewShortenURLService(repo)
	handler := handler.NewShortenerHandler(dataHandlerService)
	router := router.NewShortenerRouter(handler)

	http.ListenAndServe(":8080", router)
}
