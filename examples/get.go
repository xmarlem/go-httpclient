package examples

import "fmt"

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

	fmt.Println(fmt.Sprintf("Status code %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body %s", response.String()))

	var endpoints Endpoints

	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))

	return &endpoints, nil

}
