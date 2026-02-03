package main

import (
	"net/http"

	"github.com/Nakohartum/practicum-shortener/internal/config"
	"github.com/Nakohartum/practicum-shortener/internal/handler"
	"github.com/Nakohartum/practicum-shortener/internal/router"
	"github.com/Nakohartum/practicum-shortener/internal/service"
)

func main() {
	mux := http.NewServeMux()
	repo := config.NewInMemoryRepo()
	dataHandlerService := service.NewShortenUrlService(repo)
	handler := handler.NewShortenerHandler(dataHandlerService)
	router := router.NewShortenerRouter(handler)
	mux.HandleFunc("/", router.HandleShortenerRequest)

	http.ListenAndServe(":8080", mux)
}
