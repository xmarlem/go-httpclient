package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetRequestTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(max int) ClientBuilder
	DisableTimeouts(disableTimeouts bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

// in pratica, con i vari metodi compongo il builder struct... e poi a build time quando si chiama questo metodo
// si prendono tutti gli elementi e si compone httpClient obj
func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) DisableTimeouts(disableTimeouts bool) ClientBuilder {
	c.disableTimeouts = disableTimeouts

	return c
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}
func (c *clientBuilder) SetRequestTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}
func (c *clientBuilder) SetMaxIdleConnections(max int) ClientBuilder {
	c.maxIdleConnections = max
	return c
}
