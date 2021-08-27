package main

import (
	"github.com/bmizerany/pat"
	"github.com/yaji1122/go-simple-web/pkg/config"
	"github.com/yaji1122/go-simple-web/pkg/handler"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))
	return mux
}
