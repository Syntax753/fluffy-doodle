package db

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/boltdb/bolt"
)

type config struct {
	DBName string
}

const configFile = "config.toml"

var (
	conf config
)

func init() {
	log.Println("Starting database")
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		log.Fatal("Can't open config file\n", err)
	}
}

// NewPayDB sets up the database for payments
func NewPayDB() (*bolt.DB, error) {
	log.Println("Opening database")

	dbName := "../data/" + conf.DBName + ".db"

	

	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Initialise(env string) {

}
