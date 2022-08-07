package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestNotesManager_Add(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[int64])
		loadTestData(t, testDataPath+ActionAddNote+"Result"+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		note := new(Note)
		loadTestData(t, testDataPath+ActionAddNote+"Payload"+jsonExt, result)
		restErr := client.Notes.Add(*note)
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

		note := new(Note)
		loadTestData(t, testDataPath+ActionAddNote+"Payload"+jsonExt, result)
		restErr := client.Notes.Add(*note)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}
