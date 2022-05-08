package data

import (
	"database/sql"
	"library/config"
	"log"
)

type DatabaseSource struct {
	Driver        string
	User          string
	Password      string
	ServerAddress string
	DatabaseName  string
}

func CreateConnectionToDatabase(ds config.DatabaseSource) *sql.DB {
	db, err := sql.Open(ds.Driver, ds.User+":"+ds.Password+ds.ServerAddress+"/"+ds.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
