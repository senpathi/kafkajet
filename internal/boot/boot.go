package boot

import (
	"os"
	"os/signal"

	"github.com/senpathi/kafkajet/internal/datastore"
	"github.com/senpathi/kafkajet/internal/datastore/repository"
	"github.com/senpathi/kafkajet/internal/http"
	"github.com/senpathi/kafkajet/internal/kafka"
)

func Boot() {
	/******************-Load config-*************************************/
	loadConfig()

	/******************-Init Kafka Client-*************************************/
	kClient := kafka.NewClient()

	/******************-Init  Repository-*************************************/
	clusterRepo := datastore.NewClusterRepository()

	/******************-Init Http Handlers-*************************************/
	router := &http.Router{}

	router.Init(kClient)

	/******-Handle OS interrupt Signals-*****/
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Kill, os.Interrupt)
	go func() {
		<-sig
		router.Stop()
	}()

	router.Start()
}

func loadConfig() {
	(&repository.DBConfig{}).Load()
}
