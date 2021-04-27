package main

import (
	"net/http"

	"github.com/NganJason/hotel-booking/internal/config"
	"github.com/NganJason/hotel-booking/internal/handlers"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
)

func routes(app *config.AppConfig) http.Handler {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(middleware.Recoverer)	
	router.Use(SessionLoad)
	router.Use(WriteToConsole)
	
	router.HandleFunc("/", handlers.Repo.HandleHome).Methods("GET")
	router.HandleFunc("/about", handlers.Repo.HandleAbout).Methods("GET")
	router.HandleFunc("/generals", handlers.Repo.HandleGenerals).Methods("GET")
	router.HandleFunc("/major", handlers.Repo.HandleMajor).Methods("GET")

	router.HandleFunc("/search-availability", handlers.Repo.HandleSearchAvailability).Methods("GET")
	router.HandleFunc("/search-availability", handlers.Repo.PostAvailability).Methods("POST")
	router.HandleFunc("/search-availability-json", handlers.Repo.AvailabilityJSON).Methods("POST")
	router.HandleFunc("/choose-room/{id}", handlers.Repo.ChooseRoom).Methods("GET")
	router.HandleFunc("/book-room", handlers.Repo.BookRoom).Methods("GET")

	router.HandleFunc("/make-reservation", handlers.Repo.HandlerMakeReservation).Methods("GET")
	router.HandleFunc("/post-reservation", handlers.Repo.PostReservation).Methods("POST")
	router.HandleFunc("/reservation-summary", handlers.Repo.ReservationSummary)

	fs := http.FileServer(http.Dir("./static/"))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	return router
}
