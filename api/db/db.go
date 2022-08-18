package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func New(configs ...Config) *sql.DB {
	c := resolveConfigs(configs...)
	conn, err := sql.Open(c.Driver(), c.Source())
	if err != nil {
		log.Fatalln(err)
	}
	// https://www.alexedwards.net/blog/configuring-sqldb
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(time.Minute * 5)
	return conn
}

func resolveConfigs(configs ...Config) Config {
	var c Config
	if len(configs) > 0 {
		c = configs[0]
	} else {
		c = cfg
	}
	log.Println("CONNECTION STRING:", c.String())
	return c
}
