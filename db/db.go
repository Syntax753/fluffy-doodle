package db

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/syntax753/fluffy-doodle/models"
	model "github.com/syntax753/fluffy-doodle/models"
)

type config struct {
	Schema string
}

const configFile = "config.toml"

var (
	conf config
)

type Env struct {
	db models.Datastore
}

func NewDB(string)

func init() {
	log.Println("Initialising database")
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		log.Fatalf("Can't open config file %v: %v\n", configFile, err)
	}

	data, err := ProcessFile(conf.Schema)

	if err != nil {
		log.Fatalf("Can't load transactions from %v: %v\n", conf.Schema, err)
	}

	Env := &DB{TXs: data.AsMap}

	log.Println("Database OK")
}

// TXNotFound is the error for when a transaction can not be found
type TXNotFound struct {
	ID string
}

func (tx *TXNotFound) Error() string {
	return fmt.Sprintf("TX with ID %s not found", tx.ID)
}

// Find returns a transaction for an ID
func (db *DB) Find(ID string) (model.TX, error) {

	if v, ok := db.TXs[ID]; ok {
		return v, nil
	}

	return model.TX{}, &TXNotFound{ID}
}
