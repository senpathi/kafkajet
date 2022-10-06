package datastore

import (
	"github.com/senpathi/kafkajet/internal/datastore/mongodb"
	"github.com/senpathi/kafkajet/internal/datastore/repository"
	"log"
)

var database repository.DB

func Init() {
	db, err := mongodb.Init(repository.DBConf)
	if err != nil {
		log.Fatal(err)
	}

	database = db
}
