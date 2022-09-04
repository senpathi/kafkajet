package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/senpathi/kafkajet/internal/http/handlers"
	"github.com/senpathi/kafkajet/internal/kafka"
)

type Router struct {
	server *http.Server
	client kafka.Client
}

func (r *Router) Init(cli kafka.Client) {
	muxRouter := mux.NewRouter()
	r.server = &http.Server{
		Addr: fmt.Sprintf(":%s", "8080"),
		Handler: muxhandlers.RecoveryHandler(
			muxhandlers.PrintRecoveryStack(true),
		)(muxRouter),
	}

	// view all topics
	muxRouter.Handle(
		"/topics",
		&handlers.ViewTopicsHandler{
			Client: cli,
		},
	).Methods(http.MethodGet)

	// create topics
	muxRouter.Handle(
		"/topics",
		&handlers.CreateTopicsHandler{
			Client: cli,
		},
	).Methods(http.MethodPost)
}

func (r Router) Start() {
	log.Println("starting server ...")
	err := r.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func (r *Router) Stop() {
	log.Println("shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		log.Fatalln(fmt.Sprintf("failed to gracefully shutdown the server: %r", err))
	}
}
