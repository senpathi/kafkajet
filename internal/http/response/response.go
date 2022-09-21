package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/senpathi/kafkajet/internal/domain"
	domainErr "github.com/senpathi/kafkajet/internal/errors"
)

type GenericResponse struct {
	Payload  any              `json:"payload"`
	Error    *domainErr.Error `json:"error"`
	Metadata *domain.Metadata `json:"metadata"`
}

func MakeJson(w http.ResponseWriter, payload any, err error) {
	res := GenericResponse{}
	code := http.StatusOK
	if err != nil {
		var dmErr domainErr.Error
		dmErr, code = unwrapError(err)
		res.Error = &dmErr
	} else {
		res.Payload = payload
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(res)
}

func MakeJsonWithMeta(w http.ResponseWriter, payload any, err error, meta domain.Metadata) {
	res := GenericResponse{}

	code := http.StatusOK
	if err != nil {
		var dmErr domainErr.Error
		dmErr, code = unwrapError(err)
		res.Error = &dmErr
	} else {
		res.Payload = payload
		res.Metadata = &meta
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(res)
}

func unwrapError(err error) (domainErr.Error, int) {
	dmErr, ok := err.(domainErr.Error)
	if !ok {
		dmErr = domainErr.Error{
			Code:    "5000",
			Message: "internal server error",
		}

		fmt.Println(err)

		return dmErr, http.StatusInternalServerError
	}

	fmt.Println(dmErr)

	// todo: decide response status code base on the error code
	return dmErr, http.StatusBadRequest
}
