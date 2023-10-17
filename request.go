package request

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	request *http.Request
}

func New(request *http.Request) Request {
	return Request{
		request: request,
	}
}

func (request Request) DecodeBody(body interface{}) error {
	err := json.NewDecoder(request.request.Body).Decode(&body)

	return err
}
