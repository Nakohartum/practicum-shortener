package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Nakohartum/practicum-shortener/internal/config"
	"github.com/Nakohartum/practicum-shortener/internal/service"
	"github.com/go-chi/chi"
)

func TestHandleShortenerSet(t *testing.T) {
	tests := []struct {
		name string
		actualLink string
		method string
		statusCode int
	}{
		{
			name: "positive #1",
			actualLink: "test.com",
			method: http.MethodPost,
			statusCode: http.StatusCreated,
		},
		{
			name: "not correct link",
			actualLink: "",
			method: http.MethodPost,
			statusCode: http.StatusBadRequest,
		},
		{
			name: "not correct method",
			actualLink: "",
			method: http.MethodGet,
			statusCode: http.StatusBadRequest,
		},
	}

	repo := config.NewInMemoryRepo()
	dt := service.NewShortenURLService(repo)
	sh := NewShortenerHandler(dt)

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.actualLink))

			w := httptest.NewRecorder()

			sh.HandleShortenerSet(w, request)

			res := w.Result()
			defer res.Body.Close()

			if tt.statusCode != res.StatusCode {
				t.Errorf("not correct status code. expected: %d, actual: %d", tt.statusCode, res.StatusCode)
			}
		})
	}
}

func TestHandleShortenerGet(t *testing.T) {
	tests := []struct{
		name string
		link string
		shortenedLink string
		status int
		method string
	}{
		{
			name: "positive #1",
			link: "test.com",
			shortenedLink: "z284XC",
			status: http.StatusTemporaryRedirect,
			method: http.MethodGet,
		},
		{
			name: "not correct method",
			link: "test.com",
			shortenedLink: "zzdffas",
			status: http.StatusBadRequest,
			method: http.MethodPost,
		},
	}

	repo := config.NewInMemoryRepo()
	dt := service.NewShortenURLService(repo)
	sh := NewShortenerHandler(dt)

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T) {
			repo.SetData(tt.shortenedLink, tt.link)

			r := chi.NewRouter()
			r.Handle("/{shortenedURL}", http.HandlerFunc(sh.HandleShortenerGet))

			req := httptest.NewRequest(tt.method, "/"+tt.shortenedLink, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			if tt.status != res.StatusCode{
				t.Errorf("not correct status. expected %d got %d", tt.status, res.StatusCode)
			}
		})
	}
}