package gohttp

import (
	"fmt"
	"testing"
)

func TestGetRequestBody(t *testing.T) {

	client := httpClient{}

	testCases := []struct {
		desc        string
		contentType string
		body        func() interface{}
		wantErr     bool
	}{
		{
			desc:        "nil",
			contentType: "",
			body: func() interface{} {
				return nil
			},
		},
		{
			desc:        "json",
			contentType: "application/json",
			body: func() interface{} {
				body := make(map[string]string)
				body["aaa"] = "bbb"
				return body
			},
		},
		{
			desc:        "json",
			contentType: "something else",
			body: func() interface{} {
				body := make(map[string]string)
				body["aaa"] = "bbb"
				return body
			},
		},
		{
			desc:        "xml",
			contentType: "application/xml",
			body: func() interface{} {
				return []string{"one", "two"}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			body, err := client.getRequestBody(tC.contentType, tC.body())

			if err != nil && !tC.wantErr {
				t.Errorf("unexpected error: %v", err)
			}

			fmt.Println(string(body))
		})
	}

}
