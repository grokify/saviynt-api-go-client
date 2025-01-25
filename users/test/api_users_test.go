/*
Saviynt Users API

Testing UsersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package users

import (
	"context"
	"testing"

	openapiclient "github.com/grokify/saviynt-api-go-client/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_users_UsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
