package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDecksManager_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

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
		defer httpmock.Reset()

		registerErrorResponse(t)

		decks, restErr := client.Decks.GetAll()
		assert.Nil(t, decks)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})

	t.Run("http request error", func(t *testing.T) {
		defer httpmock.Reset()

		decks, restErr := client.Decks.GetAll()
		assert.Nil(t, decks)
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusInternalServerError, restErr.StatusCode)
		assert.Equal(t, "Internal Server Error", restErr.Message)
	})
}

func TestDecksManager_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[int64])
		loadTestData(t, testDataPath+ActionCreateDeck+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Decks.Create("test")
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		restErr := client.Decks.Create("test")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}

func TestDecksManagerDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		result := new(Result[string])
		loadTestData(t, testDataPath+ActionDeleteDecks+jsonExt, result)
		responder, err := httpmock.NewJsonResponder(http.StatusOK, result)
		assert.NoError(t, err)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		restErr := client.Decks.Delete("test")
		assert.Nil(t, restErr)
	})

	t.Run("error", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		restErr := client.Decks.Delete("test")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})
}
