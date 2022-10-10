package boot

import (
	"github.com/senpathi/kafkajet/internal/service"
	"os"
	"os/signal"

	"github.com/senpathi/kafkajet/internal/datastore"
	"github.com/senpathi/kafkajet/internal/datastore/repository"
	"github.com/senpathi/kafkajet/internal/http"
)

func Boot() {
	/******************-Load config-*************************************/
	loadConfig()

	/******************-Init  Repository-*************************************/
	datastore.Init()
	clusterRepo := datastore.NewClusterRepository()

	/******************-Init  Service Layer-*************************************/
	kService := service.NewService(clusterRepo)

	/******************-Init Http Handlers-*************************************/
	router := &http.Router{}

	router.Init(kService)

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
