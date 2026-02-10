package router

import (
	"github.com/Nakohartum/practicum-shortener/internal/handler"
	"github.com/go-chi/chi"
)

type ShortenerRouter struct{
	handler *handler.ShortenerHandler
	r chi.Router
}

func NewShortenerRouter(handler *handler.ShortenerHandler, shortenBase string) chi.Router{
	router := chi.NewRouter()
	router.HandleFunc("/", handler.HandleShortenerSet)
    router.HandleFunc(shortenBase+"/{shortenedURL}", handler.HandleShortenerGet)
	return router
}