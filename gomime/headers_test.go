package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid content type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("invalid user agent header")
	}

	if ContentTypeJSON != "application/json" {
		t.Error("invalid content type json header")
	}

	if ContentTypeOctetStream != "application/octet-stream" {
		t.Error("invalid content type octet header")
	}
}
