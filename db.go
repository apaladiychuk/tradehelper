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

		database = "postgres"
		addr = "localhost:5432"
		user = "postgres"
		pwd = "postgres"

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

func (c *CryptoCompare) SaveAssets(a Assets) error {
	a.Time = a.Time * 100

	conn, err := Connect()
	if err != nil {
		return err
	}
	var ea Assets
	if err := conn.Model(&ea).Where("time = ? ", a.Time).Select(); err == nil {
		// skip existing row
		// todo we can update existing row
		fmt.Printf(" skip value for %d \n", a.Time)
		return nil
	}
	if _, err := conn.Model(&a).
		Returning("Id", &a.Id).Insert(); err != nil {
		return err
	}
	return nil
}
func (c *CryptoCompare) SaveError(string) error {
	return nil
}
