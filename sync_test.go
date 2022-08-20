package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSyncManager_Trigger(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[string])
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Sync.Trigger()
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[string])
		loadTestData(t, testDataPath+errorTestDataFileName, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Sync.Trigger()
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}
