package handler

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/Nakohartum/practicum-shortener/internal/service"
	"github.com/go-chi/chi"
)

type ShortenerHandler struct {
	dataHandlerService service.DataHandler
	shortenBase string
}

func NewShortenerHandler(dataHandlerService service.DataHandler, shortenBase string) *ShortenerHandler {
	return &ShortenerHandler{
		dataHandlerService: dataHandlerService,
		shortenBase: shortenBase,
	}
}

func (sh *ShortenerHandler) HandleShortenerSet(rw http.ResponseWriter, req *http.Request ){
	if req.Method != http.MethodPost {
		http.Error(rw, "not correct method", http.StatusBadRequest)
		return
	}

	if url, err := io.ReadAll(req.Body); err != nil {
		http.Error(rw, "error happened reading body", http.StatusBadRequest)
		return
	} else {
		if len(url) < 1{
			http.Error(rw, "error happened reading body", http.StatusBadRequest)
			return
		}
		shortentRes, err := shortenURL(6)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		rw.Header().Set("Content-Type", "text/plain")

		res := sh.shortenBase + "/" + shortentRes
		sh.dataHandlerService.SetData(shortentRes, string(url))
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(res))
	}
}

func shortenURL(nBytes int) (string, error) {
	b := make([]byte, nBytes)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (sh *ShortenerHandler) HandleShortenerGet(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet{
		http.Error(rw, "not correct method", http.StatusBadRequest)
		return
	}

	shortentURL := string(chi.URLParam(req, "shortenedURL"))

	if shortentURL == ""{
		http.Error(rw, "not correct path", http.StatusBadRequest)
		return
	}

	res := sh.dataHandlerService.GetData(shortentURL)
	rw.Header().Set("Location", string(res))
	rw.WriteHeader(http.StatusTemporaryRedirect)
}

