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
)

type Router struct {
	server *http.Server
}

func (r *Router) Init() {
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
		&handlers.ViewTopicsHandler{},
	).Methods(http.MethodGet)
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
