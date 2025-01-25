/*
Saviynt Users API

Testing UsersAPIService

*/

package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_users_UsersAPIService(t *testing.T) {
	apiClient, _, wantTest, err := client()
	require.Nil(t, err)

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {
		if !wantTest {
			t.Skip("skip test") // remove to run test
		}

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Greater(t, len(resp.Userdetails), 0)
	})
}
