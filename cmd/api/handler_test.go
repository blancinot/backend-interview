package main

import (
	"database/sql"

	accountapp "github.com/gustvision/backend-interview/pkg/account/app"
	accountsql "github.com/gustvision/backend-interview/pkg/account/sql"
	userapp "github.com/gustvision/backend-interview/pkg/user/app"
	usersql "github.com/gustvision/backend-interview/pkg/user/sql"
)

const (
	dsn = "host=0.0.0.0 port=5432 user=postgres password=secret dbname=postgres sslmode=disable"
)

// TODO: use docker test lib like ory/dockertest to run postgres
// WARN: For now, postgres must be accessible @0.0.0.0
//       (cf docker-compose up or docker run...)
func newTestHandler() (h handler, err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	userSQLStore := usersql.Store{DB: db}
	accountSQLStore := accountsql.Store{DB: db}
	h.user = &userapp.App{
		Store: &userSQLStore,
	}
	h.account = &accountapp.App{
		Store:            &accountSQLStore,
		StoreTransaction: &accountSQLStore,
	}
	return
}
