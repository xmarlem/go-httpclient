package examples

import (
	"fmt"
	"net/http"
)

type Endpoints struct {
	CurrentUserUrl   string `json:"current_user_url"`
	AuthorizationUrl string `json:"authorizations_url"`
	RepositoryUrl    string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", fmt.Sprintf("Status code %d", response.StatusCode))
	fmt.Printf("%s\n", fmt.Sprintf("Status %s", response.Status))
	fmt.Printf("%s\n", fmt.Sprintf("Body %s", response.String()))

	var endpoints Endpoints

	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))

	return &endpoints, nil

}

func GetEndpointsWithUserAgent() (*Endpoints, error) {

	h := http.Header{}

	response, err := httpClient.Get("https://api.github.com", h)
	if err != nil {
		return nil, err
	}

	var ep Endpoints
	if err = response.UnmarshalJson(&ep); err != nil {
		return nil, err
	}

	return &ep, nil

}
