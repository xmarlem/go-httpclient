package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	t.Parallel()
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "mocked-http-client")
	client.builder = &clientBuilder{
		headers:   commonHeaders,
		userAgent: "cool-user-agent",
	}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// validation
	if len(finalHeaders) != 3 {
		t.Errorf("we expect 3 headers, got %d", len(finalHeaders))
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("X-Request-Id header is not as expected")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Content-Type header is not as expected")
	}

	if finalHeaders.Get("User-Agent") != "mocked-http-client" {
		t.Error("invalid user agent received")
	}
}
