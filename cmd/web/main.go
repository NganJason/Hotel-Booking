package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NganJason/hotel-booking/internal/config"
	"github.com/NganJason/hotel-booking/internal/handlers"
	"github.com/NganJason/hotel-booking/internal/helpers"
	"github.com/NganJason/hotel-booking/internal/models"
	"github.com/NganJason/hotel-booking/internal/render"
	"github.com/alexedwards/scs/v2"
)

// Port number for server
const PORT_NUMBER = ":8080"
var app config.AppConfig
var session * scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {	
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server is listening to %s", PORT_NUMBER)

	http.Handle("/", routes(&app))
	http.ListenAndServe(":8080", nil)

}

func run() error{
		// Define type of value to store in session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot create template cache %v", err))
		return err
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	// Initiate repository pattern
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	return nil
}