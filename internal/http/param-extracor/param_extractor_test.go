package param_extracor

import (
	"kafkajet/internal/http/request"
	"net/http"
	"testing"
)

func TestParamExtractor_ExtractQuery(t *testing.T) {
	req := &http.Request{
		Method:           "POST",
		URL:              nil,
		Proto:            "",
		ProtoMajor:       0,
		ProtoMinor:       0,
		Header:           nil,
		Body:             nil,
		GetBody:          nil,
		ContentLength:    0,
		TransferEncoding: nil,
		Close:            false,
		Host:             "",
		Form:             nil,
		PostForm:         nil,
		MultipartForm:    nil,
		Trailer:          nil,
		RemoteAddr:       "",
		RequestURI:       "",
		TLS:              nil,
		Cancel:           nil,
		Response:         nil,
	}
	extractor := NewParamExtractor()
	v := request.ReadMessagesQuery{}
	//v := float64(1)
	extractor.ExtractQuery(&v, req)
}
