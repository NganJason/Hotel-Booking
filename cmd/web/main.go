package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NganJason/hotel-booking/internal/config"
	"github.com/NganJason/hotel-booking/internal/handlers"
	"github.com/NganJason/hotel-booking/internal/models"
	"github.com/NganJason/hotel-booking/internal/render"
	"github.com/alexedwards/scs/v2"
)

// Port number for server
const PORT_NUMBER = ":8080"
var app config.AppConfig
var session * scs.SessionManager

func main() {	

	// Define type of value to store in session
	gob.Register(models.Reservation{})

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
		log.Fatal(fmt.Sprintf("Cannot create template cache %v", err))
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	// Initiate repository pattern
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)


	fmt.Printf("Server is listening to %s", PORT_NUMBER)

	http.Handle("/", routes(&app))
	http.ListenAndServe(":8080", nil)

}

