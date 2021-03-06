package config

import (
	"html/template"
	"log"

	"github.com/NganJason/hotel-booking/internal/models"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache 	map[string]*template.Template
	InProduction 	bool
	Session 		*scs.SessionManager
	InfoLog 		*log.Logger
	ErrorLog 		*log.Logger
	MailChan		chan models.MailData
}