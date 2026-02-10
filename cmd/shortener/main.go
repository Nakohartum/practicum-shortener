package main

import (
	"net/http"
	"github.com/Nakohartum/practicum-shortener/internal/config"
	"github.com/Nakohartum/practicum-shortener/internal/handler"
	"github.com/Nakohartum/practicum-shortener/internal/router"
	"github.com/Nakohartum/practicum-shortener/internal/service"
)

func main() {
	readFlags()
	repo := config.NewInMemoryRepo()
	dataHandlerService := service.NewShortenURLService(repo)
	handler := handler.NewShortenerHandler(dataHandlerService, serverFlags.shortenBaseAddress)
	router := router.NewShortenerRouter(handler, serverFlags.shortenBaseAddress)

	http.ListenAndServe(serverFlags.address, router)
}
