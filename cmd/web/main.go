package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NganJason/hotel-booking/pkg/config"
	"github.com/NganJason/hotel-booking/pkg/handlers"
	"github.com/NganJason/hotel-booking/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// Port number for server
const PORT_NUMBER = ":8080"
var app config.AppConfig
var session * scs.SessionManager

func main() {	

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	// Initiate repository pattern
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Handlers
	// http.HandleFunc("/", handlers.Repo.HandleHome)
	// http.HandleFunc("/about", handlers.Repo.HandleAbout)

	fmt.Printf("Server is listening to %s", PORT_NUMBER)
	// _ = http.ListenAndServe(PORT_NUMBER, nil) 

	srv := &http.Server {
		Addr: PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

