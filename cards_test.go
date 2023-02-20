package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCardsManager_Get(t *testing.T) {
	findCardsPayload := []byte(`{
    "action": "findCards",
    "version": 6,
    "params": {
        "query": "deck:current"
    }
  }`)

	cardsInfoPayload := []byte(`{
    "action": "cardsInfo",
    "version": 6,
    "params": {
        "cards": [1498938915662, 1502098034048]
    }
  }`)

	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		registerMultipleVerifiedPayloads(t,
			[][2][]byte{
				// Get will do two api calls, first findCards to get the card id's
				{
					findCardsPayload,
					loadTestResult(t, ActionFindCards),
				},
				// Then cardsInfo to transform those into actual anki cards
				{
					cardsInfoPayload,
					loadTestResult(t, ActionCardsInfo),
				},
			})

		payload := "deck:current"
		notes, restErr := client.Cards.Get(payload)
		assert.Nil(t, restErr)
		assert.Equal(t, len(*notes), 2)

	})

	t.Run("errorFailSearch", func(t *testing.T) {
		defer httpmock.Reset()

		registerErrorResponse(t)

		_, restErr := client.Cards.Get("deck:current")
		assert.NotNil(t, restErr)
		assert.Equal(t, http.StatusBadRequest, restErr.StatusCode)
		assert.Equal(t, "some error message", restErr.Message)
	})

}
