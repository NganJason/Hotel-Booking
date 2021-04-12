package main

import (
	"net/http"

	"github.com/NganJason/hotel-booking/pkg/config"
	"github.com/NganJason/hotel-booking/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middlewares
	mux.Use(middleware.Recoverer)	
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.HandleHome)
	mux.Get("/about", handlers.Repo.HandleAbout)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}