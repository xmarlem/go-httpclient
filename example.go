package main

import (
	"fmt"

	"github.com/xmarlem/go-httpclient/gohttp"
)

var githubHttpClient = getGithubClient()

func getGithubClient() gohttp.Client {

	builder := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5)

	client := builder.Build()

	return client
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// func createUser(user User) {

// 	resp, err := githubHttpClient.Post("https://api.github.com", nil, user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	bytes, _ := io.ReadAll(resp.Body)
// 	fmt.Println(string(bytes))

// }

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())
}

func main() {

	getUrls()
}
