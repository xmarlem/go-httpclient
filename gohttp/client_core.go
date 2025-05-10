package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/xmarlem/go-httpclient/gomime"
)

const (
	// NOTE: Non voglio settare il client.Timeout globale... meglio usare context.WithTimeout per limitare da fuori ...
	// in quanto il timeout globale non e' retry-safe.
	//
	defaultMaxIdleConnections    = 5
	defaultResponseHeaderTimeout = 10 * time.Second
	defaultConnectionTimeout     = 1 * time.Second
	defaultIdleConnTimeout       = 90 * time.Second //
)

// given the body in any format, it parses and returns the body as a slice of bytes
func (c *httpClient) getRequestBody(contentType string, body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	// devo validare se e' un json, xml etc...
	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJSON:
		return json.Marshal(body)
	case gomime.ContentTypeXML:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}

}

// CORE function implementing all the calls
func (c *httpClient) do(
	method, url string,
	headers http.Header,
	body interface{},
) (*Response, error) {

	// collect all the headers into one single place (both common and request headers)
	fullHeaders := c.getRequestHeaders(headers)

	// if body is nil --> we get nil, nil
	requestBody, err := c.getRequestBody(fullHeaders.Get(gomime.HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	// For debug
	// for _, h := range fullHeaders {
	// 	fmt.Println(h)
	// }

	// required for testing...
	if mock := mockupServer.getMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	// Create a new http request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %v", err)
	}

	req.Header = fullHeaders

	client := c.getHttpClient()

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	finalResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       responseBody,
	}

	return &finalResponse, nil

}

// factory method implemented as a Singleton (do once!)
func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {

		// se il caller ha definito un custom http client dall'esterno, allora usa quello, altrimenti ...
		// if the user provided its own custom client...
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseHeaderTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
				IdleConnTimeout: c.getIdleConnTimeout(),
			},
		}
	})

	return c.client
}

// =======
// Getters
// =======

func (c *httpClient) getIdleConnTimeout() time.Duration {
	if c.builder.idleConnTimeout > 0 {
		return c.builder.idleConnTimeout
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getResponseHeaderTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disableTimeouts {
		return 0
	}

	return defaultResponseHeaderTimeout
}
