package gohttp

import (
	"net/http"

	"github.com/xmarlem/go-httpclient/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	merged := http.Header{}
	for _, h := range headers {
		for key, values := range h {
			for _, v := range values {
				merged.Add(key, v)
			}
		}
	}
	return merged
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {

	result := make(http.Header)
	// Add common headers from the HTTP client instance:
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers from the current request:
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Set the user-agent if it's not already there
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}

		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return result

}
