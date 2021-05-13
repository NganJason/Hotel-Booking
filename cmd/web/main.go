package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/NganJason/hotel-booking/internal/config"
	"github.com/NganJason/hotel-booking/internal/driver"
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
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer close(app.MailChan)
	listenForMail()

	
	fmt.Printf("Server is listening to %s", PORT_NUMBER)

	http.Handle("/", routes(&app))
	http.ListenAndServe(":8080", nil)

}

func run() (*driver.DB, error) {
	// Define type of value to store in session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan


	// read flags
	// inProduction := flag.Bool("production", true, "Application is in production")
	// useCache := flag.Bool("cache", true, "Use template cache")
	// dbHost := flag.String("dbhost", "localhost", "Database host")
	// dbName := flag.String("dbname", "", "Database name")
	// dbUser := flag.String("dbuser", "", "Database user")
	// dbPass := flag.String("dbpass", "", "Database password")
	// dbPort := flag.String("dbport", "5432", "Database port")
	// dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable prefer require)")

	// flag.Parse()

	// if *dbName == "" || *dbUser == "" {
	// 	fmt.Println("Missing required flags")
	// 	os.Exit(1)
	// }
	// Change this to true when in production
	app.InProduction = true

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

	// connect to database
	log.Println("Connecting to database...")
	// connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
	// db, err := driver.ConnectSQL(connectionString)
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=hotel-booking user=jason.ngan password=")
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot create template cache %v", err))
		return nil, err
	}

	app.TemplateCache = tc
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	// Initiate repository pattern
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	return db, nil
}