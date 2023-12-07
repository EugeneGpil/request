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

func (request Request) GetSimpleQueryParams() map[string]string {
	values := request.request.URL.Query()

	res := make(map[string]string)
	for key := range values {
		res[key] = values.Get(key)
	}

	return res
}
