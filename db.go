package main

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg"
	"log"
	"os"
	"time"
)

var Conn *pg.DB

type DbRepository interface {
	SaveAssets(Assets) error
	SaveError(string) error
}

func Connect() (*pg.DB, error) {
	if Conn == nil {
		var user, pwd, database, addr string

		database = "pg_database"
		addr = "db_host"
		user = "pg_user"
		pwd = "pg_password"

		db := pg.Connect(&pg.Options{
			Addr:        addr,
			Database:    database,
			User:        user,
			Password:    pwd,
			IdleTimeout: 5 * time.Minute,
			MaxAge:      15 * time.Minute,
		})
		if db != nil {
			Conn = db
		} else {
			return nil, errors.New("CONNECT")
		}

		logger := log.New(os.Stdout, "", log.LstdFlags)
		pg.SetLogger(logger)

		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			fmt.Printf("[pg] %s %s\n", time.Since(event.StartTime), query)
		})
	}
	return Conn, nil
}

func (c *CryptoCompare) SaveAssets(Assets) error {

	return nil
}
func (c *CryptoCompare) SaveError(string) error {
	return nil
}
