package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMediaManager_Retrieve(t *testing.T) {
	retrieveMediaRequest := []byte(`{
    "action": "retrieveMediaFile",
    "version": 6,
    "params": {
        "filename": "_hello.txt"
    }
  }`)
	retrieveMediaResult := []byte(`{
    "result": "SGVsbG8sIHdvcmxkIQ==",
    "error": null
  }`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t,
			retrieveMediaRequest,
			retrieveMediaResult)

		data, restErr := client.Media.RetrieveMediaFile("_hello.txt")
		assert.Equal(t, "SGVsbG8sIHdvcmxkIQ==", *data)
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		data, restErr := client.Media.RetrieveMediaFile("_hello.txt")
		assert.Nil(t, data)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}
