package apiserver

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"net/http"
	"restAPI/internal/app/store/sqlstore"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	store := sqlstore.New(db)
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
