package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Nakohartum/practicum-shortener/internal/handler"
	"github.com/go-chi/chi"
)

type ShortenerRouter struct{
	handler *handler.ShortenerHandler
	r chi.Router
}

func NewShortenerRouter(handler *handler.ShortenerHandler, r chi.Router) *ShortenerRouter{
	return &ShortenerRouter{
		handler: handler,
		r: r,
	}
}

func (sr *ShortenerRouter) HandleShortenerRequest(rw http.ResponseWriter, req *http.Request){
	sr.r.Route("/", func(r chi.Router) {
		r.Post("/", sr.handler.HandleShortenerSet)
		r.Get("/{shortenedURL}", sr.handler.HandleShortenerGet)
	})
}