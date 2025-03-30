package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/xmarlem/go-httpclient/gohttp"
)

var githubHttpClient = getGithubClient()

func getGithubClient() gohttp.HTTPClient {
	client := gohttp.New()

	client.SetMaxIdleConnections(20)
	client.SetConnectionTimeout(2 * time.Second)
	client.SetRequestTimeout(4 * time.Second)
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {

	resp, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	bytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bytes))

}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)

	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	getUrls()
}
