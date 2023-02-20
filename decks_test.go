package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestDecksManager_GetAll(t *testing.T) {
  getAllRequest := []byte(`{
    "action": "deckNames",
    "version": 6
}`)
  getAllResult := []byte(`{
    "result": [
        "Default",
        "Deck01",
        "Deck02"
    ],
    "error": null
}`)
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

    registerVerifiedPayload(t, getAllRequest, getAllResult)

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
  createRequest := []byte(`{
    "action": "createDeck",
    "version": 6,
    "params": {
        "deck": "Japanese::Tokyo"
    }
}`)
  createResponse := []byte(`{
    "result": 1659294179522,
    "error": null
}`)


	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

    registerVerifiedPayload(t, createRequest, createResponse)

    restErr := client.Decks.Create("Japanese::Tokyo")
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
  deleteDeckRequest := []byte(`{
    "action": "deleteDecks",
    "version": 6,
    "params": {
        "decks": ["test"],
        "cardsToo": true
    }
}`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

    registerVerifiedPayload(t, deleteDeckRequest, genericSuccessJson)

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
