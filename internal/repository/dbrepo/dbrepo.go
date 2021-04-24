package dbrepo

import (
	"database/sql"

	"github.com/NganJason/hotel-booking/internal/config"
	"github.com/NganJason/hotel-booking/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}