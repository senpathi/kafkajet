package handlers

import "net/http"

type AddClusterHandler struct {
}

func (a AddClusterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func (a AddClusterHandler) run(writer http.ResponseWriter, request *http.Request) error {

	return nil
}
