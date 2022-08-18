package db

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var cfg = config{}

func LoadConfigs(flagSet *flag.FlagSet) {
	flagSet.StringVar(&cfg.driver, "db_driver", os.Getenv("DATABASE_DRIVER"), "database's driver")
	flagSet.StringVar(&cfg.dbUser, "db_user", os.Getenv("DATABASE_USERNAME"), "database's username")
	flagSet.StringVar(&cfg.dbPass, "db_pass", os.Getenv("DATABASE_PASSWORD"), "database's password")
	flagSet.StringVar(&cfg.dbHost, "db_host", os.Getenv("DATABASE_HOSTNAME"), "database's hostname")
	flagSet.IntVar(&cfg.dbPort, "db_port", mustParsePort(), "database's port")
	flagSet.StringVar(&cfg.dbName, "db_name", os.Getenv("DATABASE_NAME"), "database name")
}

func mustParsePort() int {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln(err)
	}
	return port
}

type Config interface {
	Source() string
	Driver() string
	String() string
}

type config struct {
	driver string
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
}

func (c config) Driver() string {
	return c.driver
}

func (c config) String() string {
	return fmt.Sprintf("[%s] %s", c.driver, c.Source())
}

func (c config) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName)
}

func DefaultConfig() Config {
	return cfg
}
