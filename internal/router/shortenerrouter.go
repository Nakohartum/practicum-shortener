package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Nakohartum/practicum-shortener/internal/handler"
)

type ShortenerRouter struct{
	handler *handler.ShortenerHandler
}

func NewShortenerRouter(handler *handler.ShortenerHandler) *ShortenerRouter{
	return &ShortenerRouter{
		handler: handler,
	}
}

func (sr *ShortenerRouter) HandleShortenerRequest(rw http.ResponseWriter, req *http.Request){
	urlParts := strings.Split(strings.Trim(req.URL.Path, "/"), "/")

	fmt.Println(strings.Trim(req.URL.Path, "/"))
	switch{
	case urlParts[0] == "":
		sr.handler.HandleShortenerSet(rw, req)
	case urlParts[0] != "":
		sr.handler.HandleShortenerGet(rw, req)
	}
}