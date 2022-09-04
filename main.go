package main

import (
	"github.com/senpathi/kafkajet/internal/kafka"
	"os"
	"os/signal"

	"github.com/senpathi/kafkajet/internal/http"
)

func main() {
	/******************-Init Kafka Client-*************************************/
	kClient := kafka.NewClient()

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
