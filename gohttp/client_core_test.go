package gohttp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// initialization
	client := httpClient{}

	commonHeaders := make(http.Header)

	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders
	// execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// validation
	if len(finalHeaders) != 3 {
		t.Errorf("we expect 3 headers, got %d", len(finalHeaders))
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Content-Type header is not as expected")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("User-Agent header is not as expected")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("X-Request-Id header is not as expected")
	}
}

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
