package examples

import (
	"net/http"
	"time"

	"github.com/xmarlem/go-httpclient/gohttp"
	"github.com/xmarlem/go-httpclient/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJSON)
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Pippo"). // if already defined as common header, this one is not used
		Build()

	return client
}
