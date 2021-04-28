package handlers

import (
	"errors"
	"net/http"
)

type ReadPartitionMessageHandler struct {
}

func (r ReadPartitionMessageHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

func (r ReadPartitionMessageHandler) run(res http.ResponseWriter, req *http.Request) error {
	qValues := req.URL.Query()
	clusterId := qValues.Get(`cluster_id`)
	if clusterId == `` {
		return errors.New(`cluster_id is missing in url`)
	}
	clusterId = `1` //todo: remove this

	topic := qValues.Get(`topic`)
	if topic == `` {
		return errors.New(`cluster_id is missing in url`)
	}
	topic = `test`
	return nil
}
