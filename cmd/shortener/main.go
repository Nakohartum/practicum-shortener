package main

import (
	"net/http"

	"github.com/Nakohartum/practicum-shortener/internal/config"
	"github.com/Nakohartum/practicum-shortener/internal/handler"
	"github.com/Nakohartum/practicum-shortener/internal/router"
	"github.com/Nakohartum/practicum-shortener/internal/service"
	"github.com/go-chi/chi"
)

func main() {
	mux := http.NewServeMux()
	repo := config.NewInMemoryRepo()
	dataHandlerService := service.NewShortenURLService(repo)
	handler := handler.NewShortenerHandler(dataHandlerService)
	r := chi.NewRouter()
	router := router.NewShortenerRouter(handler, r)
	
	mux.HandleFunc("/", router.HandleShortenerRequest)
	

	http.ListenAndServe(":8080", mux)
}
