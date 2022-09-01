package handlers

import (
	"github.com/senpathi/kafkajet/internal/http/response"
	"net/http"
)

type ViewTopicsHandler struct {
}

func (v *ViewTopicsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	testTopics := []string{
		"topic1",
		"topic2",
		"topic3",
	}

	response.MakeJson(w, response.ViewTopics{Topics: testTopics}, nil)
}
