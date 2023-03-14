package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSyncManager_Trigger(t *testing.T) {

	syncRequest := []byte(`{
  "action": "sync",
  "version": 6
}`)
	syncResult := []byte(`{
  "result": null,
  "error": null
}`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t, syncRequest, syncResult)

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
