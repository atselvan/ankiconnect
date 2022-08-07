package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/privatesquare/bkst-go-utils/utils/fileutils"
	"github.com/stretchr/testify/assert"
)

const (
	testDataPath          = "data/test/"
	errorTestDataFileName = "error.json"
	jsonExt               = ".json"
)

func TestDecksManager_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[[]string])
		loadTestData(t, testDataPath+ActionDeckNames+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		decks, restErr := client.Decks.GetAll()
		assert.NotNil(t, decks)
		assert.Nil(t, restErr)
		assert.Equal(t, 3, len(*decks))
	})

	t.Run("error", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[[]string])
		loadTestData(t, testDataPath+errorTestDataFileName, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		decks, restErr := client.Decks.GetAll()
		assert.Nil(t, decks)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})

	t.Run("http request error", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		decks, restErr := client.Decks.GetAll()
		assert.Nil(t, decks)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusInternalServerError, restErr.StatusCode)
		assert.Equal(t, "Internal Server Error", restErr.Message)
	})
}

func TestDecksManager_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[int64])
		loadTestData(t, testDataPath+ActionCreateDeck+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Decks.Create("test")
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

		restErr := client.Decks.Create("test")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestDecksManagerDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.ActivateNonDefault(client.httpClient.GetClient())
		defer httpmock.DeactivateAndReset()

		result := new(Result[string])
		loadTestData(t, testDataPath+ActionDeleteDecks+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Decks.Delete("test")
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

		restErr := client.Decks.Delete("test")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func loadTestData(t *testing.T, path string, out interface{}) {
	err := fileutils.ReadJsonFile(path, &out)
	assert.NoError(t, err)
}
