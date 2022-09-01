package main

import (
	"os"
	"os/signal"

	"github.com/senpathi/kafkajet/internal/http"
)

func main() {
	/******************-Init Http Handlers-*************************************/
	router := &http.Router{}

	router.Init()

	/******-Handle OS interrupt Signals-*****/
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Kill, os.Interrupt)
	go func() {
		<-sig
		router.Stop()
	}()

	router.Start()
}
