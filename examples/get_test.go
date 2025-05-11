package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xmarlem/go-httpclient/gohttpmock"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package examples")

	// Tells the HTTP library to mock any further requests from here.
	gohttpmock.StartMockServer()

	os.Exit(m.Run())
}

func TestUserAgent(t *testing.T) {

	gohttpmock.DeleteMocks()
	gohttpmock.AddMock(gohttpmock.Mock{
		Method:       http.MethodGet,
		Url:          "https://api.github.com",
		ResponseBody: `{"current_user_url": "https://api.github.com/user"}`,
	})

	ep, err := GetEndpointsWithUserAgent()

	require.NoError(t, err)
	require.NotEmpty(t, ep)

}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		gohttpmock.DeleteMocks()
		// Initialization
		gohttpmock.AddMock(gohttpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timemout getting github endpoints"),
		})

		// execution
		endpoints, err := GetEndpoints()

		// validation
		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}
		if err.Error() != "timemout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttpmock.DeleteMocks()

		// Initialization
		gohttpmock.AddMock(gohttpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": 123}`,
			ResponseStatusCode: http.StatusOK,
		})

		// execution
		endpoints, err := GetEndpoints()

		// validation
		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("an error was expected")
		}
		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message received")
		}

	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttpmock.DeleteMocks()

		// Initialization
		gohttpmock.AddMock(gohttpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
			ResponseStatusCode: http.StatusOK,
		})
		// execution
		endpoints, err := GetEndpoints()
		// validation
		if err != nil {
			t.Errorf("no error was expected and we got: %s", err)
		}
		if endpoints == nil {
			t.Error("endpoints were expected at this point and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}

	})

}
