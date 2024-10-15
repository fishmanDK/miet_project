package app

import "github.com/jmoiron/sqlx"

func (a *app) connectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", a.cfg.Postgres.ToString())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}