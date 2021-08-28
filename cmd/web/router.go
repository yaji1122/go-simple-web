package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yaji1122/go-simple-web/pkg/config"
	"github.com/yaji1122/go-simple-web/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/", handler.Repo.Home)
	router.Get("/about", handler.Repo.About)
	return router
}
