package response

import (
	"encoding/json"
	"net/http"

	"github.com/senpathi/kafkajet/internal/domain"
)

type GenericResponse struct {
	Payload  any              `json:"payload"`
	Error    *domain.Error    `json:"error"`
	Metadata *domain.Metadata `json:"metadata"`
}

func MakeJson(w http.ResponseWriter, payload any, err error) {
	res := GenericResponse{}
	code := http.StatusOK
	if err != nil {
		var dmErr domain.Error
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
		var dmErr domain.Error
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

func unwrapError(err error) (domain.Error, int) {
	dmErr, ok := err.(domain.Error)
	if !ok {
		dmErr = domain.Error{
			Code:    "5000",
			Message: "internal server error",
		}

		return dmErr, http.StatusInternalServerError
	}

	// todo: decide response status code base on the error code
	return dmErr, http.StatusBadRequest
}
