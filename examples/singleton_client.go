package examples

import (
	"time"

	"github.com/xmarlem/go-httpclient/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetRequestTimeout(3 * time.Second).
		Build()

	return client
}
