package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/senpathi/kafkajet/internal/domain"
	"github.com/senpathi/kafkajet/internal/http/response"
	"github.com/senpathi/kafkajet/internal/service"
)

type ViewTopicsHandler struct {
	Service *service.ClusterService
}

func (v *ViewTopicsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	topics, err := v.Service.ViewClusterTopics("localhost")

	response.MakeJson(w, response.ViewTopics{Topics: topics}, err)
}

type CreateTopicsHandler struct {
	Service *service.ClusterService
}

func (c *CreateTopicsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.MakeJson(w, nil, err)
	}

	var topicDetails []domain.TopicDetails
	err = json.Unmarshal(body, &topicDetails)
	if err != nil {
		response.MakeJson(w, nil, err)
	}

	topics, err := c.Service.CreateTopics("localhost", topicDetails)

	response.MakeJson(w, response.ViewTopics{Topics: topics}, err)
}
