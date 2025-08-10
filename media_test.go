package ankiconnect

import (
	"encoding/base64"
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

func TestMediaManager_Store(t *testing.T) {
	request := []byte(`{
		"action": "storeMediaFile",
		"version": 6,
		"params": {
			"filename": "_test_store.txt",
			"data": "c3RvcmUgbWVkaWEgZmlsZSB0ZXN0"
		}
	}`)

	response := []byte(`{
		"result": "_test_store.txt",
		"error": null
	}`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t, request, response)

		filename := "_test_store.txt"
		encodedMediaContent := base64.StdEncoding.EncodeToString([]byte("store media file test"))
		savedFileName, restRrr := client.Media.StoreMediaFile(filename, encodedMediaContent)
		assert.Nil(t, restRrr)
		assert.Equal(t, filename, *savedFileName)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		savedFileName, restErr := client.Media.StoreMediaFile("_test_store.txt", "some media content")
		assert.NotNil(t, restErr)
		assert.Nil(t, savedFileName)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestManager_Get(t *testing.T) {
	request := []byte(`{
		"action": "getMediaFileNames",
		"version": 6,
		"params": {
			"pattern": "*"
		} 
	}`)

	response := []byte(`{
		"result": ["_test_file_1.txt", "_test_file_2.txt"],
		"error": null
	}`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t, request, response)

		filenames, restErr := client.Media.GetMediaFileNames("*")
		assert.Nil(t, restErr)
		assert.Equal(t, []string{"_test_file_1.txt", "_test_file_2.txt"}, *filenames)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		filenames, restErr := client.Media.GetMediaFileNames("*")
		assert.NotNil(t, restErr)
		assert.Nil(t, filenames)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestManager_Delete(t *testing.T) {
	request := []byte(`{
		"action": "deleteMediaFile",
		"version": 6,
		"params": {
			"filename": "_delete_file_name.txt"
		}
	}`)
	response := []byte(`{
		"result": "_delete_file_name.txt",
		"error": null
	}`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerVerifiedPayload(t, request, response)

		deletedFileName, restErr := client.Media.DeleteMediaFile("_delete_file_name.txt")
		assert.Nil(t, restErr)
		assert.Equal(t, "_delete_file_name.txt", *deletedFileName)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		deletedFileName, restErr := client.Media.DeleteMediaFile("_delete_file_name.txt")
		assert.NotNil(t, restErr)
		assert.Nil(t, deletedFileName)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}
