package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSyncManager_Trigger(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[string])
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Sync.Trigger()
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		restErr := client.Sync.Trigger()
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}
