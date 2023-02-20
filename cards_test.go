package ankiconnect

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCardsManager_Add(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer httpmock.Reset()

		// Get will do two api calls, first findCards to get the card id's
		findResult := new(Result[[]int64])
		loadTestData(t, testDataPath+ActionFindCards+"Result"+jsonExt, findResult)
		findResponse, err := httpmock.NewJsonResponse(http.StatusOK, findResult)
		assert.NoError(t, err)

		// Then cardsInfo to transform those into actual anki cards 
		infoResult := new(Result[[]ResultCardsInfo])
		loadTestData(t, testDataPath+ActionCardsInfo+"Result"+jsonExt, infoResult)
		assert.Equal(t, infoResult.Result[0].ModelName, "Basic")
		infoResponse, err := httpmock.NewJsonResponse(http.StatusOK, infoResult)
		assert.NoError(t, err)

		responder := httpmock.ResponderFromMultipleResponses(
			[]*http.Response{
				findResponse,
				infoResponse,
			},
		)

		httpmock.RegisterResponder(http.MethodPost, ankiConnectUrl, responder)

		// If possible test to make sure the payload is properly transformed into
		// findNotesPayload.json (which seems to be attempted in above tests but is not
		// actually working)
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
