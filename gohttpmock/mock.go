package gohttpmock

import (
	"fmt"
	"net/http"

	"github.com/xmarlem/go-httpclient/core"
)

type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*core.Response, error) {

	if m.Error != nil {
		return nil, m.Error
	}

	response := core.Response{
		Status: fmt.Sprintf(
			"%d %s",
			m.ResponseStatusCode,
			http.StatusText(m.ResponseStatusCode),
		),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}
	return &response, nil
}
