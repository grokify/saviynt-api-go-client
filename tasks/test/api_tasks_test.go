/*
Saviynt Tasks API

Testing TasksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tasks

import (
	"context"
	"testing"

	openapiclient "github.com/saviynt/saviynt-api-go-client/tasks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_tasks_TasksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TasksAPIService CheckTaskStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TasksAPI.CheckTaskStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TasksAPIService UpdateTasks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TasksAPI.UpdateTasks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
