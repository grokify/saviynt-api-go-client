/*
Saviynt Job Control API

Testing JobControlAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package jobcontrol

import (
	"context"
	"testing"

	openapiclient "github.com/saviynt/saviynt-api-go-client/jobcontrol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_jobcontrol_JobControlAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JobControlAPIService CheckJobStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.JobControlAPI.CheckJobStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobControlAPIService CreateTriggers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.JobControlAPI.CreateTriggers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobControlAPIService CreateUpdateTrigger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.JobControlAPI.CreateUpdateTrigger(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobControlAPIService DeleteTrigger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.JobControlAPI.DeleteTrigger(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobControlAPIService FetchJobMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.JobControlAPI.FetchJobMetadata(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobControlAPIService ResumePauseJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.JobControlAPI.ResumePauseJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobControlAPIService RunJobTrigger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.JobControlAPI.RunJobTrigger(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
