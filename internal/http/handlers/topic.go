package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/senpathi/kafkajet/internal/domain"
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

type CreateTopicsHandler struct {
	Client kafka.Client
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

	topics, err := c.Client.CreateTopics(topicDetails)

	response.MakeJson(w, response.ViewTopics{Topics: topics}, err)
}
