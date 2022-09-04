package handlers

import (
	"net/http"

	"github.com/senpathi/kafkajet/internal/http/response"
	"github.com/senpathi/kafkajet/internal/kafka"
)

type ViewTopicsHandler struct {
	Client kafka.Client
}

func (v *ViewTopicsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	topics, err := v.Client.Topics()

	response.MakeJson(w, response.ViewTopics{Topics: topics}, err)
}
